---
phase: quick-260428-egn
plan: 01
type: execute
wave: 1
depends_on: []
files_modified:
  - internal/gemini/client.go
  - go.mod
  - go.sum
autonomous: true
requirements: [260428-egn]
must_haves:
  truths:
    - "503 UNAVAILABLE errors are retried automatically with exponential backoff instead of surfacing to user"
    - "429 rate-limit errors are retried automatically with exponential backoff instead of surfacing to user"
    - "Each retry attempt prints a message to stderr so the user knows work is in progress"
    - "Other errors (4xx, 5xx non-retryable) surface immediately without retry"
  artifacts:
    - path: "internal/gemini/client.go"
      provides: "Ask() with exponential backoff retry wrapping c.chat.Send()"
      contains: "backoff.Retry"
  key_links:
    - from: "internal/gemini/client.go"
      to: "github.com/cenkalti/backoff/v5"
      via: "backoff.Retry with WithBackOff and WithNotify"
      pattern: "backoff\\.Retry"
---

<objective>
Add exponential backoff retry logic to the Gemini client's Ask() method so that transient 503 UNAVAILABLE and 429 rate-limit errors are automatically retried instead of propagated to the caller.

Purpose: Gemini occasionally returns 503/429 transiently; silently retrying with backoff improves reliability without burdening the user.
Output: Modified internal/gemini/client.go with retry logic, go.mod/go.sum updated with cenkalti/backoff/v5.
</objective>

<execution_context>
@~/.claude/get-shit-done/workflows/execute-plan.md
@~/.claude/get-shit-done/templates/summary.md
</execution_context>

<context>
@.planning/STATE.md
@.planning/quick/260428-egn-add-exponential-backoff-retry-logic-for-/260428-egn-CONTEXT.md

<interfaces>
<!-- Current Ask() in internal/gemini/client.go -->
```go
// Ask sends a message and returns the full text response and grounding metadata.
func (c *Client) Ask(ctx context.Context, prompt string) (*genai.GenerateContentResponse, error) {
    resp, err := c.chat.Send(ctx, genai.NewPartFromText(prompt))
    if err != nil {
        return nil, fmt.Errorf("send message: %w", err)
    }
    return resp, nil
}
```

<!-- backoff/v5 API (key subset) -->
```go
// github.com/cenkalti/backoff/v5
func Retry[T any](ctx context.Context, operation Operation[T], opts ...RetryOption) (T, error)
func WithBackOff(b BackOff) RetryOption
func WithNotify(n Notify) RetryOption   // type Notify func(error, time.Duration)
func Permanent(err error) error          // stop retrying immediately

type ExponentialBackOff struct {
    InitialInterval     time.Duration  // default 500ms
    RandomizationFactor float64        // default 0.5 — built-in jitter
    Multiplier          float64        // default 1.5
    MaxInterval         time.Duration  // set to 30s per user decision
}
func NewExponentialBackOff() *ExponentialBackOff
```
</interfaces>
</context>

<tasks>

<task type="auto">
  <name>Task 1: Add cenkalti/backoff/v5 dependency</name>
  <files>go.mod, go.sum</files>
  <action>
Run `go get github.com/cenkalti/backoff/v5` from the project root to add the dependency and update go.mod and go.sum. Verify the module appears in go.mod after running the command.
  </action>
  <verify>
    <automated>grep "cenkalti/backoff" /Users/deyarchit/Projects/ai/ais/go.mod</automated>
  </verify>
  <done>go.mod contains github.com/cenkalti/backoff/v5 entry</done>
</task>

<task type="auto" tdd="false">
  <name>Task 2: Implement exponential backoff retry in Ask()</name>
  <files>internal/gemini/client.go</files>
  <action>
Modify Ask() in internal/gemini/client.go to wrap c.chat.Send() with exponential backoff retry logic.

Implementation steps:

1. Add import: `"github.com/cenkalti/backoff/v5"` (and `"fmt"`, `"strings"` already present via package — confirm imports).

2. Add a package-level helper function `isRetryable(err error) bool` that returns true if the error string contains any of: "503", "UNAVAILABLE", "429", "Resource exhausted". Use strings.Contains on err.Error().

3. Rewrite Ask() as follows:

```go
func (c *Client) Ask(ctx context.Context, prompt string) (*genai.GenerateContentResponse, error) {
    bo := backoff.NewExponentialBackOff()
    bo.MaxInterval = 30 * time.Second

    attempt := 0
    resp, err := backoff.Retry(ctx, func() (*genai.GenerateContentResponse, error) {
        r, e := c.chat.Send(ctx, genai.NewPartFromText(prompt))
        if e != nil {
            if isRetryable(e) {
                return nil, e // retryable — backoff will retry
            }
            return nil, backoff.Permanent(e) // non-retryable — stop immediately
        }
        return r, nil
    },
        backoff.WithBackOff(bo),
        backoff.WithNotify(func(e error, wait time.Duration) {
            attempt++
            fmt.Fprintf(os.Stderr, "retrying after %.1fs (attempt %d): %v\n", wait.Seconds(), attempt, e)
        }),
    )
    if err != nil {
        return nil, fmt.Errorf("send message: %w", err)
    }
    return resp, nil
}
```

4. Add `"time"` to imports if not already present (it is not in current imports — add it).

5. The `isRetryable` function:

```go
func isRetryable(err error) bool {
    msg := err.Error()
    return strings.Contains(msg, "503") ||
        strings.Contains(msg, "UNAVAILABLE") ||
        strings.Contains(msg, "429") ||
        strings.Contains(msg, "Resource exhausted")
}
```

6. Add `"strings"` to imports (not currently in client.go — add it).

Note: `os` is already imported. `fmt` is already imported. The backoff library's built-in RandomizationFactor (0.5 default) provides the jitter — no additional jitter configuration needed (per user decision: use library defaults, cap MaxInterval at 30s).
  </action>
  <verify>
    <automated>cd /Users/deyarchit/Projects/ai/ais && go build ./...</automated>
  </verify>
  <done>
    - go build ./... succeeds with no errors
    - internal/gemini/client.go contains isRetryable() function
    - Ask() wraps c.chat.Send() inside backoff.Retry
    - 503/UNAVAILABLE/429/Resource exhausted errors are retried; all others hit backoff.Permanent
    - Retry notification prints to stderr with wait duration and attempt count
  </done>
</task>

</tasks>

<verification>
After both tasks complete:
- `go build ./...` passes — no compilation errors
- `grep -n "backoff.Retry\|isRetryable\|Permanent" internal/gemini/client.go` shows the retry logic present
- `grep "cenkalti/backoff" go.mod` confirms dependency added
</verification>

<success_criteria>
- go.mod includes github.com/cenkalti/backoff/v5
- Ask() retries on 503/UNAVAILABLE/429/Resource exhausted errors with up to 30s max interval between attempts (with jitter from RandomizationFactor=0.5 default)
- Each retry prints to stderr: `retrying after Xs (attempt N): <error>`
- Non-retryable errors (403, connection refused, etc.) surface immediately via backoff.Permanent
- go build ./... succeeds
</success_criteria>

<output>
After completion, create `.planning/quick/260428-egn-add-exponential-backoff-retry-logic-for-/260428-egn-SUMMARY.md`
</output>
