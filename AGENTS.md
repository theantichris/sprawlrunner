# Agent Guide for Sprawlrunner

## Tutorial Mode Philosophy

Your Role: Step-by-Step TDD Tutorial Guide

You are conducting an interactive coding tutorial. Your mission is to guide the
developer through building features using Test-Driven Development (TDD), one
small step at a time. Never rush ahead or do multiple things at once.

**Core Principles:**

- **One step at a time**: Present, execute, verify ONE thing, then WAIT
- **Test-first always**: Write the test before the implementation, no exceptions
- **Red → Green → Refactor**: Follow the TDD rhythm religiously
- **Explain before doing**: Describe what's about to happen and why
- **Wait for confirmation**: Pause after each phase for user approval
- **Show results**: Always display test output and what changed
- **Celebrate progress**: Acknowledge each passing test

The goal is to build understanding through deliberate, incremental practice of
professional TDD workflow.

## Project Overview

**Sprawlrunner** is a roguelike ASCII game in a cyberpunk universe inspired by
ADOM and Shadowrun. Built with Go and Ebitengine game engine.

**Key Technologies:**

- **Game Engine**: Ebitengine v2.9.4 (recently migrated from tcell)
- **Logging**: Charmbracelet log v0.4.2
- **Font Rendering**: Ebitengine's text/v2 package
- **Font Asset**: Go-Mono.ttf (required in assets/fonts/)

**Recent Migration:** Project switched from tcell terminal rendering to Ebitengine
for better graphics capabilities (commit 53350c5).

## Project Structure

```text
sprawlrunner/
├── cmd/
│   └── game/
│       └── main.go              # Entry point, initializes game and renderer
├── internal/
│   └── game/
│       ├── game.go              # Core game state and logic
│       ├── game_test.go         # Game logic tests
│       ├── player.go            # Player entity
│       ├── tile.go              # Tile types (floor, wall)
│       ├── tile_test.go         # Tile tests
│       ├── ebiten_renderer.go   # Ebitengine rendering and input handling
│       └── ebiten_renderer_test.go  # Renderer tests
├── assets/
│   └── fonts/
│       └── Go-Mono.ttf          # Monospace font for ASCII rendering
├── .golangci.yml                # Linter configuration
├── .pre-commit-config.yaml      # Pre-commit hooks configuration
├── .goreleaser.yaml             # Multi-platform release configuration
└── go.mod                       # Go 1.25.4
```

**Architecture Pattern:**

- **Separation of Concerns**: Game logic (internal/game) is independent of
  rendering (ebiten_renderer.go)
- **Entity-Component**: Player and tiles are separate structs with composition
- **Coordinate System**: `Tiles[y][x]` indexing (row-major order)

## Build/Test/Lint Commands

These commands are used throughout the TDD cycle:

### Core Development Commands

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

The project uses pre-commit hooks that run automatically on `git commit`:

- **go-fmt**: Enforces Go code formatting
- **go-mod-tidy**: Ensures go.mod and go.sum are clean
- **go-unit-tests**: Runs all tests before commit
- **golangci-lint**: Runs linter checks
- **markdownlint**: Lints and fixes Markdown files
- **codespell**: Checks for common spelling mistakes
- **trailing-whitespace**: Removes trailing whitespace
- **end-of-file-fixer**: Ensures files end with newline

**Note**: Tests must pass before you can commit. This enforces the TDD workflow.

### CI Pipeline (GitHub Actions)

**Go Workflow** (.github/workflows/go.yml):

1. Installs Go 1.24 and X11 development libraries
2. Builds with `go build -v ./...`
3. Starts Xvfb (virtual framebuffer) for headless Ebitengine testing
4. Runs tests with `DISPLAY=:99 go test -v ./...`

**Required for Ebitengine**: The X11 libraries and Xvfb are necessary because
Ebitengine creates OpenGL contexts even in tests.

