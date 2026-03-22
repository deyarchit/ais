# Requirements: ais

**Defined:** 2026-03-22
**Core Value:** Every query returns a grounded, source-cited answer — either as a quick one-liner or a full conversation — without leaving the terminal.

## v1 Requirements

### Query Modes

- [ ] **MODE-01**: Running `ais -q "query"` returns an answer and exits (one-shot mode)
- [ ] **MODE-02**: Running `ais` with no args opens an interactive REPL (chat mode)
- [ ] **MODE-03**: In chat mode, the full conversation history is passed to Gemini on each turn (multi-turn context)

### Search Grounding

- [ ] **SRCH-01**: Google Search grounding is enabled on every query (no toggle needed)
- [ ] **SRCH-02**: Grounding sources/URLs are listed after each response

### Output

- [ ] **OUT-01**: Responses are rendered as markdown in the terminal (via glamour)
- [ ] **OUT-02**: Source citations appear after the rendered response body

### Configuration

- [ ] **CFG-01**: GEMINI_API_KEY environment variable is used for authentication
- [ ] **CFG-02**: Missing API key produces a clear, actionable error message

### Error Handling

- [ ] **ERR-01**: Missing API key error tells user exactly what to set
- [ ] **ERR-02**: Network/API errors show the failure reason and suggest next steps
- [ ] **ERR-03**: Unknown flags or bad input show usage help

### Tooling

- [ ] **TOOL-01**: `make build` produces `./bin/ais` binary
- [ ] **TOOL-02**: Code passes existing golangci-lint rules

## v2 Requirements

### Input

- **INP-01**: Stdin pipe support — `echo "query" | ais` for scripting workflows

### Output

- **OUT-03**: Streaming output — response streams to terminal as it arrives
- **OUT-04**: Visual chat prompt showing turn number or context length

### Customization

- **CUST-01**: `-s` flag to set a system prompt / persona per session

## Out of Scope

| Feature | Reason |
|---------|--------|
| Config file | Env var is sufficient; config file adds complexity without value for v1 |
| Multiple AI providers | Gemini-only keeps the scope tight |
| Persistent conversation history | Sessions are ephemeral; disk storage is v2+ |
| Web UI / TUI | Terminal text output is the right interface for a search tool |
| OAuth / complex auth | API key is the right auth model for a personal CLI tool |

## Traceability

| Requirement | Phase | Status |
|-------------|-------|--------|
| MODE-01 | Phase 1 | Pending |
| MODE-02 | Phase 1 | Pending |
| MODE-03 | Phase 1 | Pending |
| SRCH-01 | Phase 1 | Pending |
| SRCH-02 | Phase 1 | Pending |
| OUT-01 | Phase 1 | Pending |
| OUT-02 | Phase 1 | Pending |
| CFG-01 | Phase 1 | Pending |
| CFG-02 | Phase 1 | Pending |
| ERR-01 | Phase 1 | Pending |
| ERR-02 | Phase 1 | Pending |
| ERR-03 | Phase 1 | Pending |
| TOOL-01 | Phase 1 | Pending |
| TOOL-02 | Phase 1 | Pending |

**Coverage:**
- v1 requirements: 14 total
- Mapped to phases: 14
- Unmapped: 0 ✓

---
*Requirements defined: 2026-03-22*
*Last updated: 2026-03-22 after initial definition*
