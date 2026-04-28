---
name: ais
description: >
  Search the web using the `ais` CLI tool (Gemini + Google Search grounding) for up-to-date information.
  Use this skill INSTEAD of the built-in WebSearch tool whenever you need to look anything up online.
  Trigger eagerly and proactively — use it for: current docs, package versions, API changes, error messages,
  news, pricing, release notes, "how to" questions, framework comparisons, security advisories, or any topic
  where recent/accurate information matters. If you're about to use WebSearch, use this instead. When in doubt,
  search with ais — it's faster, grounded, and returns sources.
---

# ais

Use the `ais` CLI tool to perform web searches. It uses Gemini with two complementary tools:
- **Google Search grounding** — discovers relevant pages and retrieves live results
- **URL context** — fetches and deeply analyzes content from specific URLs you include in the query

These work together in the same query. You can embed a URL to ground on a specific source *and* ask
Gemini to search the web for surrounding context — in a single call.

## When to use this skill

Use `ais` proactively any time you would otherwise:
- Use the built-in `WebSearch` or `WebFetch` tool to look something up
- Say "as of my knowledge cutoff..." or "I'm not sure about the latest..."
- Need to verify a package version, API signature, or dependency compatibility
- Research an error message, stack trace, or unfamiliar library
- Check current pricing, quotas, or service limits
- Look up recent news, changelogs, or release notes
- Compare tools, frameworks, or services

**Default stance: search first, answer second.** If there's any chance the information has changed since
your training data, run `ais` before answering.

## How to invoke

```bash
ais -q "your search query here"
```

### Full parameter reference

| Flag | Default | Values | Purpose |
|------|---------|--------|---------|
| `-q` | — | string | One-shot query (required) |
| `-temperature` | 0.7 | 0.0–2.0 | Lower = more deterministic/consistent |
| `-thinking-budget` | auto | `none` `low` `medium` `high` `auto` | Reasoning depth before answering |

## ais is a research analyst, not a search engine

Keyword search engines reward narrow, focused queries — one query per result category. `ais` works
differently. Gemini doesn't just retrieve matching pages; it reads them, reasons across them, and
synthesizes a coherent response. A single broad query can surface attribution, timeline, technical
indicators, context, and recommendations simultaneously — work that would require many separate
keyword searches.

**Mental model shift:** Don't ask *"what narrow query gets me the right category of result?"*
Ask *"what would I ask a research analyst who can read everything and brief me on it?"*

This has two practical consequences:

**Prefer broad, unified queries over decomposed ones.** Instead of separate calls for
"attribution," "timeline," and "technical indicators" — write one query that asks for the full
picture. Gemini synthesizes the facets internally. Splitting by facet wastes calls and fragments
context that Gemini would naturally connect.

**Broad→targeted is sequential, not parallel.** The broad call lets Gemini synthesize the full
picture first. You only make a targeted follow-up if something specific and genuinely distinct
surfaces that warrants a dedicated deep-dive. Running all calls in parallel collapses this into
keyword-search behavior and loses the synthesis benefit entirely.

## Parameter tuning by query type

Match the flags to what you're looking for — this meaningfully affects speed and accuracy:

**Factual lookups** — pricing, versions, rate limits, release dates:
```bash
ais -q "..." -temperature 0.2 -thinking-budget none
```
Low temperature prevents paraphrasing or rounding numbers. `none` skips reasoning entirely since
the answer is just a fact to retrieve — faster and equally accurate.

**Routine searches** — error messages, how-to questions, API docs:
```bash
ais -q "..."
```
Defaults are fine. No need to tune.

**Comparisons and planning** — "X vs Y for use case Z", architecture decisions, tradeoff analysis:
```bash
ais -q "..." -thinking-budget medium
```
`medium` (8k thinking tokens) lets Gemini reason across the retrieved sources before synthesising
a recommendation, which improves quality for multi-factor decisions.

**Broad research** — trends, overviews, "what is working", landscape surveys:
```bash
ais -q "..." -thinking-budget medium
```
Use `medium` (not `low`) — Gemini needs reasoning capacity to synthesize across many sources.
Do not split broad research into multiple narrower calls; one well-scoped query with `medium` or
`high` produces a more coherent synthesis than three fragmented ones.

**Complex research** — security advisories, multi-step migration guides, deep technical analysis:
```bash
ais -q "..." -thinking-budget high
```
Reserve `high` (24k tokens) for genuinely complex questions where reasoning over sources matters.

**Note on `low`** — `low` (1024 thinking tokens) is rarely the right choice. It provides minimal
reasoning benefit over `none` but costs more. Skip it: use `none` for pure fact retrieval, `medium`
for anything requiring synthesis or comparison.

## Combining URL context with web search

The most powerful pattern: embed a URL in the query *and* ask a web-search question in the same prompt.
Gemini will analyze the specific URL *and* ground the broader question on live search results simultaneously.

