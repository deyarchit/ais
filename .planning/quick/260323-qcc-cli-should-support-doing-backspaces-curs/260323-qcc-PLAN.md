---
phase: quick
plan: 260323-qcc
type: execute
wave: 1
depends_on: []
files_modified:
  - internal/repl/repl.go
  - go.mod
  - go.sum
autonomous: true
requirements: []

must_haves:
  truths:
    - "Backspace erases the previous character in the prompt (does not echo ^H)"
    - "Left/right arrow keys move the cursor within the current input line"
    - "Up/down arrow keys cycle through previous inputs within the session"
    - "Ctrl+A moves cursor to start of line; Ctrl+E moves to end"
    - "Ctrl+C discards current input and re-prompts; Ctrl+D exits cleanly"
  artifacts:
    - path: "internal/repl/repl.go"
      provides: "REPL using readline instead of bufio.Scanner"
      contains: "chzyer/readline"
  key_links:
    - from: "internal/repl/repl.go"
      to: "github.com/chzyer/readline"
      via: "readline.New / rl.Readline()"
---

<objective>
Replace the raw `bufio.Scanner` input loop in the REPL with `github.com/chzyer/readline` so the prompt behaves like a standard terminal line editor: backspace works, arrow keys move the cursor, up/down recalls history, and Ctrl+A/E jump to line start/end.

Purpose: The current implementation reads bytes directly from stdin, so the terminal never interprets control sequences. Users see `^H` instead of an erased character and `^[[D` instead of cursor movement. This is the root cause of every "backspace / cursor skip" complaint.

Output: Updated `internal/repl/repl.go` using `readline.New`, updated `go.mod` / `go.sum`.
</objective>

<execution_context>
@~/.claude/get-shit-done/workflows/execute-plan.md
@~/.claude/get-shit-done/templates/summary.md
</execution_context>

<context>
@.planning/STATE.md
@internal/repl/repl.go
@cmd/ais/main.go
@go.mod
</context>

<tasks>

<task type="auto">
  <name>Task 1: Add readline dependency</name>
  <files>go.mod, go.sum</files>
  <action>
    Run:

      go get github.com/chzyer/readline@latest

    This adds the readline library to go.mod and populates go.sum. No source changes yet.
    Confirm the module appears in go.mod under `require`.
  </action>
  <verify>
    <automated>grep "chzyer/readline" /Users/deyarchit/Projects/ai/ais/go.mod</automated>
  </verify>
  <done>`github.com/chzyer/readline` appears in go.mod require block.</done>
</task>

<task type="auto">
  <name>Task 2: Replace bufio.Scanner with readline in repl.go</name>
  <files>internal/repl/repl.go</files>
  <action>
    Rewrite the input loop in `internal/repl/repl.go`. Key changes:

    1. Remove `bufio` import; add `"github.com/chzyer/readline"`.

    2. Before the loop, construct the readline instance:

       ```go
       rl, err := readline.New(prompt)
       if err != nil {
           return fmt.Errorf("readline init: %w", err)
       }
       defer rl.Close()
       ```

       The `prompt` const (`"ais> "`) is passed directly to `readline.New` — readline prints
       it and handles redrawing after spinner output.

    3. Replace the scanner loop body with:

       ```go
       for {
           line, err := rl.Readline()
           if errors.Is(err, readline.ErrInterrupt) {
               // Ctrl+C — discard current input, re-prompt
               continue
           }
           if errors.Is(err, io.EOF) {
               // Ctrl+D — clean exit (D-02)
               _, _ = fmt.Fprintln(os.Stdout)
               break
           }
           if err != nil {
               return fmt.Errorf("reading input: %w", err)
           }

           input := strings.TrimSpace(line)
           if input == "" {
               continue
           }
           // ... spinner, client.Ask, render — unchanged
       }
       return nil
       ```

    4. Remove the old `_, _ = fmt.Fprint(os.Stdout, prompt)` call and the post-loop
       `scanner.Err()` check — both are replaced by readline semantics above.

    5. Keep all other logic (spinner, classifyAPIError, render.Markdown, render.Sources,
       showRefs) exactly as-is.

    Design notes:
    - `readline.ErrInterrupt` is the signal for Ctrl+C; do NOT exit on it.
    - `io.EOF` is the signal for Ctrl+D; exit cleanly.
    - History within the session is automatic (readline maintains an in-memory ring).
    - Persistent history to disk is NOT needed — keep it session-only to avoid
      leaking queries to disk by default.
    - Do NOT configure a history file path (`readline.Config.HistoryFile`) for this task.
  </action>
  <verify>
    <automated>cd /Users/deyarchit/Projects/ai/ais && go build ./...</automated>
  </verify>
  <done>
    `go build ./...` completes with no errors.
    Manual smoke-test: run `make build && ./bin/ais`, type a few characters, press backspace
    — the character is erased rather than echoing `^H`. Press up-arrow — previous input
    is recalled. Ctrl+C clears the line. Ctrl+D exits.
  </done>
</task>

</tasks>

<verification>
After both tasks:

1. `go build ./...` — must pass with zero errors or warnings.
2. `go vet ./...` — must pass.
3. Run REPL manually (`./bin/ais`), verify:
   - Backspace erases characters (no `^H` echo)
   - Left/right arrows reposition the cursor
   - Up arrow recalls the previous input
   - Ctrl+C clears current input, re-prompts
   - Ctrl+D exits the program cleanly
</verification>

<success_criteria>
Terminal line editing works correctly in the REPL: backspace erases, arrow keys move the cursor, up/down cycles history, Ctrl+A/E jump to line boundaries, Ctrl+C re-prompts, Ctrl+D exits. `go build ./...` and `go vet ./...` both pass with no errors.
</success_criteria>

<output>
After completion, create `.planning/quick/260323-qcc-cli-should-support-doing-backspaces-curs/260323-qcc-SUMMARY.md` summarising what was changed and any decisions made.
</output>
