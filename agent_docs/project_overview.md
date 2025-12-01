# Project Overview

**Sprawlrunner** is a roguelike ASCII game in a cyberpunk universe inspired by
ADOM and Shadowrun. Built with Go and Ebitengine game engine.

**Key Technologies:**

- **Game Engine**: Ebitengine v2.9.4 (recently migrated from tcell)
- **Logging**: Charmbracelet log v0.4.2
- **Font Rendering**: Ebitengine's text/v2 package
- **Font Asset**: Go-Mono.ttf (required in assets/fonts/)

## Project Structure

```text
sprawlrunner/
├── cmd/
│   └── game/
│       └── main.go              # Entry point, initializes game and renderer
├── internal/
│   └── game/
│       ├── game.go              # Core game state and logic
│       ├── game_test.go         # Tests for core game state and logic
│       ├── player.go            # Player entity and behavior
│       ├── player_test.go       # Tests for player-specific behavior
│       ├── tile.go              # Tile definitions, map representation
│       ├── tile_test.go         # Tests for tiles and map behavior
│       ├── ebiten_renderer.go   # Ebitengine rendering and input handling
│       ├── ebiten_renderer_test.go # Tests for renderer-specific integration
│       └── errors.go            # Sentinel error definitions
├── assets/
│   └── fonts/
│       └── Go-Mono.ttf          # Required font asset (monospaced)
├── .github/
│   └── workflows/
│       └── ci.yml               # CI configuration (tests, lint, build)
├── .golangci.yml                # golangci-lint configuration
├── .pre-commit-config.yaml      # Pre-commit hooks
├── go.mod                       # Go module definition
├── go.sum                       # Module checksums
└── AGENTS.md                    # Agent guide
```
