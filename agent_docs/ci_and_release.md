# Build/Test/Lint Commands

These commands are used throughout the TDD cycle:

## Core Development Commands

**Test single**:

```bash
go test -v -run TestName ./path/to/package
```

Use during RED/GREEN phases for focused testing.

**Test all**:

```bash
go test -v ./...
```

Use during REFACTOR phase to ensure nothing broke.

**Build binary**:

```bash
go build -v ./cmd/game
```

Build the game executable.

**Build all**:

```bash
go build -v ./...
```

Verify compilation of all packages.

**Run game**:

```bash
go run ./cmd/game
```

Test the game manually after TDD cycle.

**Format**:

```bash
go fmt ./...
```

Run during REFACTOR phase (enforced by pre-commit).

**Lint**:

```bash
golangci-lint run
```

Run during REFACTOR phase for code quality.

**Tidy dependencies**:

```bash
go mod tidy
```

Run when adding new imports.

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

CI is configured in `.github/workflows/go.yml` and runs on every push/PR:

- `go test ./...` - Run all tests
- `golangci-lint run` - Lint checks
- `go build ./cmd/game` - Verify binary builds

Additional workflows:

- `markdown.yml` - Markdown linting
- `release.yml` - Automated releases on tag push

If any check fails, CI will fail and block merging.

## Release and Deployment

Releases are automated via **GoReleaser** and **GitHub Actions**.

### Release Configuration

- **GoReleaser config**: `.goreleaser.yaml` (defines build targets and artifacts)
- **Release workflow**: `.github/workflows/release.yml` (triggered on tag push)

### Semantic Versioning

Follow [Semantic Versioning](https://semver.org/):

- `v0.x.x` - Pre-1.0 development
- `vMAJOR.MINOR.PATCH` - e.g., `v1.2.3`
  - **MAJOR**: Breaking changes
  - **MINOR**: New features (backward compatible)
  - **PATCH**: Bug fixes

### Release Process

1. **Ensure clean state**:

   ```bash
   git status  # Should show clean working tree
   go test ./...  # All tests pass
   golangci-lint run  # No lint issues
   ```

2. **Create and push tag**:

   ```bash
   git tag v0.1.0
   git push origin v0.1.0
   ```

3. **GoReleaser runs automatically** (via GitHub Actions on tag push):
   - Builds binaries for configured platforms
   - Creates GitHub Release with release notes
   - Attaches compiled binaries to release

4. **Verify release**:
   - Check GitHub Releases page for new release
   - Download and test binaries for target platforms

### Manual Release (if needed)

If automated release fails or testing locally:

```bash
goreleaser release --snapshot --clean  # Test without publishing
goreleaser release --clean             # Publish (requires GITHUB_TOKEN)
```

### When Adding New Build Steps

- Update `.goreleaser.yaml` with new build configuration
- Test with `--snapshot` flag first
- Ensure CI covers new build steps
- Document any new dependencies or build requirements here
