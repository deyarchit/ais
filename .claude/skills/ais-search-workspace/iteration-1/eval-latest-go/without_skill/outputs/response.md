# Latest Stable Version of Go

## Version: Go 1.23

As of my training knowledge cutoff (August 2025), the latest stable release of Go is **Go 1.23**, released in August 2024. Go 1.22 was released in February 2024, and Go 1.23 followed in August 2024.

Note: Go 1.24 may have been released in early 2025 (Go follows a roughly 6-month release cadence, typically releasing in February and August), but I do not have confident detailed knowledge of its headline features. The most well-documented recent major release in my training data is Go 1.23.

---

## Go 1.23 Headline Features and Changes

### Language Changes

- **Range over integers**: Building on Go 1.22's introduction of range-over integers, Go 1.23 continued refining iterator support.
- **Range over function iterators (experimental in 1.22, progressing in 1.23)**: The ability to range over functions that yield values (iterator functions), following the `iter` package patterns, moved forward significantly.

### Standard Library

- **New `iter` package**: Introduced the `iter` package providing support for user-defined iterators that work with the `for range` loop. This is a foundational addition for the Go ecosystem's approach to iteration.
- **`slices` and `maps` package updates**: Additional functions were added to the `slices` and `maps` packages to work with the new iterator patterns (e.g., `slices.All`, `slices.Values`, `maps.All`, `maps.Keys`, `maps.Values`).
- **`structs` package**: A new `structs` package was introduced.
- **`unique` package**: A new `unique` package was added for interning/canonicalizing comparable values.

### Timer Changes

- **`time.Timer` and `time.Ticker` garbage collection fix**: Timers and tickers that are no longer referenced are now eligible for garbage collection even if their channel has not been read. This fixes a long-standing subtle resource leak.

### Toolchain

- **`go` tool improvements**: Various improvements to the Go toolchain, including better support for the `toolchain` directive in `go.mod`.

### Platform Support

- Continued improvements to WebAssembly (WASM) support and other platform targets.

---

## Go 1.22 Headline Features (for context)

Go 1.22 (February 2024) introduced several notable changes:

- **Loop variable semantics fix**: Each iteration of a `for` loop now creates a new variable (fixing the long-standing loop variable capture bug that had caused countless subtle errors).
- **Range over integers**: `for i := range N` syntax for ranging over integers.
- **`math/rand/v2`**: A new, improved random number generation package.
- **HTTP routing enhancements**: The `net/http` package's `ServeMux` gained support for method-based routing (e.g., `GET /path`) and wildcards in route patterns.

---

## Summary

The most significant recent additions to Go center on:
1. The iterator/range-over-function feature (a major ergonomic improvement)
2. The loop variable fix (Go 1.22) eliminating a classic Go gotcha
3. New standard library packages (`iter`, `unique`, `structs`)
4. HTTP routing improvements in the standard library

These changes reflect Go's continued evolution while maintaining its commitment to backward compatibility and simplicity.

---

*Note: This answer is based on training knowledge with a cutoff of August 2025. The actual latest stable version at the time of reading may be Go 1.24 or later. For the most current information, consult https://go.dev/doc/devel/release.*
