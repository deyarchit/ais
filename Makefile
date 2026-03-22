## build-go: compile the Go binary (requires build-web first)
build:
	go build -o ./bin/ais ./cmd/ais

# ── Development ────────────────────────────────────────────────────────────────
dev:
	go run ./cmd/ais

## Tidy modules
tidy:
	go mod tidy

## Lint with golangci-lint
lint:
	golangci-lint config verify
	golangci-lint run ./...

## Format with golangci-lint (auto-fix formatting)
fmt:
	golangci-lint fmt ./...

update-codemaps:
	@echo "Launching Claude Code to run skill: update-codemaps..."
	claude -p "run the update-codemaps skill" --model "haiku" --allowedTools "Bash,Read,Edit,Write" --output-format stream-json --verbose --include-partial-messages | \
  jq -rj 'select(.type == "stream_event" and .event.delta.type? == "text_delta") | .event.delta.text'
	@echo "Updated codemaps"

## test: run all Go tests
test:
	go test ./... -coverprofile=coverage.txt -coverpkg=./internal/...
	go tool cover -html=coverage.txt -o coverage.html

## Prepare for pull request
pr: tidy lint fmt test

.PHONY: *
