---
phase: quick
plan: 260323-u7z
subsystem: build
tags: [makefile, install, developer-experience]
dependency_graph:
  requires: []
  provides: [make-install-target]
  affects: [Makefile]
tech_stack:
  added: []
  patterns: [make-dependency-chain]
key_files:
  created: []
  modified:
    - Makefile
decisions:
  - "install target depends on build so binary is always fresh before copy"
  - "no sudo in Makefile; user runs sudo make install as needed (standard Unix convention)"
metrics:
  duration: "2min"
  completed_date: "2026-03-23"
  tasks_completed: 1
  files_modified: 1
---

# Quick Task 260323-u7z: Add make install target

**One-liner:** `make install` target added to Makefile that builds then copies `./bin/ais` to `/usr/local/bin/ais`.

## What Was Done

Added a single `install` target to the Makefile immediately after the `build` target. The target:

1. Declares `build` as a prerequisite — the binary is always rebuilt before copying
2. Runs `cp ./bin/ais /usr/local/bin/ais` to place the binary on the system PATH
3. Uses a `## install:` comment consistent with the existing comment style

## Usage

```bash
sudo make install   # build + copy to /usr/local/bin (sudo required on macOS)
ais --help          # works from any directory
```

Note: `sudo` is required on macOS because `/usr/local/bin` is owned by root. This is standard Unix convention — the Makefile itself does not include `sudo` so it remains portable.

## Verification

Binary builds successfully (`go build` exits 0). The `cp` step fails without sudo due to macOS permissions, which is expected and correct behavior.

```
/Users/deyarchit/Projects/ai/ais/bin/ais --help  # works, confirms binary is valid
```

## Deviations from Plan

None - plan executed exactly as written.

## Known Stubs

None.

## Self-Check: PASSED

- Makefile contains `install:` target — FOUND
- Makefile `install` depends on `build` — FOUND
- Commit 02576e1 exists — FOUND
