package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/briandowns/spinner"

	"ais/internal/gemini"
	"ais/internal/render"
	"ais/internal/repl"
)

func main() {
	query := flag.String("q", "", "Query to ask in one-shot mode")
	showRefs := flag.Bool("show-refs", false, "Show source citation block")
	flag.Parse()

	if *query != "" {
		if strings.TrimSpace(*query) == "" {
			fmt.Fprintln(os.Stderr, "error: query cannot be empty")
			flag.Usage()
			os.Exit(1)
		}
		if err := runOneShot(context.Background(), *query, *showRefs); err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
		return
	}

	// Chat REPL mode (MODE-02, D-11)
	if err := repl.Run(context.Background(), *showRefs); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

// classifyAPIError wraps an SDK error with a human-readable suggestion based on
// the error category. Substring matching is used because the Gemini SDK does not
// expose typed sentinel errors (per D-04).
func classifyAPIError(err error) error {
	msg := err.Error()
	switch {
	case strings.Contains(msg, "403") || strings.Contains(msg, "API key not valid"):
		return fmt.Errorf("%w — verify your API key is valid and has not expired", err)
	case strings.Contains(msg, "429") || strings.Contains(msg, "Resource exhausted"):
		return fmt.Errorf("%w — you have exceeded your API quota — wait before retrying", err)
	case strings.Contains(msg, "connection refused") || strings.Contains(msg, "deadline exceeded") || strings.Contains(msg, "no such host"):
		return fmt.Errorf("%w — check your internet connection", err)
	default:
		return err
	}
}

// runOneShot executes a single query and exits. Creates a fresh Client so no
// prior history is attached (one-shot = stateless, per D-10).
func runOneShot(ctx context.Context, query string, showRefs bool) error {
	client, err := gemini.NewClient(ctx)
	if err != nil {
		return err
	}

	// Animated spinner while waiting for API response (D-08, D-09)
	s := spinner.New(spinner.CharSets[14], 80*time.Millisecond)
	s.Suffix = " Thinking..."
	s.Start()

	resp, err := client.Ask(ctx, query)
	s.Stop()
	if err != nil {
		return fmt.Errorf("query failed: %w", classifyAPIError(err))
	}

	// Render markdown response (OUT-01, D-13)
	render.Markdown(gemini.ResponseText(resp))

	// Print source citations (OUT-02, SRCH-02, D-05, D-06)
	if showRefs {
		render.Sources(gemini.ExtractSources(resp))
	}

	return nil
}
