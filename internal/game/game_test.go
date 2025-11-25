package game

import "testing"

func TestNewGame(t *testing.T) {
	tests := []struct {
		name   string
		width  int
		height int
	}{
		{
			name:   "standard size",
			width:  80,
			height: 24,
		},
		{
			name:   "small room",
			width:  10,
			height: 10,
		},
		{
			name:   "minimum size",
			width:  3,
			height: 3,
		},
		{
			name:   "rectangular wide",
			width:  100,
			height: 30,
		},
		{
			name:   "rectangular tall",
			width:  40,
			height: 60,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game := NewGame(tt.width, tt.height)

			// Check dimensions
			if game.Width != tt.width {
				t.Errorf("want width %d, got %d", tt.width, game.Width)
			}

			if game.Height != tt.height {
				t.Errorf("want height %d, got %d", tt.height, game.Height)
			}

			// Check tiles array is properly sized
			if len(game.Tiles) != tt.height {
				t.Errorf("want tiles height %d, got %d", tt.height, len(game.Tiles))
			}

			for y := range tt.height {
				if len(game.Tiles[y]) != tt.width {
					t.Errorf("want tiles[%d] width %d, got %d", y, tt.width, len(game.Tiles[y]))
				}
			}

			// Check borders are walls
			for x := range tt.width {
				// Top border
				if game.Tiles[0][x].Glyph != '#' {
					t.Errorf("want top border at x=%d glyph '#', got %c", x, game.Tiles[0][x].Glyph)
				}

				if game.Tiles[0][x].Walkable {
					t.Errorf("want top border at x=%d walkable false, got true", x)
				}

				// Bottom border
				if game.Tiles[tt.height-1][x].Glyph != '#' {
					t.Errorf("want bottom border at x=%d glyph '#', got %c", x, game.Tiles[tt.height-1][x].Glyph)
				}

				if game.Tiles[tt.height-1][x].Walkable {
					t.Errorf("want bottom border at x=%d walkable false, got true", x)
				}
			}

			for y := range tt.height {
				// Left border
				if game.Tiles[y][0].Glyph != '#' {
					t.Errorf("want left border at y=%d glyph '#', got %c", y, game.Tiles[y][0].Glyph)
				}

				if game.Tiles[y][0].Walkable {
					t.Errorf("want left border at y=%d walkable false, got true", y)
				}

				// Right border
				if game.Tiles[y][tt.width-1].Glyph != '#' {
					t.Errorf("want right border at y=%d glyph '#', got %c", y, game.Tiles[y][tt.width-1].Glyph)
				}

				if game.Tiles[y][tt.width-1].Walkable {
					t.Errorf("want right border at y=%d walkable false, got true", y)
				}
			}

			// Check interior tiles are floors (if room is large enough)
			if tt.width > 2 && tt.height > 2 {
				centerX := tt.width / 2
				centerY := tt.height / 2

				if game.Tiles[centerY][centerX].Glyph != '.' {
					t.Errorf("want center tile glyph '.', got %c", game.Tiles[centerY][centerX].Glyph)
				}

				if !game.Tiles[centerY][centerX].Walkable {
					t.Errorf("want center tile walkable true, got false")
				}
			}

			// Check player position
			expectedX := tt.width / 2
			expectedY := tt.height / 2

			if game.Player.X != expectedX {
				t.Errorf("want player X %d, got %d", expectedX, game.Player.X)
			}

			if game.Player.Y != expectedY {
				t.Errorf("want player Y %d, got %d", expectedY, game.Player.Y)
			}

			// Check player glyph
			if game.Player.Glyph != '@' {
				t.Errorf("want player glyph '@', got %c", game.Player.Glyph)
			}
		})
	}
}

