package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/log"
	"github.com/gdamore/tcell/v2"
	"github.com/theantichris/sprawlrunner/internal/game"
)

// main is the entry point for the Sprawlrunner game binary. It initializes
// the logger and runs the game, exiting with an error if something goes wrong.
func main() {
	logger := log.NewWithOptions(os.Stderr, log.Options{
		Formatter:       log.JSONFormatter,
		ReportCaller:    true,
		ReportTimestamp: true,
		Level:           log.DebugLevel,
	})

	if err := run(); err != nil {
		logger.Fatalf("error running game: %v", err)
	}
}

// run initializes the terminal screen, creates a new game, and enters the
// main event loop. It returns an error if initialization or game execution fails.
func run() error {
	screen, err := tcell.NewScreen()
	if err != nil {
		return fmt.Errorf("creating screen: %w", err)
	}

	defer screen.Fini()

	if err = screen.Init(); err != nil {
		return fmt.Errorf("initializing screen: %w", err)
	}

	defaultStyle := tcell.StyleDefault.
		Foreground(tcell.ColorGray).
		Background(tcell.ColorBlack)

	game := initializeGame(screen)

	screen.SetStyle(defaultStyle)
	drawGame(screen, game, defaultStyle)

	// Main event loop
	for {
		event := screen.PollEvent()

		switch eventType := event.(type) {
		case *tcell.EventKey:
			if handleKey(eventType, game) {
				return nil // true means quit requested
			}

			drawGame(screen, game, defaultStyle)
		case *tcell.EventResize:
			drawGame(screen, game, defaultStyle)
		}
	}
}

// initializeGame creates and returns a new game based on the terminal screen
// size.
func initializeGame(screen tcell.Screen) *game.Game {
	width, height := screen.Size()

	// Derive map size from terminal, with a minimum size for safety.
	mapWidth := width - 2
	mapHeight := height - 2

	if mapWidth < 10 {
		mapWidth = 10
	}

	if mapHeight < 10 {
		mapHeight = 10
	}

	return game.NewGame(mapWidth, mapHeight)
}

// handleKey processes a single key event and updates game state. It returns
// true if the caller should quit the game.
func handleKey(event *tcell.EventKey, game *game.Game) bool {
	switch event.Key() {
	case tcell.KeyEscape, tcell.KeyCtrlC:
		return true

	case tcell.KeyUp:
		game.MovePlayer(0, -1)

	case tcell.KeyDown:
		game.MovePlayer(0, 1)

	case tcell.KeyLeft:
		game.MovePlayer(-1, 0)

	case tcell.KeyRight:
		game.MovePlayer(1, 0)
	default:
		switch event.Rune() {
		case 'q', 'Q':
			return true

		case 'h':
			game.MovePlayer(-1, 0)

		case 'j':
			game.MovePlayer(0, 1)

		case 'k':
			game.MovePlayer(0, -1)
		case 'l':
			game.MovePlayer(1, 0)
		}
	}

	return false
}

// drawGame renders the current game state (map and player) onto the given
// screen using the provided base style for tiles.
func drawGame(screen tcell.Screen, game *game.Game, style tcell.Style) {
	screen.Clear()

	// Offset map slightly from the top left for a cleaner look.
	const offsetX = 1
	const offsetY = 1

	// Draw tiles.
	for y := 0; y < game.Height; y++ {
		for x := 0; x < game.Width; x++ {
			tile := game.Tiles[y][x]
			screen.SetContent(offsetX+x, offsetY+y, tile.Glyph, nil, style)
		}
	}

	// Draw player with a distinct style to stand out.
	playerStyle := style.Foreground(tcell.ColorWhite)
	screen.SetContent(
		offsetX+game.Player.X,
		offsetY+game.Player.Y,
		game.Player.Glyph,
		nil,
		playerStyle,
	)

	screen.Show()
}
