---
phase: quick
plan: 260323-qcc
subsystem: repl
tags: [readline, terminal, ux, input-handling]
dependency_graph:
  requires: []
  provides: [terminal-line-editing]
  affects: [internal/repl/repl.go]
tech_stack:
  added: [github.com/chzyer/readline v1.5.1]
  patterns: [readline-based REPL input loop]
key_files:
  created: []
  modified:
    - internal/repl/repl.go
    - go.mod
    - go.sum
decisions:
  - "Use chzyer/readline in-memory history only — no HistoryFile to avoid leaking queries to disk"
  - "ErrInterrupt (Ctrl+C) re-prompts without exit; io.EOF (Ctrl+D) exits cleanly"
metrics:
  duration: "~5 min"
  completed_date: "2026-03-23"
  tasks_completed: 2
  files_changed: 3
---

# Quick Task 260323-qcc: CLI should support backspace/cursor editing

**One-liner:** Replaced bufio.Scanner with chzyer/readline to give the REPL proper terminal line-editing (backspace, arrow keys, history, Ctrl+A/E).

## What Was Done

The REPL previously used `bufio.NewScanner(os.Stdin)` for input, which operates in cooked terminal mode without line-editing support. Users saw `^H` when pressing backspace and `^[[D` / `^[[C` escape sequences when pressing arrow keys, because no readline layer was translating those keystrokes.

### Task 1: Add readline dependency

Added `github.com/chzyer/readline v1.5.1` via `go get`. The module was recorded in `go.mod` and `go.sum`.

**Commit:** `5d05131`

### Task 2: Replace bufio.Scanner with readline in repl.go

Rewrote the input loop in `internal/repl/repl.go`:

- Removed `bufio` import; added `github.com/chzyer/readline`.
- Before the loop: `rl, err := readline.New(prompt)` — readline takes over prompt printing and puts the terminal into raw mode for keystroke handling.
- Loop body: `line, err := rl.Readline()` replaces `scanner.Scan()` + `scanner.Text()`.
- `readline.ErrInterrupt` (Ctrl+C): continue (re-prompt, do not exit).
- `io.EOF` (Ctrl+D): print newline and break (clean exit).
- Any other error: return it wrapped.
- Removed the old `fmt.Fprint(os.Stdout, prompt)` call and `scanner.Err()` check — both superseded by readline semantics.
- All other logic (spinner, classifyAPIError, render.Markdown, render.Sources, showRefs) is unchanged.

**Commit:** `23b6393`

## Terminal behaviours now working

| Action | Before | After |
|--------|--------|-------|
| Backspace | echoed `^H` | erases previous character |
| Left/right arrow | echoed `^[[D` / `^[[C` | moves cursor within line |
| Up/down arrow | no effect | recalls previous inputs |
| Ctrl+A / Ctrl+E | no effect | jumps to line start/end |
| Ctrl+C | SIGINT / exit | clears current input, re-prompts |
| Ctrl+D | EOF, sometimes garbled | exits cleanly with newline |

## Decisions Made

1. **No persistent history file** — `readline.Config.HistoryFile` was intentionally left unconfigured. History lives in memory for the session only, avoiding leaking user queries to disk.
2. **ErrInterrupt vs io.EOF split** — Ctrl+C maps to `readline.ErrInterrupt` (continue), Ctrl+D maps to `io.EOF` (break). This matches standard UNIX shell semantics.

## Deviations from Plan

None — plan executed exactly as written.

## Known Stubs

None.

## Verification

- `go build ./...` — passed with zero errors.
- `go vet ./...` — passed with zero warnings.

## Self-Check: PASSED

- `/Users/deyarchit/Projects/ai/ais/internal/repl/repl.go` — modified, contains `readline.New`.
- `/Users/deyarchit/Projects/ai/ais/go.mod` — contains `github.com/chzyer/readline v1.5.1`.
- Commits `5d05131` and `23b6393` present in git log.