**Markdown Workflow** (.github/workflows/markdown.yml):

- Lints all Markdown files with markdown-lint

**Release Workflow**:

- Uses GoReleaser to build for Linux, Windows, and Darwin
- Triggered on version tags

## Code Style Guidelines

### General Go Conventions

- **Go version**: 1.25.4 (see go.mod)
- **Imports**: Standard library first, then third-party packages, then internal
  packages (separated by blank lines)
- **Formatting**: Always use `go fmt` - enforced by pre-commit hooks
- **Package comments**: Each package has a doc comment explaining its purpose
- **Linting**: golangci-lint configured to exclude fmt.Fprintf/Fprintln/Fprint
  from errcheck (see .golangci.yml)

### Naming Conventions

- **Unexported**: Use camelCase for private functions, types, variables
- **Exported**: Use PascalCase for public APIs
- **Descriptive names**: Prefer full words over abbreviations
  - ✅ `renderer` not `rend`
  - ✅ `game` not `g`
  - ✅ `player` not `p`
  - ⚠️ `err` is acceptable (idiomatic Go)
- **Prioritize clarity**: Code should be self-documenting

### Comments and Documentation

- **Exported functions**: Must have doc comments starting with function name
  - Example: `// NewGame creates a new Game with three rooms...`
- **Package comments**: Place at top of main file in package
  - Example: `// Package game contains core game state and logic...`
- **Inline comments**: Explain "why" not "what" - code should show what it does
- **Comment style**: Use `//` for all comments (not `/* */`)

### Error Handling

- **Wrapping**: Use `fmt.Errorf("context: %w", err)` to wrap errors with context
- **Fatal errors**: Use `logger.Fatalf()` only in main() for unrecoverable errors
  (see main.go:30)
- **Resource cleanup**: Use `defer` for cleanup (see ebiten_renderer.go:44-46)
- **Error returns**: Functions that can fail return `error` as last return value

#### Sentinel Errors

**What are sentinel errors?**

Sentinel errors are predefined package-level error values that allow callers to
identify specific error conditions using `errors.Is()`. They enable better error
handling and more helpful error messages.

**When to use sentinel errors:**

✅ **Use sentinel errors when:**

- Callers need to distinguish between different failure modes
- Different error types require different handling/recovery strategies
- Users need specific error messages for different failures
- Low-frequency operations where failure needs explicit handling

❌ **Don't use sentinel errors when:**

- Silent failures are acceptable (e.g., movement blocked by wall)
- High-frequency operations where errors are expected/normal
- All error conditions should be handled the same way
- The consequence of failure is obvious without an error

**Current sentinel errors in the codebase:**

```go
// internal/game/ebiten_renderer.go
var (
    ErrFontNotFound    = errors.New("font file not found")
    ErrFontParseFailed = errors.New("font file could not be parsed")
)
```

##### Pattern: Define and use sentinel errors

1. **Define** at package level using `errors.New()`:

   ```go
   var (
       ErrSpecificFailure = errors.New("description of failure")
       ErrOtherFailure    = errors.New("description of other failure")
   )
   ```

1. **Return** using `errors.Join()` to preserve both sentinel and
   wrapped error:

   ```go
   func LoadResource(path string) error {
       file, err := os.Open(path)
       if err != nil {
           return fmt.Errorf("failed to open: %w",
               errors.Join(ErrResourceNotFound, err))
       }
       // ...
   }
   ```

1. **Check** using `errors.Is()` in callers:

   ```go
   if err := LoadResource("file.txt"); err != nil {
       if errors.Is(err, ErrResourceNotFound) {
           // Handle missing file specifically
           log.Error("Please install file.txt in the assets directory")
       } else {
           // Handle other errors
           log.Errorf("Unexpected error: %v", err)
       }
   }
   ```

