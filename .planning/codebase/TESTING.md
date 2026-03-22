# Testing

## Framework

- **Standard Go testing** (`testing` package) — no third-party test framework detected
- Tests run with: `go test ./... -coverprofile=coverage.txt -coverpkg=./internal/...`
- Coverage report: `go tool cover -html=coverage.txt -o coverage.html`

## Commands

| Command | Description |
|---------|-------------|
| `make test` | Run all tests + generate coverage report |
| `make pr` | Full pre-PR pipeline: tidy → lint → fmt → test |

## Coverage

- Coverage profile written to `coverage.txt`
- HTML report written to `coverage.html`
- Coverage measured over `./internal/...` packages

## Linter Exclusions for Tests

- `errcheck` excluded in `_test.go` files — error returns in tests don't need explicit checks
- `bodyclose` excluded in `_test.go` files — HTTP response body close not required in tests

## Notes

- No test files exist yet (project in initialization phase)
- Test structure will follow standard Go conventions: `*_test.go` files alongside source
- Internal packages (`./internal/`) are the primary coverage target

---
*Mapped: 2026-03-22*
