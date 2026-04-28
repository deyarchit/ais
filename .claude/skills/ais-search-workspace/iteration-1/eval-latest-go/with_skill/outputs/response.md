# Latest Stable Go Version

As of March 23, 2026, the latest stable version of the Go programming language is **Go 1.26.1**. Go 1.26.0 was released on February 10, 2026.

## Go 1.25 Headline Features (August 2025)

While Go 1.25 was not a flashy release with major syntax changes, it focused on practical enhancements, runtime safety, and smarter tooling:

- **Generic Optimization** — Removed the "Core Types" concept introduced in Go 1.18, simplifying the language and enhancing generics flexibility by defining rules directly on concrete types and type sets.
- **Safer Nil-Pointer Handling** — Fixed a bug from Go 1.21 that sometimes prevented nil pointer panics, ensuring reliable panics on nil dereferences.
- **Experimental `encoding/json/v2` package** — Almost entirely rewritten: 3-10x faster deserialization, zero heap allocation, streaming support for large documents, and more flexible custom serialization.
- **`testing/synctest` package** — Graduated to general availability; runs tests in isolated "bubbles" with virtualized time (was experimental in Go 1.24).
- **Tooling Advancements** — `go vet` catches more common mistakes; `go doc -http` serves documentation locally; `go build -asan` auto-detects memory leaks.
- **DWARF v5** — Compiler and linker generate DWARF version 5 debug info, reducing storage footprint and shortening link time for large binaries.
- **`runtime.AddCleanup`** — More flexible and efficient finalization mechanism replacing older finalizer approach.
- **FIPS 140-3 compliance** — Standard library includes FIPS-approved cryptographic algorithms without code modifications.

*(Sources retrieved via `ais` CLI — see sources.txt)*