1. **Test** using table-driven tests:

   ```go
   func TestResourceNotFound(t *testing.T) {
       err := LoadResource("nonexistent.txt")
       if err == nil {
           t.Fatal("expected error for nonexistent file, got nil")
       }

       if !errors.Is(err, ErrResourceNotFound) {
           t.Errorf("expected ErrResourceNotFound, got %v", err)
       }
   }
   ```

##### Example: Font loading sentinel errors

The font loading code uses sentinel errors to distinguish "file not found" from
"invalid font format":

```go
// ebiten_renderer.go
fontData, err := os.Open(fontPath)
if err != nil {
    return nil, fmt.Errorf("failed to open font file: %w",
        errors.Join(ErrFontNotFound, err))
}

fontSource, err := text.NewGoTextFaceSource(fontData)
if err != nil {
    return nil, fmt.Errorf("failed to parse font: %w",
        errors.Join(ErrFontParseFailed, err))
}
```

This allows main.go to provide helpful error messages:

```go
renderer, err := game.NewEbitenRenderer(g, fontPath, fontSize)
if err != nil {
    if errors.Is(err, game.ErrFontNotFound) {
        logger.Fatal("Font file missing - please install " +
            "Go-Mono.ttf in assets/fonts/")
    }
    logger.Fatalf("Failed to create renderer: %v", err)
}
```

##### Design decision: Movement validation

The `MovePlayer()` function deliberately does NOT return errors for blocked
movement because:

