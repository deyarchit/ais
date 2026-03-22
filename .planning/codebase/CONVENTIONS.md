# Code Conventions

## Language & Runtime

- **Go 1.25.7** — latest stable, uses new WaitGroup API (`wg.Go(...)`)
- **Module:** `ais`
- **Formatter:** `gofmt` + `goimports` (run via `make fmt`)

## Logging

- **Only `log/slog`** — `log.*` and `fmt.Print*` are forbidden by linter
- Keys must be `snake_case`
- Messages must be `Capitalized` and static (dynamic data goes in Attrs)
- Forbidden keys: `time`, `level`, `msg`, `source`
- Use key-value pairs only (`slog.String("key", val)`, not positional args)

## Error Handling

- Every error return must be checked (no silent ignores)
- Use `%w` in `fmt.Errorf` for wrapping (never `%v`)
- Use `errors.Is` for comparison, `errors.As` for type assertions
- Never return `nil` in an `err != nil` branch (nilerr)
- Sentinel errors named with `Err` prefix: `var ErrNotFound = errors.New(...)`
- Error strings: lowercase, no trailing punctuation

## Naming

- Standard Go initialisms: `URL`, `ID`, `HTTP`, `JSON` (not `Url`, `Id`, etc.)
- Context parameter: always first, always named `ctx`
- Error always last in multi-return
- Receiver names: short (1-2 letters from type name), consistent across methods
- Exported functions must not return unexported types

## Function Design

- Max 8 parameters (use config struct if exceeded)
- Max 3 return values (use struct if exceeded)
- Max cognitive complexity: 25
- Early returns / guard clauses preferred over nested if/else
- No `else` after `return`/`continue`/`break`

## Imports

- No dot imports (`import . "pkg"`)
- Side-effect imports (`import _ "pkg"`) only in `main` or test files, with comment
- `goimports` enforces grouping: stdlib / external / internal

## Concurrency

- `sync.WaitGroup` must be passed by pointer
- Use `wg.Go(func() {...})` pattern (Go 1.25+ API)
- Use `sync/atomic` for shared variables
- Context keys must be unexported custom types

## Nolint Directives

- Must name the specific linter: `//nolint:errcheck`
- Must include explanation: `//nolint:errcheck // intentionally ignored because...`
- No bare `//nolint` allowed

## os.Exit / log.Fatal

- Only allowed in `main()` and `init()` — everywhere else return an error

---
*Mapped: 2026-03-22*
