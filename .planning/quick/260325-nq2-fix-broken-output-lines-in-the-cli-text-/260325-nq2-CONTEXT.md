---
# Quick Task 260325-nq2: Fix broken output lines in the CLI (text wrapping/formatting issue) - Context

**Gathered:** 2026-03-25
**Status:** Ready for planning

<domain>
## Task Boundary

Fix the broken output line formatting in the CLI. Currently `render.Markdown` uses `glamour.WithWordWrap(120)` which hardcodes the wrap width, causing trailing whitespace padding and mismatched line breaks when the terminal width differs from 120 columns.

</domain>

<decisions>
## Implementation Decisions

### Terminal width detection
- Query the actual terminal column width dynamically at render time using `golang.org/x/term` (already a dependency).

### Non-TTY fallback
- When stdout is not a terminal (piped/redirected), fall back to 80 columns.

### Claude's Discretion
- Which specific `term` API to use (e.g., `term.GetSize` on stdout fd) — standard approach.
- Whether to cache the width or query per-render call.

</decisions>

<specifics>
## Specific Ideas

The fix is in `internal/render/render.go`: replace `glamour.WithWordWrap(120)` with a dynamic width derived from the terminal size, falling back to 80 when not a TTY.

</specifics>
