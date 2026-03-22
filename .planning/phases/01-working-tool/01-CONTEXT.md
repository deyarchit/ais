# Phase 1: Working Tool - Context

**Gathered:** 2026-03-22
**Status:** Ready for planning

<domain>
## Phase Boundary

A developer can query Gemini with Google Search grounding from the terminal in both one-shot and chat modes. Covers the core binary, both modes, markdown rendering, source citation output, and `make build`. Error handling polish and linting compliance are Phase 2.

</domain>

<decisions>
## Implementation Decisions

### Chat REPL UX
- **D-01:** Input prompt is `ais> ` (branded, makes it obvious you're in the REPL)
- **D-02:** Exit via Ctrl+C or Ctrl+D only — no `exit`/`quit` keywords
- **D-03:** No welcome header — drop straight into the `ais> ` prompt on start
- **D-04:** No special REPL commands (no `/clear`, no `/help`) — keep it simple

### Source citation format
- **D-05:** After each response body, print a blank line then `Sources:` label followed by a numbered list of bare URLs
- **D-06:** If the API returns no grounding URLs, print `Sources: none`
- **D-07:** No horizontal rule separator — just blank line + "Sources:" label is sufficient

### Loading / waiting UX
- **D-08:** Show an animated spinner while waiting for the API response (applies to both one-shot and chat mode)
- **D-09:** Consistent behaviour — spinner in both `-q` one-shot mode and interactive REPL

### Already-decided (from project context)
- **D-10:** One-shot mode: `ais -q "query"` — exits cleanly after printing response + sources
- **D-11:** Chat mode: `ais` with no args — multi-turn REPL, full conversation history passed to Gemini each turn
- **D-12:** Google Search grounding always enabled — no per-query toggle
- **D-13:** Responses rendered via `github.com/charmbracelet/glamour`
- **D-14:** Auth via `GEMINI_API_KEY` environment variable
- **D-15:** Entry point is `cmd/ais/` — Makefile target `build` produces `./bin/ais`
- **D-16:** Phase 1 is non-streaming — full response arrives before rendering (streaming is v2/OUT-03)

### Claude's Discretion
- Exact spinner library/implementation (bubbletea, a simple goroutine loop, or similar)
- Internal package structure under `cmd/ais/` and `internal/`
- Exact glamour style/theme (auto, dark, light)
- Error message text for missing API key (CFG-01 guidance — Phase 2 polishes this further)

</decisions>

<specifics>
## Specific Ideas

- REPL should feel like a shell command prompt — `ais> ` branding is enough, no extra decoration
- Sources block is informational, not interactive — bare numbered URLs are sufficient
- Spinner shows the tool is working; silence while waiting would feel broken

</specifics>

<canonical_refs>
## Canonical References

**Downstream agents MUST read these before planning or implementing.**

### Project requirements and scope
- `.planning/REQUIREMENTS.md` — Full requirement list; Phase 1 covers MODE-01, MODE-02, MODE-03, SRCH-01, SRCH-02, OUT-01, OUT-02, CFG-01, TOOL-01
- `.planning/ROADMAP.md` — Phase 1 success criteria (6 items) — executor must verify all pass

### Existing scaffold
- `Makefile` — Current build target (`cmd/server` → must be pivoted to `cmd/ais`); lint, fmt, test targets defined
- `go.mod` — Module name is `ais`; Go 1.25.7

No external specs or ADRs beyond the above — all requirements are captured in decisions and REQUIREMENTS.md.

</canonical_refs>

<code_context>
## Existing Code Insights

### Reusable Assets
- None yet — bare scaffold (only `go.mod` and `Makefile`)

### Established Patterns
- `Makefile` lint target runs `golangci-lint` — plans must not break this; slog for logging, errors wrapped (Phase 2 full lint pass, but Phase 1 code should not introduce violations)

### Integration Points
- `cmd/ais/main.go` is the new entry point (replacing `cmd/server/`)
- `Makefile` `build` target must be updated: `go build -o ./bin/ais ./cmd/ais`

</code_context>

<deferred>
## Deferred Ideas

- Streaming output — explicitly v2 (OUT-03); Phase 1 buffers full response then renders
- `/clear` REPL command to reset conversation history — not in Phase 1
- Stdin pipe support (`echo "query" | ais`) — v2 (INP-01)
- Visual chat prompt showing turn number — v2 (OUT-04)
- `-s` system prompt flag — v2 (CUST-01)

</deferred>

---

*Phase: 01-working-tool*
*Context gathered: 2026-03-22*
