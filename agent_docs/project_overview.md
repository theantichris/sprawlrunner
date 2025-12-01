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
├── agent_docs/                  # Agent documentation
│   ├── project_overview.md      # This file - full project structure
│   ├── tdd_workflow.md          # TDD tutorial workflow and examples
│   ├── code_conventions.md      # Go style and error handling patterns
│   └── ci_and_release.md        # Build, test, lint, and release commands
├── .github/
│   └── workflows/
│       ├── go.yml               # CI configuration (tests, lint, build)
│       ├── markdown.yml         # Markdown linting
│       └── release.yml          # Release automation with GoReleaser
├── .gitignore                   # Git ignore patterns
├── .golangci.yml                # golangci-lint configuration
├── .goreleaser.yaml             # GoReleaser configuration for builds
├── .pre-commit-config.yaml      # Pre-commit hooks (gofmt, golangci-lint)
├── .codespellrc                 # Spell-check configuration
├── .harper-dictionary.txt       # Custom dictionary for spell-check
├── AGENTS.md                    # Agent guide (entry point)
├── README.md                    # Project README
├── LICENSE                      # Project license
├── go.mod                       # Go module definition
└── go.sum                       # Module checksums
```
