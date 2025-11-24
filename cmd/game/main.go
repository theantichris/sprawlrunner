package main

import (
	"log"

	"github.com/gdamore/tcell/v2"
)

func main() {
	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("creating screen: %v", err)
	}

	if err = screen.Init(); err != nil {
		log.Fatalf("initializing screen: %v", err)
	}

	defer screen.Fini()

	// Basic neon default style: bright text on dark background
	defaultStyle := tcell.StyleDefault.
		Foreground(tcell.ColorDarkCyan).
		Background(tcell.ColorBlack)

	screen.SetStyle(defaultStyle)
	screen.Clear()

	title := "sprawlrunner (q to quit)"
	putCentered(screen, title)

	screen.Show()

	// Main event loop (just waits for quit key for now)
	for {
		event := screen.PollEvent()

		switch eventType := event.(type) {
		case *tcell.EventKey:
			switch eventType.Key() {
			case tcell.KeyEscape, tcell.KeyCtrlC:
				return
			default:
				if eventType.Rune() == 'q' || eventType.Rune() == 'Q' {
					return
				}
			}
		case *tcell.EventResize:
			screen.Clear()
			putCentered(screen, title)
			screen.Show()
		}
	}
}

// putCentered puts a string in the center of the screen.
func putCentered(screen tcell.Screen, text string) {
	width, height := screen.Size()
	x := (width - len(text)) / 2
	y := height / 2

	for index, rune := range text {
		screen.SetContent(x+index, y, rune, nil, tcell.StyleDefault)
	}
}
