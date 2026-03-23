---
phase: 02-production-ready
plan: 01
subsystem: api
tags: [go, error-handling, gemini, cli]

# Dependency graph
requires:
  - phase: 01-working-tool
    provides: gemini.Client.Ask and the REPL/one-shot entry points that surface API errors
provides:
  - Actionable error messages for missing API key, API auth/quota/network failures, and empty query input
  - classifyAPIError helper present in both cmd/ais/main.go and internal/repl/repl.go
affects: [02-02-lint-compliance]

# Tech tracking
tech-stack:
  added: []
  patterns: [substring-based SDK error classification via classifyAPIError helper]

key-files:
  created: []
  modified:
    - internal/gemini/client.go
    - cmd/ais/main.go
    - internal/repl/repl.go

key-decisions:
  - "Duplicate classifyAPIError in both main.go and repl.go rather than change Run() signature — avoids breaking the caller interface while keeping both call sites identical"
  - "Substring matching on SDK error message strings for error classification — SDK does not expose typed sentinel errors"

patterns-established:
  - "classifyAPIError pattern: wrap SDK errors with category-specific suggestions before printing to stderr"
  - "Empty-input guard: strings.TrimSpace check before branching to API call"

requirements-completed: [CFG-02, ERR-01, ERR-02, ERR-03]

# Metrics
duration: 4min
completed: 2026-03-23
---

# Phase 2 Plan 01: Error Handling Summary

**Actionable error messages for missing API key, auth/quota/network API failures, and empty/whitespace query input via classifyAPIError substring-matching helper**

## Performance

- **Duration:** 4 min
- **Started:** 2026-03-23T10:08:53Z
- **Completed:** 2026-03-23T10:13:08Z
- **Tasks:** 2
- **Files modified:** 3

## Accomplishments

- Fixed error-strings lint violation in client.go: `"GEMINI_API_KEY is not set"` (was "GEMINI_API_KEY environment variable is not set")
- Added classifyAPIError helper in main.go and repl.go with auth/quota/network error classification via substring matching
- Added empty/whitespace query guard in main() with `flag.Usage()` output and exit 1
- Applied classifyAPIError at both Ask error sites (runOneShot and REPL loop)
- Documented whitespace-only input guard in repl.go with D-11 reference

## Task Commits

Each task was committed atomically:

1. **Task 1: Fix missing-key error message and add classifyAPIError helper** - `8800ab3` (feat)
2. **Task 2: Apply classifyAPIError in the REPL and confirm whitespace guard** - `f463d1d` (feat)

## Files Created/Modified

- `internal/gemini/client.go` - Fixed error-strings violation: "GEMINI_API_KEY is not set"
- `cmd/ais/main.go` - Added classifyAPIError helper, empty/whitespace query guard, strings import, classifyAPIError call at Ask error site
- `internal/repl/repl.go` - Added classifyAPIError helper, applied at Ask error site, documented whitespace guard

## Decisions Made

- Duplicated classifyAPIError in both main.go and repl.go rather than refactoring Run() to accept a function parameter — avoids changing the Run signature and breaking callers while keeping both implementations identical (by plan design, D-04)
- Substring matching on SDK error messages is appropriate since the Gemini SDK does not expose typed sentinel errors

## Deviations from Plan

None - plan executed exactly as written.

## Issues Encountered

Pre-existing golangci-lint violations in render.go and repl.go (`fmt.Print*` forbidden by forbidigo rule) were observed but are out of scope for this plan. They are addressed in plan 02-02 (lint compliance). These were pre-existing before this plan's changes.

## User Setup Required

None - no external service configuration required.

## Next Phase Readiness

- Error handling complete for CFG-02, ERR-01, ERR-02, ERR-03
- Ready for plan 02-02: lint compliance (TOOL-02) — replace fmt.Print* with fmt.Fprint* variants in render.go and repl.go

## Self-Check: PASSED

- FOUND: internal/gemini/client.go
- FOUND: cmd/ais/main.go
- FOUND: internal/repl/repl.go
- FOUND: .planning/phases/02-production-ready/02-01-SUMMARY.md
- FOUND commit: 8800ab3 (Task 1)
- FOUND commit: f463d1d (Task 2)

---
*Phase: 02-production-ready*
*Completed: 2026-03-23*
