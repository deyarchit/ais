# Directory Structure

## Current Layout

```
ais/
├── go.mod              # Module: "ais", Go 1.25.7
├── Makefile            # Build, dev, lint, test, PR targets
├── .planning/          # GSD planning directory
│   └── codebase/       # Codebase map documents
└── bin/                # Compiled binaries (gitignored, created by build)
    └── server          # Output of `make build`
```

## Planned Layout (inferred from Makefile)

```
ais/
├── cmd/
│   └── server/         # Main entry point (go run ./cmd/server)
├── internal/           # Private application packages
│   ├── handlers/       # HTTP handlers (inferred)
│   ├── services/       # Business logic (inferred)
│   └── store/          # Data access (inferred)
├── go.mod
├── go.sum
└── Makefile
```

## Key Locations

| Path | Purpose |
|------|---------|
| `./cmd/server/` | Application entry point (planned) |
| `./internal/` | All private application code (planned) |
| `./bin/server` | Compiled binary output |
| `coverage.txt` | Test coverage profile |
| `coverage.html` | Human-readable coverage report |

## Naming Conventions

- Module name: `ais` (short, no domain prefix)
- Standard Go project layout (`cmd/`, `internal/`)
- Binary named `server`

## Notes

- Project is in initialization phase — `cmd/` and `internal/` directories do not yet exist
- `golangci-lint` config file expected (referenced in Makefile but not yet visible)
- `update-codemaps` Makefile target runs Claude Code haiku to refresh `.planning/codebase/`

---
*Mapped: 2026-03-22*