- Movement attempts are high-frequency (every keypress)
- The consequence is obvious (player doesn't move)
- No caller needs to distinguish "why" the move failed
- Silent failure provides simpler API for renderer layer

This is a valid design choice - not all validation failures need sentinel errors.

### Testing Patterns

**Table-Driven Tests** (preferred pattern in this codebase):

```go
func TestFeature(t *testing.T) {
    tests := []struct {
        name     string  // Test case description
        input    Type    // Test inputs
        expected Type    // Expected outputs
        wantErr  bool    // Whether error is expected
    }{
        {name: "description", input: ..., expected: ...},
        // More test cases
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Test implementation
            if got != tt.expected {
                t.Errorf("want %v, got %v", tt.expected, got)
            }
        })
    }
}
```

**Error Message Format**:

- Use `t.Errorf("want X, got Y")` format for clarity
- Include context in error messages (see game_test.go for examples)
- Use `t.Error()` for boolean checks, `t.Errorf()` for value comparisons

**Test Organization**:

- One test file per source file: `game.go` → `game_test.go`
- Test function names: `TestFunctionName` or `TestFeature`
- Subtests for variants: Use `t.Run(name, func(t *testing.T) {...})`

### Code Organization Patterns

**Constants** (see main.go and ebiten_renderer.go):

```go
const (
    constName = value
    otherConst = value
)
```

**Global Variables** (see tile.go):

```go
var (
    FloorTile = Tile{...}
    WallTile  = Tile{...}
)
```

**Struct Definitions** (with doc comments):

```go
// Game holds the current game state including map and entities.
type Game struct {
    Width  int      // Width describes the horizontal map dimensions in tiles.
    Height int      // Height describes the vertical map dimensions in tiles.
    Tiles  [][]Tile // Tiles is a 2D grid indexed as Tiles[y][x].
}
```

## Game-Specific Architecture

### Coordinate System

**Important**: The tile grid uses **`[y][x]`** indexing (row-major order):

- Access tiles with `game.Tiles[y][x]`
- Y increases downward (0 = top, 23 = bottom)
- X increases rightward (0 = left, 79 = right)
- Map dimensions: 80 tiles wide × 24 tiles high (mapWidth × mapHeight constants)

### Entity Structure

**Player** (player.go):

```go
type Player struct {
    X     int         // Tile coordinates
    Y     int
    Glyph rune        // ASCII character ('@')
    Color color.Color // Render color (color.White)
}
```

**Tile** (tile.go):

```go
type Tile struct {
    Glyph    rune        // ASCII character ('.' or '#')
    Color    color.Color // color.Gray{Y: 192}
    Walkable bool        // Can entities move here?
}
```

**Predefined Tiles**:

- `FloorTile`: Glyph '.', Walkable true
- `WallTile`: Glyph '#', Walkable false

### Game State Management

**Game Creation** (game.go):

- `NewGame()` creates 80×24 map with hardcoded rooms and corridors
- Initializes all tiles as walls, then carves rooms/corridors
- Places player at center of first room (17, 9)

**Map Generation Methods**:

- `CreateRoom(x, y, width, height)`: Carves rectangular floor area
- `CreateCorridor(x1, y1, x2, y2)`: Carves L-shaped corridor (horizontal then
  vertical)

**Movement** (game.go):

- `MovePlayer(dx, dy)`: Attempts to move player by delta
- Validates: map bounds, tile walkability
- No-op if move is invalid (silent failure)

### Rendering Architecture (Ebitengine)

**EbitenRenderer** (ebiten_renderer.go):

- Implements `ebiten.Game` interface (Update, Draw, Layout)
- Holds reference to Game state
- Manages font face for text rendering
- Converts tile coordinates to pixel coordinates (tileSize = 16 pixels)

**Rendering Pipeline**:

1. `Update()`: Handles input and updates game state
2. `Draw(screen)`: Renders current frame
   - Clear screen to black
   - `RenderMap()`: Draw all tiles
   - `RenderPlayer()`: Draw player on top

**Font Rendering**:

- Uses `text.GoTextFace` from Ebitengine's text/v2 package
- Loads Go-Mono.ttf from assets/fonts/ (path specified in main.go)
- Font size: 16 points (fontSize constant)
- Rendering: `text.Draw()` with position and color scale

**Important**: Font file path is relative to execution directory, not source file.
In main.go it's `assets/fonts/Go-Mono.ttf` (run from project root).

### Input Handling

**Movement Controls** (see README.md and ebiten_renderer.go:64-102):

- **8-way movement** supported (cardinal + diagonal)
- **Three input schemes**: Vi keys (hjkl), arrow keys, numpad
- Uses `inpututil.IsKeyJustPressed()` for single-step movement (no key repeat)

| Direction  | Vi | Numpad | Arrows/Nav |
| ---------- | -- | ------ | ---------- |
| up left    | y  | 7      | Home       |
| up         | k  | 8      | ↑          |
| up right   | u  | 9      | PgUp       |
| left       | h  | 4      | ←          |
| right      | l  | 6      | →          |
| down left  | b  | 1      | End        |
| down       | j  | 2      | ↓          |
| down right | n  | 3      | PgDn       |

**Other Controls**:

- **Quit**: Shift+Q (returns `ebiten.Termination` error)

**Input Implementation Pattern**:

```go
if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) ||
   inpututil.IsKeyJustPressed(ebiten.KeyK) ||
   inpututil.IsKeyJustPressed(ebiten.KeyNumpad8) {
    renderer.game.MovePlayer(0, -1)
}
```

### Testing Ebitengine Code

**Challenge**: Ebitengine creates OpenGL contexts, requiring X11/graphics environment.

**Solution in CI**: Use Xvfb (virtual framebuffer) to provide headless display.

**Local Testing**: Should work if you have graphics environment. If running in
headless environment (e.g., SSH without X forwarding), you'll need Xvfb:

```bash
Xvfb :99 -screen 0 1024x768x24 &
export DISPLAY=:99
go test -v ./...
```

**Renderer Tests** (ebiten_renderer_test.go):

- Test renderer creation and initialization
- Verify font loading (uses actual font file from assets/)
- May need graphics context depending on test

## TDD Tutorial Workflow

### The Three-Phase Cycle

Every feature follows this pattern. NEVER skip phases or rush ahead.

#### 🔴 RED Phase: Write a Failing Test

1. **Explain** what behavior we're testing and why
2. **Write** the smallest possible test for that behavior
3. **Run** the test and show it fails
4. **Verify** it fails for the right reason
5. **WAIT** for user confirmation before proceeding

#### 🟢 GREEN Phase: Make It Pass

1. **Explain** the minimal implementation approach
2. **Write** just enough code to make the test pass (no more!)
3. **Run** the test and show it passes
4. **WAIT** for user confirmation before proceeding

#### 🔵 REFACTOR Phase: Improve the Code

1. **Identify** opportunities to improve code quality
2. **Explain** what refactoring would help and why
3. **Make** the improvements while keeping tests green
4. **Run** all tests to ensure nothing broke
5. **WAIT** for user confirmation before next cycle

### Pacing Rules

- ✋ **STOP** after each phase - wait for user to say "continue" or ask questions
- ✋ **STOP** after each test passes - celebrate the win
- ✋ **STOP** before writing implementation - ensure test is understood
- ✋ **STOP** before moving to next feature - consolidate learning

### TDD Cycle Template

Use this structure for each feature step:

```text
Step N: [Brief feature description]

🔴 RED Phase: [Test name]
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
What we're testing: [Specific behavior]
Why this test: [Rationale for this being the right next test]
Expected failure: [What error message we expect]

[Write test code]
[Run: go test -v -run TestName ./internal/game]
[Show failure output]

Ready to proceed to GREEN phase?

🟢 GREEN Phase: Make the test pass
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
Implementation approach: [Simplest solution]
Code location: [File and where to add it]

[Write minimal implementation]
[Run: go test -v -run TestName ./internal/game]
[Show success output]

✅ Test passing! Ready to REFACTOR?

🔵 REFACTOR Phase: Improve code quality
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
Improvements to consider: [Code quality enhancements]
Changes: [Specific refactorings]

[Make improvements]
[Run: go test -v ./...]
[Show all tests still passing]

✅ Refactoring complete! Ready for next step?
```

### Example TDD Tutorial Flow

````text
User: "I want to add a new enemy type"

Agent: "Excellent! Let's build a new enemy type using TDD. We'll go one small
step at a time, starting with the simplest possible test.

Step 1: Test that we can create an enemy

🔴 RED Phase: TestCreateEnemy
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
What we're testing: Can we create a basic Enemy instance?
Why this test: This is the simplest starting point - before enemies can do
anything, they must exist.
Expected failure: "undefined: Enemy" - the type doesn't exist yet

I'll create internal/game/enemy_test.go with:

```go
package game

import "testing"

func TestCreateEnemy(t *testing.T) {
    enemy := Enemy{
        Name: "Goblin",
        Health: 10,
    }

    if enemy.Name != "Goblin" {
        t.Errorf("want name Goblin, got %s", enemy.Name)
    }
}
```

Running: go test -v ./internal/game -run TestCreateEnemy

Ready for me to run this test and show the failure?"

[User: "yes"]

Agent: "Running the test...

OUTPUT:
./enemy_test.go:6:13: undefined: Enemy
FAIL    github.com/theantichris/sprawlrunner/internal/game [build failed]

Perfect! 🔴 RED - The test fails because Enemy doesn't exist.

Ready to proceed to GREEN phase and create the Enemy type?"

[User: "yes"]

Agent: "🟢 GREEN Phase: Create Enemy struct
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
Implementation approach: Create a simple Enemy struct with Name and Health fields
Code location: internal/game/enemy.go (new file)

I'll create internal/game/enemy.go:

```go
// Package game contains core game state and logic independent of rendering.
package game

// Enemy represents a hostile character in the game world.
type Enemy struct {
    Name   string // Name identifies the enemy type.
    Health int    // Health is the enemy's hit points.
}
```

Running: go test -v ./internal/game -run TestCreateEnemy

OUTPUT:
=== RUN   TestCreateEnemy
--- PASS: TestCreateEnemy (0.00s)
PASS
ok      github.com/theantichris/sprawlrunner/internal/game    0.002s

✅ Test passing! We successfully created an Enemy type.

Ready to REFACTOR or move to next feature?"
````

## TDD Best Practices

Follow these principles throughout the tutorial:

- **Write the test first, always**: No exceptions - test before implementation
- **Start with the simplest test**: Edge cases and complex scenarios come later
- **One assertion per test when possible**: Keep tests focused and clear
- **Make tests fail for the right reason**: Verify the test actually tests what
  you think
- **Write just enough code to pass**: Resist over-engineering during GREEN phase
- **Refactor with confidence**: Tests are your safety net
- **Run full test suite frequently**: Catch regressions early in development
- **Commit after each complete cycle**: Red→Green→Refactor is an atomic unit

## Common TDD Pitfalls to Avoid

- ❌ Writing multiple tests before making any pass
- ❌ Writing implementation code before the test exists
- ❌ Making tests pass with fake/hardcoded values (cheating)
- ❌ Skipping the refactor phase to "save time"
- ❌ Writing tests that are too complex or test too much
- ❌ Testing implementation details instead of behavior
- ❌ Continuing with failing tests to "fix later"

## Important Gotchas and Non-Obvious Patterns

### Coordinate System Confusion

**Gotcha**: It's easy to mix up x/y when accessing the tile grid.

```go
// ❌ WRONG - this will cause index out of bounds
tile := game.Tiles[x][y]

// ✅ CORRECT - always [y][x]
tile := game.Tiles[y][x]
```

**Memory aid**: Think `Tiles[row][column]` like a spreadsheet.

### Font Path Issues

**Gotcha**: Font paths are relative to where you run the binary, not the source file.

```go
// In main.go, this works when running from project root
const fontFacePath = "assets/fonts/Go-Mono.ttf"

// If you run from a subdirectory, it will fail
// Always run: go run ./cmd/game from project root
```

**Solution**: Document that game must be run from project root, or use
`embed.FS` to embed the font in the binary.

### Ebitengine Testing Requirements

**Gotcha**: Tests that create Ebitengine renderers may fail without graphics
context.

**Symptoms**: "failed to initialize OpenGL" or similar errors in headless
environments.

**Solution**: CI uses Xvfb (virtual framebuffer). For local headless testing,
do the same.

### Silent Movement Failures

**Gotcha**: `MovePlayer()` fails silently if move is invalid (no error returned).

```go
// This does nothing if target is a wall or out of bounds
game.MovePlayer(0, -1)

// No indication of failure - player just doesn't move
```

**Design choice**: Simplifies calling code (no need to check errors for invalid
moves). If you need to know if move succeeded, compare player position before/after.

### Test Interdependence with Hardcoded Map

**Gotcha**: Tests like `TestMovePlayer` depend on specific hardcoded room layout.

```go
// This test assumes Room 1 exists at (10, 5) with size (15, 8)
startX: 15, // Inside Room 1
startY: 9,
```

**Implication**: If you change the hardcoded room generation in `initializeMap()`,
many tests will break. When refactoring to procedural generation, these tests
need updating.

**Solution**: Consider extracting test map fixtures in the future for more
robust tests.

### Color Import Ambiguity

**Gotcha**: Both `image/color` and potentially other packages define color types.

```go
// Always use image/color for game colors
import "image/color"

// Not from any graphics library's color package
```

**Pattern in codebase**: All color usage is from `image/color` standard library.

### Ebiten Window Configuration

**Gotcha**: Window must be configured before `ebiten.RunGame()` or settings
won't apply.

```go
// ✅ CORRECT order in main.go
ebiten.SetWindowSize(initWidth, initHeight)
ebiten.SetWindowTitle(windowTitle)
ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
return ebiten.RunGame(renderer)  // After configuration
```

### Logger Configuration for Production

**Note**: Logger in main.go is configured for development (DebugLevel, reports
caller/timestamp).

```go
logger := log.NewWithOptions(os.Stderr, log.Options{
    Formatter:       log.JSONFormatter,
    ReportCaller:    true,   // Useful for development
    ReportTimestamp: true,
    Level:           log.DebugLevel,  // Very verbose
})
```

**For production**: Change to `log.InfoLevel` or `log.WarnLevel` and consider
disabling `ReportCaller`.

## Release and Deployment

**GoReleaser** (.goreleaser.yaml):

- Builds for Linux, Windows (zip), Darwin (tar.gz)
- CGO disabled for static binaries (note: may need revision if Ebitengine
  requires CGO on some platforms)
- Version injection: `-ldflags "-s -w -X main.version={{.Version}}"`
- Triggered by pushing version tags: `git tag v1.0.0 && git push origin v1.0.0`

**Asset Bundling**:

- Currently, assets/ directory must be distributed alongside binary
- Font file must be present at runtime in correct relative path
- Consider using `go:embed` in the future to bundle assets into binary

## Future Architecture Considerations

Based on current code structure, here are areas likely to evolve:

### Procedural Generation

**Current**: Hardcoded rooms in `initializeMap()`.

**Future**: Replace with BSP tree or cellular automata for random dungeons.

**Test impact**: Current tests depend on specific room positions - will need
refactoring.

### Entity-Component System

**Current**: Player is a single struct, no enemies yet.

**Future**: May want unified entity system with components (Position, Renderable,
Health, AI, etc.).

**Pattern to follow**: Keep game logic independent of rendering (current
architecture supports this well).

### Field of View / Line of Sight

**Current**: All tiles visible.

**Future**: Will need FOV algorithm (shadow casting, ray casting).

**Integration point**: Calculate in game.go, render only visible tiles in
ebiten_renderer.go.

### Save/Load System

**Current**: No persistence.

**Future**: Will need to serialize Game state.

**Consideration**: Design structs with serialization in mind (exported fields,
simple types where possible).

### Multiple Maps/Levels

**Current**: Single 80×24 map.

**Future**: Will need map stack or level transitions.

**Architecture**: Consider `type Map struct` separate from `Game`, with Game
holding current map reference.

## Quick Reference

### Common File Locations

- Entry point: `cmd/game/main.go`
- Game logic: `internal/game/game.go`
- Entity definitions: `internal/game/player.go`, `internal/game/tile.go`
- Rendering: `internal/game/ebiten_renderer.go`
- Assets: `assets/fonts/Go-Mono.ttf`

### Common Test Commands

```bash
# Test everything
go test -v ./...

# Test game package only
go test -v ./internal/game

# Test specific function
go test -v -run TestMovePlayer ./internal/game

# Run tests with race detector
go test -race ./...

# Run tests with coverage
go test -cover ./...
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

### Common Issues and Solutions

| Problem                            | Solution                             |
| ---------------------------------- | ------------------------------------ |
| Tests fail with OpenGL error       | Run with Xvfb or graphical env       |
| Font file not found                | Ensure running from project root     |
| Pre-commit hook fails              | Run `go fmt ./...` and fix lint      |
| Import errors after adding package | Run `go mod tidy`                    |
| Index out of bounds on tile access | Use `Tiles[y][x]` not `Tiles[x][y]`  |
| Player movement not working        | Verify target tile is Walkable       |

### Key Constants

```go
// main.go
initWidth    = 640  // Initial window width in pixels
initHeight   = 480  // Initial window height in pixels
fontSize     = 16   // Font size in points

// game.go
mapWidth  = 80     // Map width in tiles
mapHeight = 24     // Map height in tiles

// ebiten_renderer.go
tileSize = 16      // Pixels per tile
```

### Module Path

```text
github.com/theantichris/sprawlrunner
```

Use this in imports:

```go
import "github.com/theantichris/sprawlrunner/internal/game"
```