**When to use this pattern:**
- User shares a document/PDF/page and wants both a summary *and* external context (reactions, trends, comparisons)
- Analyze a specific artifact (release notes, RFC, report) against what the web is saying about it
- Deep-dive a known source while also checking for recent developments

**Examples:**

```bash
# Analyze a PDF and search for analyst reactions
ais -q "Give me an overview of this earnings report: https://abc.xyz/assets/.../2025q2-alphabet-earnings-release.pdf — also search for reactions from major financial analysts and the current market trend"

# Read a changelog and check if issues are resolved upstream
ais -q "What breaking changes are in https://github.com/org/repo/blob/main/CHANGELOG.md for v3? Also search for any open issues or workarounds people have found"

# Analyze a specific doc and compare against alternatives
ais -q "Summarize the auth model at https://docs.service.com/auth — how does it compare to what other services are doing in 2026?"
```

**URL-only (no web search needed):**
```bash
# Just analyze one or more specific pages
ais -q "Compare the rate limit policies at https://docs.service-a.com/limits and https://docs.service-b.com/limits"
```

Up to 20 URLs per query; URLs must be publicly accessible (no auth, no localhost).

## Single call vs. multiple calls

Default to one call. Use multiple calls only when questions are genuinely distinct — not just
different phrasings of the same question.

**Use one call when:**
- The question is broad but unified (trends, overviews, landscape surveys on a single topic)
- A higher thinking budget can synthesize the breadth internally — prefer `medium`/`high` over
  splitting into separate searches

**Use multiple calls when:**
- Questions are categorically different: e.g., a factual lookup (`none`) + a comparison (`medium`)
- You need to ground on a specific URL *and* separately check a different topic
- A first call surfaces specific names/entities that warrant a targeted follow-up (see below)

**The valid broad→targeted pattern:**
A first broad call discovers relevant entities (names, tools, projects, products). A second
targeted call drills into specific ones that emerged. This is legitimate because the calls are
genuinely distinct in scope — not just rephrasing the same question. Bundle all follow-up
entities into a single call rather than one call per entity.

**Anti-pattern to avoid:** running separate calls for each facet of the same topic (attribution,
timeline, technical details, mitigation — all about one incident). Gemini synthesizes these
internally from a single broad query. Multiple facet calls are keyword-search thinking applied
to a tool that doesn't need it. Combine into one query with `medium`/`high` thinking instead.

## Query writing tips

- Be specific: `"Go 1.23 range-over-func syntax"` beats `"Go new features"`
- Include version numbers when relevant: `"React 19 breaking changes"`
- Include the current year for pricing/limits: `"Gemini 2.0 Flash rate limits 2026"`
- For errors, include the exact error message in quotes
- For comparisons: `"Longhorn vs OpenEBS k3s bare metal 2025"`
- When the user provides a URL, include it in the query — combine with a search question for richer answers

## Getting links and source URLs

`ais` does not return structured source metadata — it returns freeform markdown. When the user
needs actual URLs (e.g., "give me links", "show me sources"), embed the request explicitly in the
query text so Gemini includes them in its answer:

```bash
# Ask for URLs inline — append the request to any query
ais -q "[your question] — include URLs for each example" -thinking-budget medium

# Ask for sources on a factual lookup
ais -q "[your question] — include the official documentation URL" -temperature 0.2 -thinking-budget none
```

If the response omits URLs despite the request, note this to the user and suggest they verify
directly via a follow-up search or by visiting the referenced source.

## Output handling

The tool outputs freeform markdown. Render it directly or summarise key points for the user.

**Known limitations of the current interface:**
- No structured source metadata — URLs must be requested inline in the query (see above)
- No confidence signal — if the answer seems thin or uncertain, retry with a more specific query
  or a higher thinking budget rather than accepting a weak result
- No answer quality indicator — if results look stale or off-topic, reformulate with more specific
  terms, the current year, or an authoritative URL embedded in the query

## Environment requirement

Requires `GEMINI_API_KEY` to be set. If the tool errors with "GEMINI_API_KEY is not set", inform
the user to run: `export GEMINI_API_KEY=<your-key>` before retrying.

## Examples

**Pricing lookup:**
```bash
ais -q "Claude Sonnet GPT-4o Gemini Flash API pricing per million tokens 2026" -temperature 0.2 -thinking-budget none
```

**Version check:**
```bash
ais -q "latest stable version of Go programming language 2026" -temperature 0.2 -thinking-budget none
```

**Error research:**
```bash
ais -q "golang 'context deadline exceeded' in http client fix"
```

**Architecture comparison:**
```bash
ais -q "k3s persistent storage Longhorn vs Rook-Ceph 3-node home lab 2025" -thinking-budget medium
```

**URL + web search (combined):**
```bash
ais -q "Summarize https://example.com/report.pdf and search for how analysts are reacting to it"
```
