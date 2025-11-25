package main

import (
	"testing"

	"github.com/gdamore/tcell/v2"
	"github.com/theantichris/sprawlrunner/internal/game"
)

func TestHandleKey(t *testing.T) {
	tests := []struct {
		name     string
		key      tcell.Key
		rune     rune
		wantQuit bool
		startX   int
		startY   int
		expectX  int
		expectY  int
	}{
		{
			name:     "escape key quits",
			key:      tcell.KeyEscape,
			wantQuit: true,
			startX:   5,
			startY:   5,
			expectX:  5,
			expectY:  5,
		},
		{
			name:     "ctrl-c quits",
			key:      tcell.KeyCtrlC,
			wantQuit: true,
			startX:   5,
			startY:   5,
			expectX:  5,
			expectY:  5,
		},
		{
			name:     "q key quits",
			key:      tcell.KeyRune,
			rune:     'q',
			wantQuit: true,
			startX:   5,
			startY:   5,
			expectX:  5,
			expectY:  5,
		},
		{
			name:     "Q key quits",
			key:      tcell.KeyRune,
			rune:     'Q',
			wantQuit: true,
			startX:   5,
			startY:   5,
			expectX:  5,
			expectY:  5,
		},
		{
			name:     "up arrow moves up",
			key:      tcell.KeyUp,
			wantQuit: false,
			startX:   5,
			startY:   5,
			expectX:  5,
			expectY:  4,
		},
		{
			name:     "down arrow moves down",
			key:      tcell.KeyDown,
			wantQuit: false,
			startX:   5,
			startY:   5,
			expectX:  5,
			expectY:  6,
		},
		{
			name:     "left arrow moves left",
			key:      tcell.KeyLeft,
			wantQuit: false,
			startX:   5,
			startY:   5,
			expectX:  4,
			expectY:  5,
		},
		{
			name:     "right arrow moves right",
			key:      tcell.KeyRight,
			wantQuit: false,
			startX:   5,
			startY:   5,
			expectX:  6,
			expectY:  5,
		},
		{
			name:     "h key moves left",
			key:      tcell.KeyRune,
			rune:     'h',
			wantQuit: false,
			startX:   5,
			startY:   5,
			expectX:  4,
			expectY:  5,
		},
		{
			name:     "j key moves down",
			key:      tcell.KeyRune,
			rune:     'j',
			wantQuit: false,
			startX:   5,
			startY:   5,
			expectX:  5,
			expectY:  6,
		},
		{
			name:     "k key moves up",
			key:      tcell.KeyRune,
			rune:     'k',
			wantQuit: false,
			startX:   5,
			startY:   5,
			expectX:  5,
			expectY:  4,
		},
		{
			name:     "l key moves right",
			key:      tcell.KeyRune,
			rune:     'l',
			wantQuit: false,
			startX:   5,
			startY:   5,
			expectX:  6,
			expectY:  5,
		},
		{
			name:     "unknown key does nothing",
			key:      tcell.KeyRune,
			rune:     'x',
			wantQuit: false,
			startX:   5,
			startY:   5,
			expectX:  5,
			expectY:  5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a game with enough space to move
			g := game.NewGame(20, 20)
			g.Player.X = tt.startX
			g.Player.Y = tt.startY

			// Create event
			event := tcell.NewEventKey(tt.key, tt.rune, tcell.ModNone)

			// Call handleKey
			actualQuit := handleKey(event, g)

			// Check quit result
			if actualQuit != tt.wantQuit {
				t.Errorf("want quit %v, got %v", tt.wantQuit, actualQuit)
			}

			// Check player position
			if g.Player.X != tt.expectX {
				t.Errorf("want player X %d, got %d", tt.expectX, g.Player.X)
			}

			if g.Player.Y != tt.expectY {
				t.Errorf("want player Y %d, got %d", tt.expectY, g.Player.Y)
			}
		})
	}
}
