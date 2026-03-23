# Phase 3: Add Additional Flags to Control Various Params for Gemini API - Context

**Gathered:** 2026-03-23
**Status:** Ready for planning

<domain>
## Phase Boundary

Add CLI flags to expose Gemini API parameters that users currently cannot control. The tool runs queries with hardcoded model config (gemini-2.5-flash, no temperature set, default thinking). This phase lets users tune temperature and thinking budget at invocation time. No new modes, no config files, no env vars — flags only.

</domain>

<decisions>
## Implementation Decisions

### Which params to expose
- **D-01:** Two new flags only: `--temperature` and `--thinking-budget`
- **D-02:** No model selection flag, no max-tokens, no system prompt in this phase

### Temperature flag
- **D-03:** Flag name: `--temperature` (long only, no short alias — consistent with existing `--show-refs`)
- **D-04:** Range: 0.0–2.0 (matches Gemini SDK range)
- **D-05:** Default: 0.7 (balanced — slightly lower than SDK default of 1.0, appropriate for a search tool)
- **D-06:** If flag not provided, pass 0.7 explicitly to the SDK

### Thinking budget flag
- **D-07:** Flag name: `--thinking-budget` with named presets: `none`, `low`, `medium`, `high`, `auto`
- **D-08:** Default: `auto` (SDK decides thinking depth dynamically — Gemini 2.5 uses reasoning when needed)
- **D-09:** Preset → token mapping:
  - `none` → 0 (disabled)
  - `low` → 1024
  - `medium` → 8192
  - `high` → 24576
  - `auto` → -1 (SDK dynamic)
- **D-10:** Invalid preset value → show error with valid options list and exit 1

### Scope — both modes
- **D-11:** Both `--temperature` and `--thinking-budget` apply to both one-shot (`-q`) and REPL mode
- **D-12:** Flags are parsed in `main()` and passed through to `gemini.NewClient()` — the client config applies to all queries in the session
- **D-13:** No mid-session param changes in REPL — flags are set at startup and fixed for the session

### Persistence
- **D-14:** Flags only — no env var fallbacks (e.g., no `AIS_TEMPERATURE`)
- **D-15:** No config file support in this phase

### Claude's Discretion
- How `genai.GenerateContentConfig` fields are set for each preset (ThinkingConfig struct details)
- Whether to validate temperature range client-side or let the SDK error
- Exact error message wording for invalid `--thinking-budget` values

</decisions>

<specifics>
## Specific Ideas

- Temperature default of 0.7 (not SDK default 1.0) — better for factual/search queries
- Thinking budget uses named presets (`low`, `medium`, `high`) rather than raw token counts — friendlier UX
- `auto` as default for thinking budget — Gemini 2.5 Flash with dynamic thinking is the right out-of-box behavior

</specifics>

<canonical_refs>
## Canonical References

**Downstream agents MUST read these before planning or implementing.**

### Current implementation
- `cmd/ais/main.go` — Flag parsing in main(), runOneShot() signature, how showRefs is threaded through
- `internal/gemini/client.go` — NewClient() signature, GenerateContentConfig construction, model name hardcoding
- `internal/repl/repl.go` — Run() signature, how showRefs is passed in

### No external specs
No ADRs or feature docs — requirements fully captured in decisions above.

</canonical_refs>

<code_context>
## Existing Code Insights

### Reusable Assets
- `flag.String`, `flag.Bool`, `flag.Float64` patterns in `cmd/ais/main.go` — add new flags in same block
- `gemini.NewClient(ctx)` — needs to accept config params (temperature, thinkingBudget); signature change required
- `genai.GenerateContentConfig` — already constructed in `NewClient`; extend it with `Temperature` and `ThinkingConfig`

### Established Patterns
- Long-form flags only (no short aliases) — consistent with `--show-refs`
- Flags parsed in `main()`, passed as args to `repl.Run()` and `runOneShot()` — new params follow the same threading pattern
- `gemini.NewClient(ctx)` currently takes no config — will need a new signature (options struct or explicit params)

### Integration Points
- `gemini.NewClient()` is called in both `main.go` (one-shot) and `repl.go` (REPL) — signature change affects both call sites
- `genai.GenerateContentConfig.Temperature` — float32 field, set directly
- `genai.GenerateContentConfig.ThinkingConfig` — set `ThinkingBudget` (int32) on the config

</code_context>

<deferred>
## Deferred Ideas

- Model selection flag (`--model gemini-2.5-pro`) — not in this phase
- System prompt flag (`-s`) — already in REQUIREMENTS.md as CUST-01 (v2)
- Max output tokens flag — not requested for this phase
- Env var fallbacks (`AIS_TEMPERATURE`) — explicitly deferred
- Mid-session `/set` commands in REPL — explicitly deferred

</deferred>

---

*Phase: 03-add-additional-flags-to-control-various-params-for-gemini-api*
*Context gathered: 2026-03-23*
