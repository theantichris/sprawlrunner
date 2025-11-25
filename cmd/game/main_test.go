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

func TestInitializeGame(t *testing.T) {
	tests := []struct {
		name           string
		screenWidth    int
		screenHeight   int
		expectedWidth  int
		expectedHeight int
	}{
		{
			name:           "standard terminal",
			screenWidth:    80,
			screenHeight:   24,
			expectedWidth:  78,
			expectedHeight: 22,
		},
		{
			name:           "large terminal",
			screenWidth:    120,
			screenHeight:   40,
			expectedWidth:  118,
			expectedHeight: 38,
		},
		{
			name:           "small terminal",
			screenWidth:    15,
			screenHeight:   15,
			expectedWidth:  13,
			expectedHeight: 13,
		},
		{
			name:           "narrow terminal enforces minimum width",
			screenWidth:    11,
			screenHeight:   30,
			expectedWidth:  10,
			expectedHeight: 28,
		},
		{
			name:           "short terminal enforces minimum height",
			screenWidth:    30,
			screenHeight:   11,
			expectedWidth:  28,
			expectedHeight: 10,
		},
		{
			name:           "tiny terminal enforces both minimums",
			screenWidth:    5,
			screenHeight:   5,
			expectedWidth:  10,
			expectedHeight: 10,
		},
		{
			name:           "minimum viable terminal",
			screenWidth:    12,
			screenHeight:   12,
			expectedWidth:  10,
			expectedHeight: 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			screen := tcell.NewSimulationScreen("UTF-8")
			err := screen.Init()

			if err != nil {
				t.Fatalf("want screen init success, got error: %v", err)
			}

			defer screen.Fini()

			screen.SetSize(tt.screenWidth, tt.screenHeight)

			game := initializeGame(screen)

			if game == nil {
				t.Fatal("want game, got nil")
			}

			if game.Width != tt.expectedWidth {
				t.Errorf("want width %d, got %d", tt.expectedWidth, game.Width)
			}

			if game.Height != tt.expectedHeight {
				t.Errorf("want height %d, got %d", tt.expectedHeight, game.Height)
			}

			// Verify game is properly initialized with tiles
			if len(game.Tiles) != game.Height {
				t.Errorf("want tiles height %d, got %d", game.Height, len(game.Tiles))
			}

			if len(game.Tiles) > 0 && len(game.Tiles[0]) != game.Width {
				t.Errorf("want tiles width %d, got %d", game.Width, len(game.Tiles[0]))
			}

			// Verify player is positioned
			if game.Player.X < 0 || game.Player.X >= game.Width {
				t.Errorf("want player X in range [0, %d), got %d", game.Width, game.Player.X)
			}

			if game.Player.Y < 0 || game.Player.Y >= game.Height {
				t.Errorf("want player Y in range [0, %d), got %d", game.Height, game.Player.Y)
			}
		})
	}
}
