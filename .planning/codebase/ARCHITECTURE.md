# Architecture

## Pattern

Layered server application following standard Go project layout. The application is in early initialization phase with the entry point and internal package structure planned but not yet populated.

## Layers

| Layer | Location | Responsibility |
|-------|----------|----------------|
| Command | `./cmd/server/` | Application entry point, config wiring, server startup |
| Internal | `./internal/` | Domain logic, services, handlers (planned) |
| Cross-cutting | `./internal/` | Logging (slog), validation, auth |

## Entry Points

- **Server binary:** `./cmd/server` — built to `./bin/server`
- **Dev run:** `go run ./cmd/server`

## Data Flow

```
HTTP Request
  → cmd/server (entry/routing)
  → internal/handlers (request handling)
  → internal/services (business logic)
  → internal/store (data access)
  → Response
```

## Key Design Constraints (from linter config)

- Structured logging only via `slog` — no `fmt.Print*` for logs
- Errors must be wrapped with context (not swallowed)
- Structured concurrency — goroutine usage controlled
- No copy-paste code — enforced by duplication linter
- Strict nil checks enforced

## Build System

- **Makefile** targets: `build`, `dev`, `test`, `lint`, `fmt`, `tidy`, `pr`
- `make pr` = full pre-PR pipeline (tidy → lint → fmt → test)
- **Test coverage** output: `coverage.txt` + `coverage.html`
- **Lint:** golangci-lint with 20+ enabled linters

---
*Mapped: 2026-03-22*
