package main

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/theantichris/sprawlrunner/internal/game"
)

const (
	initWidth    = 1280
	initHeight   = 800
	fontFacePath = "assets/fonts/Go-Mono.ttf"
	fontSize     = 16
	windowTitle  = "sprawlrunner"
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
	g := game.NewGame()
	renderer, err := game.NewEbitenRenderer(g, fontFacePath, fontSize)
	if err != nil {
		return err
	}

	ebiten.SetWindowSize(initWidth, initHeight)
	ebiten.SetWindowTitle(windowTitle)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	return ebiten.RunGame(renderer)
}
