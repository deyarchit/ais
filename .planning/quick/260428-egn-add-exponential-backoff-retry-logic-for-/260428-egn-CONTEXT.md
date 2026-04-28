---
quick_id: 260428-egn
description: Add exponential backoff retry logic for Gemini 503 UNAVAILABLE errors
gathered: 2026-04-28
status: Ready for planning
---

# Quick Task 260428-egn: Add exponential backoff retry logic for Gemini 503 UNAVAILABLE errors - Context

**Gathered:** 2026-04-28
**Status:** Ready for planning

<domain>
## Task Boundary

Add exponential backoff retry logic to the Gemini client so that transient 503 UNAVAILABLE and 429 rate-limit errors are automatically retried instead of surfaced as errors to the user. The retry logic lives in `internal/gemini/client.go` in the `Ask()` method.

</domain>

<decisions>
## Implementation Decisions

### Which errors to retry on
- Retry on **503 UNAVAILABLE** and **429 rate-limit** errors
- Do NOT retry on other error codes (4xx client errors, etc.)

### Retry library and parameters
- Use **`github.com/cenkalti/backoff/v5`** (user-specified library)
- Apply **jitter** to avoid thundering herd
- **Max wait** of 30 seconds between attempts
- Use the library's built-in exponential backoff with defaults, capped at 30s max interval

### User feedback during retries
- Print a retry message to **stderr** so the user knows it's working
- Format: something like `retrying after Xs (attempt N)...` printed during each retry

### Claude's Discretion
- Exact message wording for retry output
- Whether to cap total retry duration or number of attempts (use library defaults unless they conflict with 30s max wait constraint)

</decisions>

<specifics>
## Specific Ideas

- Library: https://pkg.go.dev/github.com/cenkalti/backoff/v5
- The error occurs at `c.chat.Send()` in `Ask()` — `internal/gemini/client.go:75`
- Error detection needs to parse the gRPC status code from the error (Status: UNAVAILABLE / 503, or 429)

</specifics>

<canonical_refs>
## Canonical References

- https://pkg.go.dev/github.com/cenkalti/backoff/v5 — the backoff library to use
</canonical_refs>
