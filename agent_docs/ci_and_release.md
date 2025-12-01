# Build/Test/Lint Commands

These commands are used throughout the TDD cycle:

## Core Development Commands

- **Test single**: `go test -v -run TestName ./path/to/package` - Use during
  RED/GREEN phases for focused testing
- **Test all**: `go test -v ./...` - Use during REFACTOR phase to ensure nothing
  broke
- **Build binary**: `go build -v ./cmd/game` - Build the game executable
- **Build all**: `go build -v ./...` - Verify compilation of all packages
- **Run game**: `go run ./cmd/game` - Test the game manually after TDD cycle
- **Format**: `go fmt ./...` - Run during REFACTOR phase (enforced by pre-commit)
- **Lint**: `golangci-lint run` - Run during REFACTOR phase for code quality
- **Tidy deps**: `go mod tidy` - Run when adding new imports

### Pre-commit Hooks

The project uses pre-commit to enforce formatting and linting:

- Config file: `.pre-commit-config.yaml`
- Typical hooks:
  - `gofmt`
  - `golangci-lint`
  - Others as configured

To install and run:

```bash
pre-commit install
pre-commit run --all-files
```

## CI Pipeline (GitHub Actions)

CI is configured in `.github/workflows/` and typically runs:

1. `go test ./...`
1. `golangci-lint run`
1. `go build ./cmd/game`

If tests, lint, or build fail, CI will fail.

## Release and Deployment

Releases are handled via GoReleaser and GitHub Releases.

Typical flow:

1. Merge PR into main with passing CI.
1. Tag a new version (e.g., v0.1.0).
1. Let GoReleaser or GitHub Actions build and publish binaries.

When adding new build steps:

- Ensure CI covers them.
- Keep commands simple and documented here.
