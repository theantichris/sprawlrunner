# Agent Guide for Sprawlrunner

## Build/Test/Lint Commands
- **Build**: `go build -v ./...` or `go build -v ./cmd/game`
- **Run**: `go run ./cmd/game`
- **Test all**: `go test -v ./...`
- **Test single**: `go test -v -run TestName ./path/to/package`
- **Lint**: `golangci-lint run`
- **Format**: `go fmt ./...`
- **Tidy deps**: `go mod tidy`

## Code Style Guidelines
- **Go version**: 1.25.4 (see go.mod)
- **Imports**: Standard library first, then external packages (separated by blank line)
- **Formatting**: Use `go fmt` - enforced by pre-commit hooks
- **Error handling**: Wrap errors with context using fmt.Errorf or log.Fatalf for fatal errors (see main.go:12,16)
- **Naming**: Use camelCase for unexported, PascalCase for exported; descriptive function names
- **Comments**: Document exported functions with comments starting with function name (see main.go:56)
- **Linting**: golangci-lint configured to exclude fmt.Fprintf/Fprintln/Fprint from errcheck

## Pre-commit Hooks
Project uses pre-commit hooks for: go-fmt, go-mod-tidy, go-unit-tests, golangci-lint, markdownlint, codespell, trailing-whitespace, end-of-file-fixer
