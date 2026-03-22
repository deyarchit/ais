# ais — AI Search Assistant

## What This Is

A Go CLI tool that wraps the Gemini API with always-on Google Search grounding, usable as both a one-shot search command (`ais -q "query"`) and a persistent interactive chat client (`ais`). Designed to replace a daily search workflow with AI-synthesized answers that cite their sources.

## Core Value

Every query returns a grounded, source-cited answer — either as a quick one-liner or a full conversation — without leaving the terminal.

## Requirements

### Validated

(None yet — ship to validate)

### Active

- [ ] One-shot mode: `ais -q "query"` returns answer and exits
- [ ] Interactive chat mode: `ais` (no args) opens a REPL with full multi-turn conversation context
- [ ] Google Search grounding always enabled on every query
- [ ] Responses rendered as markdown in the terminal (e.g., via glamour)
- [ ] Grounding sources/URLs listed after each response
- [ ] GEMINI_API_KEY env var used for authentication
- [ ] Streaming output — response streams to terminal as it arrives
- [ ] Graceful error handling: missing API key, network failure, API errors

### Out of Scope

- Config file — env var only, keeps it simple for daily use
- Multiple AI providers — Gemini-only for now
- Persisting conversation history to disk — sessions are ephemeral
- Web UI — terminal-only

## Context

- Go 1.25.7 project with strict golangci-lint config (slog for logging, errors must be wrapped)
- Existing scaffold has a `cmd/server/` pattern in Makefile — will pivot to `cmd/ais/` for the CLI binary
- Gemini Go SDK: `google.golang.org/genai`
- Markdown rendering: `github.com/charmbracelet/glamour` (standard Go CLI choice)
- The `ais` module name already matches the tool name — no rename needed

## Constraints

- **Tech stack**: Go only — no other runtimes
- **Auth**: GEMINI_API_KEY env var — no OAuth, no config files
- **Linting**: Must pass existing golangci-lint rules (slog for logs, fmt.Print* forbidden for log output, errors wrapped)
- **Quality**: Usable daily — error messages must be actionable, not cryptic

## Key Decisions

| Decision | Rationale | Outcome |
|----------|-----------|---------|
| Flag-based mode switching (`-q`) | Predictable: no args = chat, `-q` = one-shot | — Pending |
| Always-on search grounding | Simplifies UX — no per-query toggle needed | — Pending |
| Streaming output | Better UX for long AI responses | — Pending |
| Glamour for markdown rendering | De-facto standard for Go CLIs, handles terminal width | — Pending |
| Pivot from `cmd/server` to `cmd/ais` | Tool is a CLI, not a server | — Pending |

## Evolution

This document evolves at phase transitions and milestone boundaries.

**After each phase transition** (via `/gsd:transition`):
1. Requirements invalidated? → Move to Out of Scope with reason
2. Requirements validated? → Move to Validated with phase reference
3. New requirements emerged? → Add to Active
4. Decisions to log? → Add to Key Decisions
5. "What This Is" still accurate? → Update if drifted

**After each milestone** (via `/gsd:complete-milestone`):
1. Full review of all sections
2. Core Value check — still the right priority?
3. Audit Out of Scope — reasons still valid?
4. Update Context with current state

---
*Last updated: 2026-03-22 after initialization*
