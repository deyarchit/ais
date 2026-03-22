# Technology Stack

**Analysis Date:** 2026-03-22

## Languages

**Primary:**
- Go 1.25.7 - Backend application language

## Runtime

**Environment:**
- Go 1.25.7 runtime

**Package Manager:**
- Go modules (go.mod/go.sum)
- Lockfile: Not detected (go.sum not present in repository)

## Frameworks

**Core:**
- Not yet initialized - project is in early setup stage

**Testing:**
- Go standard testing library (implied by test target in `Makefile`)

**Build/Dev:**
- golangci-lint - Linting and code quality
- gofmt/goimports - Code formatting

## Key Dependencies

**Code Quality:**
- golangci-lint - Integrated linter configured in `.golangci.yml`

## Configuration

**Environment:**
- Go module system via `go.mod` at `/Users/deyarchit/Projects/ai/ais/go.mod`
- Environment variables: `.env` and `.env.local` files are gitignored but not yet created
- Configuration managed through build targets in `Makefile` at `/Users/deyarchit/Projects/ai/ais/Makefile`

**Build:**
- `.golangci.yml` - Comprehensive linting configuration at `/Users/deyarchit/Projects/ai/ais/.golangci.yml`
- `Makefile` - Build and development targets at `/Users/deyarchit/Projects/ai/ais/Makefile`
- `go.mod` - Module definition at `/Users/deyarchit/Projects/ai/ais/go.mod`

## Linting & Code Quality

**Enabled Linters** (via golangci-lint):

**Correctness:**
- errcheck - Error checking enforcement
- govet - Standard Go vet violations
- staticcheck - Idiomatic and bug-free Go code
- ineffassign - Dead assignment detection
- unused - Unused code detection
- nilerr - Error hiding prevention
- errorlint - Error wrapping support
- bodyclose - Resource leak prevention
- durationcheck - Time math correctness
- forbidigo - Pattern forbidding (forbids `log.*` and `fmt.Print.*` in favor of `log/slog`)

**Security & Safety:**
- copyloopvar - Loop variable safety (Go 1.22+ feature)

**Clean Code & Performance:**
- revive - General best practices with custom rules
- gocritic - Code diagnostics and refactoring
- unparam - Clean function signatures

**Style & Simplification:**
- unconvert - Redundant type conversion removal
- misspell - English spelling verification

**Formatters:**
- gofmt - Canonical Go formatting
- goimports - Clean, grouped imports

## Platform Requirements

**Development:**
- Go 1.25.7 installed
- golangci-lint installed (for linting)
- Make utility (for Makefile execution)

**Production:**
- Go 1.25.7 runtime or compiled binary
- No specific platform requirements detected yet

---

*Stack analysis: 2026-03-22*