func TestMovePlayer(t *testing.T) {
	tests := []struct {
		name      string
		width     int
		height    int
		startX    int
		startY    int
		dx        int
		dy        int
		expectX   int
		expectY   int
		wantMoved bool
	}{
		{
			name:      "move right on floor",
			width:     10,
			height:    10,
			startX:    5,
			startY:    5,
			dx:        1,
			dy:        0,
			expectX:   6,
			expectY:   5,
			wantMoved: true,
		},
		{
			name:      "move left on floor",
			width:     10,
			height:    10,
			startX:    5,
			startY:    5,
			dx:        -1,
			dy:        0,
			expectX:   4,
			expectY:   5,
			wantMoved: true,
		},
		{
			name:      "move down on floor",
			width:     10,
			height:    10,
			startX:    5,
			startY:    5,
			dx:        0,
			dy:        1,
			expectX:   5,
			expectY:   6,
			wantMoved: true,
		},
		{
			name:      "move up on floor",
			width:     10,
			height:    10,
			startX:    5,
			startY:    5,
			dx:        0,
			dy:        -1,
			expectX:   5,
			expectY:   4,
			wantMoved: true,
		},
		{
			name:      "blocked by top wall",
			width:     10,
			height:    10,
			startX:    5,
			startY:    1,
			dx:        0,
			dy:        -1,
			expectX:   5,
			expectY:   1,
			wantMoved: false,
		},
		{
			name:      "blocked by bottom wall",
			width:     10,
			height:    10,
			startX:    5,
			startY:    8,
			dx:        0,
			dy:        1,
			expectX:   5,
			expectY:   8,
			wantMoved: false,
		},
		{
			name:      "blocked by left wall",
			width:     10,
			height:    10,
			startX:    1,
			startY:    5,
			dx:        -1,
			dy:        0,
			expectX:   1,
			expectY:   5,
			wantMoved: false,
		},
		{
			name:      "blocked by right wall",
			width:     10,
			height:    10,
			startX:    8,
			startY:    5,
			dx:        1,
			dy:        0,
			expectX:   8,
			expectY:   5,
			wantMoved: false,
		},
		{
			name:      "blocked by left edge of map",
			width:     10,
			height:    10,
			startX:    0,
			startY:    5,
			dx:        -1,
			dy:        0,
			expectX:   0,
			expectY:   5,
			wantMoved: false,
		},
		{
			name:      "blocked by right edge of map",
			width:     10,
			height:    10,
			startX:    9,
			startY:    5,
			dx:        1,
			dy:        0,
			expectX:   9,
			expectY:   5,
			wantMoved: false,
		},
		{
			name:      "blocked by top edge of map",
			width:     10,
			height:    10,
			startX:    5,
			startY:    0,
			dx:        0,
			dy:        -1,
			expectX:   5,
			expectY:   0,
			wantMoved: false,
		},
		{
			name:      "blocked by bottom edge of map",
			width:     10,
			height:    10,
			startX:    5,
			startY:    9,
			dx:        0,
			dy:        1,
			expectX:   5,
			expectY:   9,
			wantMoved: false,
		},
		{
			name:      "no movement",
			width:     10,
			height:    10,
			startX:    5,
			startY:    5,
			dx:        0,
			dy:        0,
			expectX:   5,
			expectY:   5,
			wantMoved: false,
		},
		{
			name:      "large step blocked by edge",
			width:     10,
			height:    10,
			startX:    5,
			startY:    5,
			dx:        10,
			dy:        0,
			expectX:   5,
			expectY:   5,
			wantMoved: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game := NewGame(tt.width, tt.height)
			game.Player.X = tt.startX
			game.Player.Y = tt.startY

			game.MovePlayer(tt.dx, tt.dy)

			if game.Player.X != tt.expectX {
				t.Errorf("want player X %d, got %d", tt.expectX, game.Player.X)
			}

			if game.Player.Y != tt.expectY {
				t.Errorf("want player Y %d, got %d", tt.expectY, game.Player.Y)
			}

			moved := game.Player.X != tt.startX || game.Player.Y != tt.startY

			if moved != tt.wantMoved {
				t.Errorf("want player moved %v, got %v", tt.wantMoved, moved)
			}
		})
	}
}
