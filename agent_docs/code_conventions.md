# Code Style Guidelines

## General Go Conventions

- **Go version**: 1.25.4 (see go.mod)
- **Imports**: Standard library first, then third-party packages, then internal
  packages (separated by blank lines)
- **Formatting**: Always use `go fmt` - enforced by pre-commit hooks
- **Package comments**: Each package has a doc comment explaining its purpose
- **Linting**: golangci-lint configured to exclude fmt.Fprintf/Fprintln/Fprint
  from errcheck (see .golangci.yml)

## Naming Conventions

- **Unexported**: Use camelCase for private functions, types, variables
- **Exported**: Use PascalCase for public APIs
- **Descriptive names**: Prefer full words over abbreviations
  - ✅ `renderer` not `rend`
  - ✅ `game` not `g`
  - ✅ `player` not `p`
  - ⚠️ `err` is acceptable (idiomatic Go)
- **Prioritize clarity over brevity**; this is a learning project.

## Comments and Documentation

- **Package-level**: Each package should start with a doc comment explaining
  its role in the game.
- **Exported functions and types**: Must have comments that explain behavior,
  not implementation details.
- **Non-obvious logic**: Add inline comments to clarify intent (especially if
  tied to game design decisions).

### Error Handling

- Always return `error` as the last return value.
- Use `fmt.Errorf("...: %w", err)` for wrapping errors.
- Use sentinel errors where behavior needs to branch based on error kind.
- Prefer `errors.Is` and `errors.As` for error handling logic.

#### Sentinel Errors

Sentinel errors are used to distinguish error cases that should be handled
differently (e.g., missing assets vs corrupted assets).

##### Pattern: Define and use sentinel errors

1. Define sentinel error:

   ```go
   var (
       ErrFontNotFound = errors.New("font not found")
   )
   ```

2. Wrap sentinel when returning:

   ```go
   if errors.Is(err, fs.ErrNotExist) {
       return fmt.Errorf("%w: %v", ErrFontNotFound, err)
   }
   ```

3. Use `errors.Is` in callers:

   ```go
   if errors.Is(err, ErrFontNotFound) {
       // handle missing font
   }
   ```

**Example from this project**: The `errors.go` file defines `ErrFontNotFound`
 for missing font assets. Renderer code wraps this sentinel when font loading
 fails, and the caller can distinguish between missing fonts vs. other errors
 using `errors.Is(err, ErrFontNotFound)`.

## Test Organization

Keep tests in *_test.go files next to the code they test.

Group related tests in the same file for readability.

Use table-driven tests for combinations of inputs and expected outputs.
