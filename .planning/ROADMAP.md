# Roadmap: ais

## Overview

Build a Go CLI tool that wraps the Gemini API with always-on Google Search grounding. Phase 1 delivers a working tool — both query modes running end-to-end with rendered output and source citations. Phase 2 hardens it — error handling, lint compliance, and correct build targets — so it's reliable for daily use.

## Phases

**Phase Numbering:**
- Integer phases (1, 2, 3): Planned milestone work
- Decimal phases (2.1, 2.2): Urgent insertions (marked with INSERTED)

Decimal phases appear between their surrounding integers in numeric order.

- [x] **Phase 1: Working Tool** - Both modes functional end-to-end with Gemini grounding, markdown rendering, and source citations (completed 2026-03-22)
- [ ] **Phase 2: Production Ready** - Error handling hardened, lint clean, build system correct

## Phase Details

### Phase 1: Working Tool
**Goal**: A developer can query Gemini with Google Search grounding from the terminal in both one-shot and chat modes
**Depends on**: Nothing (first phase)
**Requirements**: MODE-01, MODE-02, MODE-03, SRCH-01, SRCH-02, OUT-01, OUT-02, CFG-01, TOOL-01
**Success Criteria** (what must be TRUE):
  1. `ais -q "query"` prints a markdown-rendered answer with source URLs, then exits cleanly
  2. `ais` (no args) opens an interactive REPL and accepts successive queries without restarting
  3. In chat mode, follow-up questions reference earlier answers (multi-turn context is preserved across turns)
  4. Responses are visually rendered — headers, bold, code blocks — not raw markdown strings
  5. Source URLs from Google Search grounding appear after each response body
  6. `make build` produces `./bin/ais`
**Plans**: 4 plans

Plans:
- [x] 01-01-PLAN.md — Foundation: dependencies, Gemini client with grounding, render helpers, Makefile fix
- [x] 01-02-PLAN.md — One-shot mode: `-q` flag, spinner, glamour render, source citations
- [x] 01-03-PLAN.md — Chat REPL: interactive loop, multi-turn history, `ais> ` prompt
- [x] 01-04-PLAN.md — End-to-end verification: build checks + live API human checkpoint

### Phase 2: Production Ready
**Goal**: The tool handles failure gracefully and meets production code quality standards
**Depends on**: Phase 1
**Requirements**: CFG-02, ERR-01, ERR-02, ERR-03, TOOL-02
**Success Criteria** (what must be TRUE):
  1. Running `ais` without `GEMINI_API_KEY` set shows an actionable message telling the user exactly which variable to set — no Go panic, no cryptic output
  2. A network failure or API error displays the failure reason and suggests a next step (e.g., check connectivity, verify key)
  3. Unknown flags or empty input show usage help, not a stack trace
  4. `make lint` passes with zero golangci-lint violations
**Plans**: TBD

## Progress

**Execution Order:**
Phases execute in numeric order: 1 → 2

| Phase | Plans Complete | Status | Completed |
|-------|----------------|--------|-----------|
| 1. Working Tool | 4/4 | Complete   | 2026-03-23 |
| 2. Production Ready | 0/? | Not started | - |
