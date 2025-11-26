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
		logger.Fatalf("sprawl runner exited with error: %v", err)
	}
}

// run initializes the terminal screen, creates a new game, and enters the
// main event loop. It returns an error if initialization or game execution fails.
func run() error {
	screen, err := tcell.NewScreen()
	if err != nil {
		return fmt.Errorf("creating screen: %w", err)
	}

	defer func() {
		screen.Fini()

		// Always restore the terminal, even on panic.
		if r := recover(); r != nil {
			// Panic after clean up to see the stacktrace.
			panic(r)
		}
	}()

	if err = screen.Init(); err != nil {
		return fmt.Errorf("initializing screen: %w", err)
	}

	defaultStyle := tcell.StyleDefault.
		Foreground(tcell.ColorGray).
		Background(tcell.ColorBlack)

	game := initializeGame()

	screen.SetStyle(defaultStyle)
	drawGame(screen, game, defaultStyle)

	// Main event loop
	for {
		event := screen.PollEvent()

		switch eventType := event.(type) {
		case *tcell.EventKey:
			if eventType.Rune() == 'Q' {
				return nil
			}

			if dx, dy, ok := movementDeltaForKey(eventType); ok {
				game.MovePlayer(dx, dy)
				drawGame(screen, game, defaultStyle)
			}
		case *tcell.EventResize:
			drawGame(screen, game, defaultStyle)
		}
	}
}

// initializeGame creates and returns a new game based on the terminal screen
// size.
func initializeGame() *game.Game {
	return game.NewGame()
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

// movementDeltaForKey turns a keypress into a movement delta.
// Returns (dx, dy, true) if it's a movement key, otherwise (0, 0, false).
func movementDeltaForKey(event *tcell.EventKey) (int, int, bool) {
	switch event.Key() {
	case tcell.KeyUp:
		return 0, -1, true
	case tcell.KeyDown:
		return 0, 1, true
	case tcell.KeyLeft:
		return -1, 0, true
	case tcell.KeyRight:
		return 1, 0, true
	case tcell.KeyHome:
		return -1, -1, true
	case tcell.KeyPgUp:
		return 1, -1, true
	case tcell.KeyEnd:
		return -1, 1, true
	case tcell.KeyPgDn:
		return 1, 1, true
	}

	if event.Key() == tcell.KeyRune {
		switch event.Rune() {
		case 'h', '4':
			return -1, 0, true
		case 'j', '2':
			return 0, 1, true
		case 'k', '8':
			return 0, -1, true
		case 'l', '6':
			return 1, 0, true
		case 'y', '7':
			return -1, -1, true
		case 'u', '9':
			return 1, -1, true
		case 'b', '1':
			return -1, 1, true
		case 'n', '3':
			return 1, 1, true
		}
	}

	return 0, 0, false
}
