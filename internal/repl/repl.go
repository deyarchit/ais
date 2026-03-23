package repl

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/briandowns/spinner"

	"ais/internal/gemini"
	"ais/internal/render"
)

// classifyAPIError appends a human-readable suggestion to API errors.
// Mirrors the classification in cmd/ais/main.go — both call sites need this.
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

// Run starts the interactive REPL loop. It blocks until the user sends
// Ctrl+C (SIGINT) or Ctrl+D (EOF on stdin). No welcome header is printed (D-03).
// The same gemini.Client is reused across turns to preserve ChatSession history (D-11, MODE-03).
// showRefs controls whether source citations are printed after each response.
func Run(ctx context.Context, showRefs bool) error {
	client, err := gemini.NewClient(ctx)
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(os.Stdin)
	const prompt = "ais> " // D-01: branded prompt

	for {
		_, _ = fmt.Fprint(os.Stdout, prompt)

		if !scanner.Scan() {
			// Ctrl+D (EOF) — clean exit (D-02)
			_, _ = fmt.Fprintln(os.Stdout)
			break
		}

		input := strings.TrimSpace(scanner.Text())
		if input == "" {
			// Whitespace-only input is treated as empty — re-prompt without API call (D-11)
			continue
		}

		// Animated spinner while waiting for API response (D-08, D-09)
		s := spinner.New(spinner.CharSets[14], 80*time.Millisecond)
		s.Suffix = " Thinking..."
		s.Start()

		resp, err := client.Ask(ctx, input)
		s.Stop()
		if err != nil {
			// Print error but don't exit the REPL — allow retry (D-09)
			fmt.Fprintf(os.Stderr, "error: %v\n", classifyAPIError(err))
			continue
		}

		// Render response (OUT-01, D-13)
		render.Markdown(gemini.ResponseText(resp))

		// Print source citations (OUT-02, SRCH-02, D-05, D-06)
		if showRefs {
			render.Sources(gemini.ExtractSources(resp))
		}
	}

	if err := scanner.Err(); err != nil && !errors.Is(err, io.EOF) {
		return fmt.Errorf("reading input: %w", err)
	}
	return nil
}
