package main

import (
	"testing"

	"github.com/gdamore/tcell/v2"
)

func TestInitializeGame(t *testing.T) {
	screen := tcell.NewSimulationScreen("UTF-8")
	err := screen.Init()

	if err != nil {
		t.Fatalf("want screen init success, got error: %v", err)
	}

	defer screen.Fini()

	screen.SetSize(80, 24)

	game := initializeGame()

	if game == nil {
		t.Fatal("want game, got nil")
	}

	if game.Width != 80 {
		t.Errorf("want width 80, got %d", game.Width)
	}

	if game.Height != 24 {
		t.Errorf("want height 24, got %d", game.Height)
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
}

func TestMovementDeltaForKey(t *testing.T) {
	tests := []struct {
		name   string
		event  *tcell.EventKey
		wantDx int
		wantDy int
		wantOk bool
	}{
		// Arrow keys
		{
			name:   "up arrow moves up",
			event:  tcell.NewEventKey(tcell.KeyUp, 0, tcell.ModNone),
			wantDx: 0,
			wantDy: -1,
			wantOk: true,
		},
		{
			name:   "down arrow moves down",
			event:  tcell.NewEventKey(tcell.KeyDown, 0, tcell.ModNone),
			wantDx: 0,
			wantDy: 1,
			wantOk: true,
		},
		{
			name:   "left arrow moves left",
			event:  tcell.NewEventKey(tcell.KeyLeft, 0, tcell.ModNone),
			wantDx: -1,
			wantDy: 0,
			wantOk: true,
		},
		{
			name:   "right arrow moves right",
			event:  tcell.NewEventKey(tcell.KeyRight, 0, tcell.ModNone),
			wantDx: 1,
			wantDy: 0,
			wantOk: true,
		},

		// Special navigation keys
		{
			name:   "home moves diagonal up-left",
			event:  tcell.NewEventKey(tcell.KeyHome, 0, tcell.ModNone),
			wantDx: -1,
			wantDy: -1,
			wantOk: true,
		},
		{
			name:   "pgup moves diagonal up-right",
			event:  tcell.NewEventKey(tcell.KeyPgUp, 0, tcell.ModNone),
			wantDx: 1,
			wantDy: -1,
			wantOk: true,
		},
		{
			name:   "end moves diagonal down-left",
			event:  tcell.NewEventKey(tcell.KeyEnd, 0, tcell.ModNone),
			wantDx: -1,
			wantDy: 1,
			wantOk: true,
		},
		{
			name:   "pgdn moves diagonal down-right",
			event:  tcell.NewEventKey(tcell.KeyPgDn, 0, tcell.ModNone),
			wantDx: 1,
			wantDy: 1,
			wantOk: true,
		},

		// Vi-style keys
		{
			name:   "h moves left",
			event:  tcell.NewEventKey(tcell.KeyRune, 'h', tcell.ModNone),
			wantDx: -1,
			wantDy: 0,
			wantOk: true,
		},
		{
			name:   "j moves down",
			event:  tcell.NewEventKey(tcell.KeyRune, 'j', tcell.ModNone),
			wantDx: 0,
			wantDy: 1,
			wantOk: true,
		},
		{
			name:   "k moves up",
			event:  tcell.NewEventKey(tcell.KeyRune, 'k', tcell.ModNone),
			wantDx: 0,
			wantDy: -1,
			wantOk: true,
		},
		{
			name:   "l moves right",
			event:  tcell.NewEventKey(tcell.KeyRune, 'l', tcell.ModNone),
			wantDx: 1,
			wantDy: 0,
			wantOk: true,
		},
		{
			name:   "y moves diagonal up-left",
			event:  tcell.NewEventKey(tcell.KeyRune, 'y', tcell.ModNone),
			wantDx: -1,
			wantDy: -1,
			wantOk: true,
		},
		{
			name:   "u moves diagonal up-right",
			event:  tcell.NewEventKey(tcell.KeyRune, 'u', tcell.ModNone),
			wantDx: 1,
			wantDy: -1,
			wantOk: true,
		},
		{
			name:   "b moves diagonal down-left",
			event:  tcell.NewEventKey(tcell.KeyRune, 'b', tcell.ModNone),
			wantDx: -1,
			wantDy: 1,
			wantOk: true,
		},
		{
			name:   "n moves diagonal down-right",
			event:  tcell.NewEventKey(tcell.KeyRune, 'n', tcell.ModNone),
			wantDx: 1,
			wantDy: 1,
			wantOk: true,
		},

		// Numpad keys
		{
			name:   "4 moves left",
			event:  tcell.NewEventKey(tcell.KeyRune, '4', tcell.ModNone),
			wantDx: -1,
			wantDy: 0,
			wantOk: true,
		},
		{
			name:   "2 moves down",
			event:  tcell.NewEventKey(tcell.KeyRune, '2', tcell.ModNone),
			wantDx: 0,
			wantDy: 1,
			wantOk: true,
		},
		{
			name:   "8 moves up",
			event:  tcell.NewEventKey(tcell.KeyRune, '8', tcell.ModNone),
			wantDx: 0,
			wantDy: -1,
			wantOk: true,
		},
		{
			name:   "6 moves right",
			event:  tcell.NewEventKey(tcell.KeyRune, '6', tcell.ModNone),
			wantDx: 1,
			wantDy: 0,
			wantOk: true,
		},
		{
			name:   "7 moves diagonal up-left",
			event:  tcell.NewEventKey(tcell.KeyRune, '7', tcell.ModNone),
			wantDx: -1,
			wantDy: -1,
			wantOk: true,
		},
		{
			name:   "9 moves diagonal up-right",
			event:  tcell.NewEventKey(tcell.KeyRune, '9', tcell.ModNone),
			wantDx: 1,
			wantDy: -1,
			wantOk: true,
		},
		{
			name:   "1 moves diagonal down-left",
			event:  tcell.NewEventKey(tcell.KeyRune, '1', tcell.ModNone),
			wantDx: -1,
			wantDy: 1,
			wantOk: true,
		},
		{
			name:   "3 moves diagonal down-right",
			event:  tcell.NewEventKey(tcell.KeyRune, '3', tcell.ModNone),
			wantDx: 1,
			wantDy: 1,
			wantOk: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDx, gotDy, gotOk := movementDeltaForKey(tt.event)

			if gotDx != tt.wantDx {
				t.Errorf("want dx %d, got %d", tt.wantDx, gotDx)
			}

			if gotDy != tt.wantDy {
				t.Errorf("want dy %d, got %d", tt.wantDy, gotDy)
			}

			if gotOk != tt.wantOk {
				t.Errorf("want ok %v, got %v", tt.wantOk, gotOk)
			}
		})
	}
}
