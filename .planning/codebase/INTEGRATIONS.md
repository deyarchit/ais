# External Integrations

**Analysis Date:** 2026-03-22

## APIs & External Services

Not yet configured - project is in initialization stage with no external integrations implemented.

## Data Storage

**Databases:**
- Not configured

**File Storage:**
- Not configured

**Caching:**
- Not configured

## Authentication & Identity

**Auth Provider:**
- Not configured

## Monitoring & Observability

**Error Tracking:**
- Not configured

**Logs:**
- Structured logging enforced via log/slog (configured in `.golangci.yml` via sloglint linter)
- Logging approach: Key-value only format required (kv-only: true)
- Log message style: Capitalized sentences with snake_case keys
- Forbidden keys: time, level, msg, source (reserved/redundant in log aggregators)
- Configuration enforced in `.golangci.yml` at `/Users/deyarchit/Projects/ai/ais/.golangci.yml`

## CI/CD & Deployment

**Hosting:**
- Not configured

**CI Pipeline:**
- Not configured

## Environment Configuration

**Required env vars:**
- None configured yet

**Secrets location:**
- `.env` and `.env.local` files are gitignored but not yet created (see `.gitignore` at `/Users/deyarchit/Projects/ai/ais/.gitignore`)

## Webhooks & Callbacks

**Incoming:**
- Not configured

**Outgoing:**
- Not configured

## Build & Development Tools

**Build Artifacts:**
- Binary output: `./bin/server` (from Makefile target `build`)
- Entry point: `./cmd/server` (referenced in Makefile but directory not yet created)

## Development Commands

**Available make targets** (in `/Users/deyarchit/Projects/ai/ais/Makefile`):
- `make build` - Compile Go binary to `./bin/server`
- `make dev` - Run application in development mode via `go run ./cmd/server`
- `make tidy` - Tidy Go modules
- `make lint` - Run golangci-lint verification and checks
- `make fmt` - Auto-fix formatting with golangci-lint
- `make test` - Run all tests with coverage reporting (output to coverage.txt and coverage.html)
- `make pr` - Run tidy, lint, fmt, and test in sequence

---

*Integration audit: 2026-03-22*
