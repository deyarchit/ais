package gemini

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/cenkalti/backoff/v5"
	"google.golang.org/genai"
)

// Client wraps the Gemini generative model with always-on Google Search grounding.
type Client struct {
	chat   *genai.Chat
	config *genai.GenerateContentConfig
}

// ClientConfig holds tunable generation parameters supplied by the caller.
// Temperature is the sampling temperature in [0.0, 2.0].
// ThinkingBudget is the reasoning token budget; -1 means dynamic (auto), 0 disables thinking.
// Values are pre-validated by the caller (cmd/ais/main.go).
type ClientConfig struct {
	Temperature    float32
	ThinkingBudget int32
}

// NewClient creates a new Client using GEMINI_API_KEY from environment.
// It configures the model with Google Search grounding enabled (no toggle).
// Returns an error if GEMINI_API_KEY is not set.
func NewClient(ctx context.Context, cfg ClientConfig) (*Client, error) {
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("GEMINI_API_KEY is not set")
	}

	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey: apiKey,
	})
	if err != nil {
		return nil, fmt.Errorf("create gemini client: %w", err)
	}

	// Always-on Google Search grounding (D-12, SRCH-01)
	temp := cfg.Temperature
	config := &genai.GenerateContentConfig{
		Tools: []*genai.Tool{
			{GoogleSearch: &genai.GoogleSearch{}},
			{URLContext: &genai.URLContext{}},
		},
		Temperature: &temp,
	}

	// ThinkingBudget: -1 = auto (omit ThinkingConfig), anything else = explicit budget
	if cfg.ThinkingBudget != -1 {
		budget := cfg.ThinkingBudget
		config.ThinkingConfig = &genai.ThinkingConfig{
			ThinkingBudget: &budget,
		}
	}

	chat, err := client.Chats.Create(ctx, "gemini-2.5-flash", config, nil)
	if err != nil {
		return nil, fmt.Errorf("create chat session: %w", err)
	}

	return &Client{
		chat:   chat,
		config: config,
	}, nil
}

// isRetryable returns true for transient errors that should be retried.
// Gemini occasionally returns 503 UNAVAILABLE and 429 rate-limit errors transiently.
func isRetryable(err error) bool {
	msg := err.Error()
	return strings.Contains(msg, "503") ||
		strings.Contains(msg, "UNAVAILABLE") ||
		strings.Contains(msg, "429") ||
		strings.Contains(msg, "Resource exhausted")
}

// Ask sends a message and returns the full text response and grounding metadata.
// For one-shot mode, call NewClient before each query so history is fresh.
// For chat mode, reuse the same Client across turns to preserve history (MODE-03).
// Transient 503/UNAVAILABLE and 429/Resource exhausted errors are retried with
// exponential backoff (max 30s interval). Each retry prints a message to stderr.
func (c *Client) Ask(ctx context.Context, prompt string) (*genai.GenerateContentResponse, error) {
	bo := backoff.NewExponentialBackOff()
	bo.MaxInterval = 30 * time.Second

	attempt := 0
	resp, err := backoff.Retry(ctx, func() (*genai.GenerateContentResponse, error) {
		r, e := c.chat.Send(ctx, genai.NewPartFromText(prompt))
		if e != nil {
			if isRetryable(e) {
				return nil, e // retryable — backoff will retry
			}
			return nil, backoff.Permanent(e) // non-retryable — stop immediately
		}
		return r, nil
	},
		backoff.WithBackOff(bo),
		backoff.WithNotify(func(e error, wait time.Duration) {
			attempt++
			fmt.Fprintf(os.Stderr, "retrying after %.1fs (attempt %d): %v\n", wait.Seconds(), attempt, e)
		}),
	)
	if err != nil {
		return nil, fmt.Errorf("send message: %w", err)
	}
	return resp, nil
}

// ResponseText extracts the concatenated text from a response.
func ResponseText(resp *genai.GenerateContentResponse) string {
	return resp.Text()
}
