# ais

A terminal AI assistant powered by [Gemini 2.5 Flash](https://deepmind.google/technologies/gemini/) with always-on Google Search grounding. Ask questions in one-shot mode or hold a multi-turn conversation in an interactive REPL.

## Features

- **Google Search grounding** — every response is grounded in live web results
- **One-shot mode** — pipe queries from scripts or the command line
- **Interactive REPL** — multi-turn chat with persistent history per session
- **Markdown rendering** — responses rendered with syntax highlighting in the terminal
- **Source citations** — optionally print the URLs behind each response
- **Thinking budget** — control how much reasoning the model applies
- **Exponential backoff** — automatic retry on transient 503/429 errors

## Prerequisites

- Go 1.25+
- A [Gemini API key](https://aistudio.google.com/app/apikey)

## Installation

```bash
# Build and install to $GOPATH/bin
go install ./cmd/ais

# Or build locally
make build   # output: ./bin/ais
```

## Setup

```bash
export GEMINI_API_KEY=your_api_key_here
```

Add this to your shell profile (`~/.zshrc`, `~/.bashrc`, etc.) to persist it.

## Usage

### One-shot mode

```bash
ais -q "What is the current Go version?"
```

### Interactive REPL

```bash
ais
```

Starts a `ais>` prompt. History is maintained across turns within the session. Exit with `Ctrl+D`.

### Flags

| Flag | Default | Description |
|---|---|---|
| `-q <query>` | — | Run a single query and exit |
| `--show-refs` | false | Print source URLs after each response |
| `--temperature` | 0.7 | Sampling temperature `[0.0, 2.0]` |
| `--thinking-budget` | auto | Reasoning token budget: `none`, `low`, `medium`, `high`, `auto` |

### Examples

```bash
# One-shot with source citations
ais -q "Latest stable Node.js release?" --show-refs

# Low temperature for deterministic answers
ais -q "What is 2+2?" --temperature 0.1

# Deep reasoning mode
ais -q "Explain the CAP theorem" --thinking-budget high

# Interactive REPL showing sources
ais --show-refs
```

## Development

```bash
make dev      # run without building
make build    # compile to ./bin/ais
make test     # run tests with coverage
make lint     # lint with golangci-lint
make fmt      # auto-fix formatting
make pr       # tidy + lint + fmt + test (pre-PR check)
```

## License

MIT
