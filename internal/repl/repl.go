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

// Run starts the interactive REPL loop. It blocks until the user sends
// Ctrl+C (SIGINT) or Ctrl+D (EOF on stdin). No welcome header is printed (D-03).
// The same gemini.Client is reused across turns to preserve ChatSession history (D-11, MODE-03).
func Run(ctx context.Context) error {
	client, err := gemini.NewClient(ctx)
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(os.Stdin)
	const prompt = "ais> " // D-01: branded prompt

	for {
		fmt.Print(prompt)

		if !scanner.Scan() {
			// Ctrl+D (EOF) — clean exit (D-02)
			fmt.Println()
			break
		}

		input := strings.TrimSpace(scanner.Text())
		if input == "" {
			// Empty input — re-prompt without sending to API
			continue
		}

		// Animated spinner while waiting for API response (D-08, D-09)
		s := spinner.New(spinner.CharSets[14], 80*time.Millisecond)
		s.Suffix = " Thinking..."
		s.Start()

		resp, err := client.Ask(ctx, input)
		s.Stop()
		if err != nil {
			// Print error but don't exit the REPL — allow retry
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			continue
		}

		// Render response (OUT-01, D-13)
		render.Markdown(gemini.ResponseText(resp))

		// Print source citations (OUT-02, SRCH-02, D-05, D-06)
		render.Sources(gemini.ExtractSources(resp))
	}

	if err := scanner.Err(); err != nil && !errors.Is(err, io.EOF) {
		return fmt.Errorf("reading input: %w", err)
	}
	return nil
}
