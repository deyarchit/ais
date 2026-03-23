# Phase 2: Production Ready - Context

**Gathered:** 2026-03-23
**Status:** Ready for planning

<domain>
## Phase Boundary

The tool handles failure gracefully and meets production code quality standards. Covers error messages for missing API key, API/network failures, and bad input; usage help for unknown flags; and full golangci-lint compliance. No new features — hardening only.

</domain>

<decisions>
## Implementation Decisions

### Missing API key (CFG-02, ERR-01)
- **D-01:** Minimal message — just name the variable: `GEMINI_API_KEY is not set`
- **D-02:** Print to stderr, exit 1
- **D-03:** No copy-paste command, no link to API key creation page

### API/network error messages (ERR-02)
- **D-04:** Smart append — classify the error and append a category-specific suggestion
- **D-05:** Auth errors (403 / "API key not valid") → append `"verify your API key is valid and has not expired"`
- **D-06:** Quota/rate limit (429 / "Resource exhausted") → append `"you have exceeded your API quota — wait before retrying"`
- **D-07:** Network/timeout errors (connection refused, deadline exceeded) → append `"check your internet connection"`
- **D-08:** Bad request (400) and all other unclassified errors → show raw SDK error, no suggestion appended
- **D-09:** All errors print to stderr only

### Bad input handling (ERR-03)
- **D-10:** `-q ""` or `-q "   "` (empty or whitespace-only query) → print `"error: query cannot be empty"` to stderr + show usage, exit 1
- **D-11:** REPL: whitespace-only input is silently re-prompted (same treatment as empty input — already handles empty, extend to whitespace)
- **D-12:** Unknown flags handled by `flag.Parse()` default behavior — no customization to usage text
- **D-13:** Extra positional args are ignored (flag default behavior)

### Lint compliance (TOOL-02)
- **D-14:** `forbidigo` rule `^fmt\.Print.*$` applies to `fmt.Print`, `fmt.Println`, `fmt.Printf` — replace all occurrences in `render.go` and `repl.go` with `fmt.Fprint*` variants targeting `os.Stdout` (these don't match the pattern and correctly express intent)
- **D-15:** All existing error strings must be lowercase with no trailing punctuation (revive `error-strings` rule)
- **D-16:** Code must pass `make lint` (`golangci-lint config verify && golangci-lint run ./...`) with zero violations

### Claude's Discretion
- Exact string matching logic for error classification (substring match, error code extraction, or sentinel errors)
- Whether to introduce a dedicated error-wrapping helper or inline the classification in each call site
- Order of nolint annotations if any are unavoidable

</decisions>

<specifics>
## Specific Ideas

- Error classification should be pragmatic — substring matching on the SDK error message is fine if the SDK doesn't expose typed errors
- The REPL already re-prompts on empty input (repl.go:43-45) — whitespace extension is a one-line `strings.TrimSpace` check that's already there; just confirm it covers whitespace

</specifics>

<canonical_refs>
## Canonical References

**Downstream agents MUST read these before planning or implementing.**

### Requirements
- `.planning/REQUIREMENTS.md` — Phase 2 covers CFG-02, ERR-01, ERR-02, ERR-03, TOOL-02
- `.planning/ROADMAP.md` — Phase 2 success criteria (4 items) — executor must verify all pass

### Lint rules
- `.golangci.yml` — Full linter config; key rules for this phase: `forbidigo` (fmt.Print* forbidden), `revive error-strings` (lowercase, no punctuation), `errorlint` (%w wrapping required), `revive deep-exit` (os.Exit in main only)

### Existing code to modify
- `cmd/ais/main.go` — Entry point; flag parsing, one-shot dispatch, error printing
- `internal/gemini/client.go` — `NewClient` (missing key check), `Ask` (API error surface point)
- `internal/render/render.go` — Uses `fmt.Print*` throughout — lint violations to fix
- `internal/repl/repl.go` — Uses `fmt.Print(prompt)` — lint violation; whitespace input check

</canonical_refs>

<code_context>
## Existing Code Insights

### Reusable Assets
- `gemini.NewClient`: already checks for missing API key and returns an error — needs message polish only (D-01)
- `gemini.Client.Ask`: returns SDK errors unwrapped — classification logic hooks in at the call sites in `main.go` and `repl.go`
- `repl.go:43-45`: existing empty-input guard (`if input == "" { continue }`) — extend with `strings.TrimSpace` (already called on line 41, so `input` is already trimmed — the guard is already correct for whitespace)

### Established Patterns
- All error paths use `fmt.Errorf("context: %w", err)` wrapping — maintain this
- `os.Exit(1)` only in `main()` — `deep-exit` rule enforced

### Integration Points
- Error classification wraps at the point where `client.Ask` errors are handled — `main.go:runOneShot` and `repl.go` loop
- `fmt.Print*` → `fmt.Fprint*(os.Stdout, ...)` replacements are mechanical, no logic changes

</code_context>

<deferred>
## Deferred Ideas

None — discussion stayed within phase scope.

</deferred>

---

*Phase: 02-production-ready*
*Context gathered: 2026-03-23*
