package main

import (
	"flag"
	"fmt"
)

func main() {
	query := flag.String("q", "", "Query to ask in one-shot mode")
	flag.Parse()

	if *query != "" {
		// One-shot mode — implemented in plan 01-02
		fmt.Println("one-shot mode: placeholder")
		return
	}

	// Chat REPL mode — implemented in plan 01-03
	fmt.Println("chat mode: placeholder")
}
