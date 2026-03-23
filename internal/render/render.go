package render

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/glamour"
)

// Markdown renders a markdown string to the terminal using glamour auto style.
// "auto" picks dark or light theme based on terminal background.
// Falls back to printing raw text if glamour fails.
func Markdown(text string) {
	r, err := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(120),
	)
	if err != nil {
		fmt.Fprintln(os.Stdout, text)
		return
	}
	out, err := r.Render(text)
	if err != nil {
		fmt.Fprintln(os.Stdout, text)
		return
	}
	fmt.Fprint(os.Stdout, out)
}

// Sources prints the source citation block after a response.
// Per D-05: blank line, then "Sources:" label, then numbered bare URLs.
// Per D-06: if urls is empty, prints "Sources: none".
// Per D-07: no horizontal rule separator.
func Sources(urls []string) {
	fmt.Fprintln(os.Stdout)
	if len(urls) == 0 {
		fmt.Fprintln(os.Stdout, "Sources: none")
		return
	}
	fmt.Fprintln(os.Stdout, "Sources:")
	for i, u := range urls {
		fmt.Fprintf(os.Stdout, "  %d. %s\n", i+1, strings.TrimSpace(u))
	}
}
