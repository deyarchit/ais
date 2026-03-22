package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/briandowns/spinner"

	"ais/internal/gemini"
	"ais/internal/render"
	"ais/internal/repl"
)

func main() {
	query := flag.String("q", "", "Query to ask in one-shot mode")
	flag.Parse()

	if *query != "" {
		if err := runOneShot(context.Background(), *query); err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
		return
	}

	// Chat REPL mode (MODE-02, D-11)
	if err := repl.Run(context.Background()); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

// runOneShot executes a single query and exits. Creates a fresh Client so no
// prior history is attached (one-shot = stateless, per D-10).
func runOneShot(ctx context.Context, query string) error {
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
		return fmt.Errorf("query failed: %w", err)
	}

	// Render markdown response (OUT-01, D-13)
	render.Markdown(gemini.ResponseText(resp))

	// Print source citations (OUT-02, SRCH-02, D-05, D-06)
	render.Sources(gemini.ExtractSources(resp))

	return nil
}
