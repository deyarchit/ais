---
phase: quick-260428-egn
plan: "01"
subsystem: internal/gemini
tags: [retry, backoff, resilience, gemini-client]
dependency_graph:
  requires: []
  provides: [exponential-backoff-retry-in-ask]
  affects: [internal/gemini/client.go]
tech_stack:
  added: [github.com/cenkalti/backoff/v5 v5.0.3]
  patterns: [exponential-backoff-with-jitter, permanent-error-wrapping, notify-callback]
key_files:
  created: []
  modified:
    - internal/gemini/client.go
    - go.mod
    - go.sum
decisions:
  - "Use library defaults for RandomizationFactor (0.5) and Multiplier (1.5); only cap MaxInterval at 30s"
  - "isRetryable matches on string substrings (503, UNAVAILABLE, 429, Resource exhausted) since SDK does not expose typed errors"
  - "Retry notify increments an attempt counter so message prints attempt N not the raw backoff attempt count"
metrics:
  duration: "~1 min"
  completed_date: "2026-04-28"
  tasks_completed: 2
  files_changed: 3
---

# Quick Task 260428-egn: Add Exponential Backoff Retry Logic for Gemini API Summary

**One-liner:** Exponential backoff retry wrapping Gemini's `chat.Send()` using cenkalti/backoff/v5 — auto-retries 503/429 transient errors with up to 30s max interval and stderr progress messages.

## Tasks Completed

| Task | Name | Commit | Files |
|------|------|--------|-------|
| 1 | Add cenkalti/backoff/v5 dependency | 08737fd | go.mod, go.sum |
| 2 | Implement exponential backoff retry in Ask() | 4c831d8 | internal/gemini/client.go |

## What Was Built

- **`isRetryable(err error) bool`** — package-level helper that returns true if the error message contains `503`, `UNAVAILABLE`, `429`, or `Resource exhausted`.
- **`Ask()` rewritten** — wraps `c.chat.Send()` inside `backoff.Retry` using `ExponentialBackOff` with `MaxInterval = 30s`. Retryable errors pass through; non-retryable errors are wrapped in `backoff.Permanent()` to stop immediately.
- **Stderr progress** — `WithNotify` callback prints `retrying after Xs (attempt N): <error>` to stderr on each retry so the user sees work is in progress.
- **New imports added to `client.go`:** `strings`, `time`, `github.com/cenkalti/backoff/v5`.

## Deviations from Plan

None - plan executed exactly as written.

## Known Stubs

None.

## Self-Check: PASSED

- `internal/gemini/client.go` — FOUND
- `go.mod` contains `cenkalti/backoff` — FOUND
- Commit `08737fd` — FOUND
- Commit `4c831d8` — FOUND
- `go build ./...` — PASSED
