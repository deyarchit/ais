# ais — AI Search Assistant

## What This Is

A Go CLI tool that wraps the Gemini API with always-on Google Search grounding, usable as both a one-shot search command (`ais -q "query"`) and a persistent interactive chat client (`ais`). Designed to replace a daily search workflow with AI-synthesized answers that cite their sources.

## Core Value

Every query returns a grounded, source-cited answer — either as a quick one-liner or a full conversation — without leaving the terminal.

## Current State (v1.0 — Shipped 2026-03-23)

The tool is feature-complete and production-ready for daily use.

**What's working:**
- `ais -q "query"` — one-shot mode with spinner, markdown rendering, source citations
- `ais` — interactive REPL with multi-turn conversation history
- Google Search grounding always-on on every query
- `--temperature` flag (float [0.0, 2.0], default 0.7)
- `--thinking-budget` flag (presets: none/low/medium/high/auto, default auto)
- Actionable error messages for missing key, network, auth/quota failures
- `make build` + `make lint` pass clean

**Tech stack:**
- Go, `google.golang.org/genai` (Gemini SDK), `glamour` (markdown), `readline` (REPL input), `spinner`
- `cmd/ais/main.go` → `internal/gemini/`, `internal/render/`, `internal/repl/`

## Next Milestone Goals (v2 — TBD)

Candidates (not yet scoped):
- Streaming output — response streams to terminal as it arrives
- Stdin pipe support — `echo "query" | ais` for scripting
- System prompt flag (`-s`) for persona/context injection
- Visual chat prompt showing turn number or context length

*Start with `/gsd:new-milestone` to scope and plan v2.*

## Constraints

- **Tech stack**: Go only — no other runtimes
- **Auth**: GEMINI_API_KEY env var — no OAuth, no config files
- **Linting**: Must pass existing golangci-lint rules (slog for logs, `fmt.Print*` forbidden, errors wrapped)
- **Quality**: Usable daily — error messages must be actionable, not cryptic

## Key Decisions (Locked)

| Decision | Rationale | Outcome |
|----------|-----------|---------|
| Flag-based mode switching (`-q`) | Predictable: no args = chat, `-q` = one-shot | Shipped v1.0 |
| Always-on search grounding | Simplifies UX — no per-query toggle needed | Shipped v1.0 |
| Glamour for markdown rendering | De-facto standard for Go CLIs, handles terminal width | Shipped v1.0 |
| Pivot from `cmd/server` to `cmd/ais` | Tool is a CLI, not a server | Shipped v1.0 |
| New `google.golang.org/genai` SDK | Deprecated SDK would require migration later | Shipped v1.0 |
| Duplicate `classifyAPIError` (not shared) | Avoided changing `repl.Run()` signature at Phase 2 | Technical debt — refactor if 3rd call site |
| Pre-validated `ClientConfig` struct | Clean boundary between CLI flags and API params | Shipped v1.0 |
| Streaming deferred | Scope control for v1.0 | Deferred to v2 |

## Out of Scope (Permanent)

| Feature | Reason |
|---------|--------|
| Config file | Env var is sufficient; config file adds complexity without value |
| Multiple AI providers | Gemini-only keeps the scope tight |
| Persistent conversation history | Sessions are ephemeral; disk storage is v2+ |
| Web UI / TUI | Terminal text output is the right interface |
| OAuth / complex auth | API key is the right auth model for a personal CLI tool |

---
*v1.0 archived: .planning/milestones/v1.0-ROADMAP.md*
*Last updated: 2026-03-23 — v1.0 milestone complete*
