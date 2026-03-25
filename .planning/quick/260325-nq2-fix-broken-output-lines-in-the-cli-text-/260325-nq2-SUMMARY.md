---
phase: quick
plan: 260325-nq2
subsystem: render
tags: [cli, terminal, glamour, word-wrap, ux]
dependency_graph:
  requires: []
  provides: ["dynamic-terminal-width-word-wrap"]
  affects: ["internal/render/render.go"]
tech_stack:
  added: ["golang.org/x/term (now direct import, was indirect)"]
  patterns: ["term.GetSize syscall for dynamic TTY width detection"]
key_files:
  modified: ["internal/render/render.go"]
decisions:
  - "Query terminal width per render call (not cached) — cheap syscall, handles resize between renders"
  - "Fall back to 80 columns when stdout is not a TTY (piped/redirected output)"
metrics:
  duration: "5 minutes"
  completed: "2026-03-25"
  tasks: 1
  files: 1
---

# Quick Task 260325-nq2: Fix Broken Output Lines in the CLI Text Summary

**One-liner:** Dynamic terminal width word wrap via `term.GetSize` replacing hardcoded 120-column glamour wrap.

## What Was Done

Replaced `glamour.WithWordWrap(120)` with a dynamic width derived from `term.GetSize(int(os.Stdout.Fd()))` on every call to `Markdown()`. When stdout is not a TTY (piped or redirected), the width falls back to 80 columns.

## Tasks

| Task | Name | Commit | Files |
|------|------|--------|-------|
| 1 | Replace hardcoded word-wrap width with dynamic terminal width | 2c4536f | internal/render/render.go |

## Verification

- `go build ./...` passes
- `go vet ./internal/render/...` passes
- `grep -n "WithWordWrap(120)" internal/render/render.go` returns no matches
- `grep -n "term.GetSize" internal/render/render.go` returns a match (line 17)

## Deviations from Plan

None - plan executed exactly as written.

## Known Stubs

None.

## Self-Check: PASSED

- File exists: `/Users/deyarchit/Projects/ai/ais/internal/render/render.go` - FOUND
- Commit exists: `2c4536f` - FOUND
- `WithWordWrap(120)` removed - CONFIRMED
- `term.GetSize` present - CONFIRMED
