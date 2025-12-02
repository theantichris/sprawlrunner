package game

import "testing"

func TestNewGame(t *testing.T) {
	game := NewGame()

	if game.Width != 80 {
		t.Errorf("want width 80, got %d", game.Width)
	}

	if game.Height != 24 {
		t.Errorf("want height 24, got %d", game.Height)
	}

	if len(game.Tiles) != 24 {
		t.Errorf("want tiles height 24, got %d", len(game.Tiles))
	}

	for y := range 24 {
		if len(game.Tiles[y]) != 80 {
			t.Errorf("want tiles[%d] width 80, got %d", y, len(game.Tiles[y]))
		}
	}

	// Spot check Room 1 (at 10, 5, size 15x8)
	// Check a few tiles inside Room 1 are floors
	if game.Tiles[7][15].Glyph != '.' {
		t.Errorf("want Room 1 interior tile glyph '.', got %c", game.Tiles[7][15].Glyph)
	}
	if !game.Tiles[7][15].Walkable {
		t.Error("want Room 1 interior tile walkable true, got false")
	}

	// Spot check Room 2 (at 35, 3, size 12x10)
	if game.Tiles[8][41].Glyph != '.' {
		t.Errorf("want Room 2 interior tile glyph '.', got %c", game.Tiles[8][41].Glyph)
	}
	if !game.Tiles[8][41].Walkable {
		t.Error("want Room 2 interior tile walkable true, got false")
	}

	// Spot check Room 3 (at 55, 12, size 18x9)
	if game.Tiles[16][64].Glyph != '.' {
		t.Errorf("want Room 3 interior tile glyph '.', got %c", game.Tiles[16][64].Glyph)
	}
	if !game.Tiles[16][64].Walkable {
		t.Error("want Room 3 interior tile walkable true, got false")
	}

	// Check a tile that should be a wall (outside rooms/corridors)
	if game.Tiles[0][0].Glyph != '#' {
		t.Errorf("want wall tile glyph '#', got %c", game.Tiles[0][0].Glyph)
	}
	if game.Tiles[0][0].Walkable {
		t.Error("want wall tile walkable false, got true")
	}

	// Check player position (center of Room 1)
	if game.Player.X != 17 {
		t.Errorf("want player X 17, got %d", game.Player.X)
	}

	if game.Player.Y != 9 {
		t.Errorf("want player Y 9, got %d", game.Player.Y)
	}

	// Check player glyph
	if game.Player.Glyph != '@' {
		t.Errorf("want player glyph '@', got %c", game.Player.Glyph)
	}
}

func TestMovePlayer(t *testing.T) {
	tests := []struct {
		name      string
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
			startX:    15, // Inside Room 1
			startY:    9,
			dx:        1,
			dy:        0,
			expectX:   16,
			expectY:   9,
			wantMoved: true,
		},
		{
			name:      "move left on floor",
			startX:    15, // Inside Room 1
			startY:    9,
			dx:        -1,
			dy:        0,
			expectX:   14,
			expectY:   9,
			wantMoved: true,
		},
		{
			name:      "move down on floor",
			startX:    15, // Inside Room 1
			startY:    9,
			dx:        0,
			dy:        1,
			expectX:   15,
			expectY:   10,
			wantMoved: true,
		},
		{
			name:      "move up on floor",
			startX:    15, // Inside Room 1
			startY:    9,
			dx:        0,
			dy:        -1,
			expectX:   15,
			expectY:   8,
			wantMoved: true,
		},
		{
			name:      "blocked by wall at room edge",
			startX:    10, // Left edge of Room 1
			startY:    9,
			dx:        -1,
			dy:        0,
			expectX:   10,
			expectY:   9,
			wantMoved: false,
		},
		{
			name:      "blocked by wall at top of room",
			startX:    15, // Inside Room 1
			startY:    5,  // Top edge
			dx:        0,
			dy:        -1,
			expectX:   15,
			expectY:   5,
			wantMoved: false,
		},
		{
			name:      "blocked by left edge of map",
			startX:    0,
			startY:    9,
			dx:        -1,
			dy:        0,
			expectX:   0,
			expectY:   9,
			wantMoved: false,
		},
		{
			name:      "blocked by right edge of map",
			startX:    79,
			startY:    9,
			dx:        1,
			dy:        0,
			expectX:   79,
			expectY:   9,
			wantMoved: false,
		},
		{
			name:      "blocked by top edge of map",
			startX:    15,
			startY:    0,
			dx:        0,
			dy:        -1,
			expectX:   15,
			expectY:   0,
			wantMoved: false,
		},
		{
			name:      "blocked by bottom edge of map",
			startX:    15,
			startY:    23,
			dx:        0,
			dy:        1,
			expectX:   15,
			expectY:   23,
			wantMoved: false,
		},
		{
			name:      "no movement",
			startX:    15,
			startY:    9,
			dx:        0,
			dy:        0,
			expectX:   15,
			expectY:   9,
			wantMoved: false,
		},
		{
			name:      "large step blocked by edge",
			startX:    15,
			startY:    9,
			dx:        100,
			dy:        0,
			expectX:   15,
			expectY:   9,
			wantMoved: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game := NewGame()
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

func TestCreateRoom(t *testing.T) {
	game := &Game{
		Width:  20,
		Height: 15,
		Tiles:  make([][]Tile, 15),
	}

	// Initialize all tiles as walls
	for y := range 15 {
		row := make([]Tile, 20)
		for x := range 20 {
			row[x] = WallTile
		}
		game.Tiles[y] = row
	}

	// Create a room at (5, 3) with size 8x6
	game.CreateRoom(5, 3, 8, 6)

	// Verify tiles inside the room are floors
	for y := 3; y < 9; y++ { // 3 to 8 (3+6-1)
		for x := 5; x < 13; x++ { // 5 to 12 (5+8-1)
			if game.Tiles[y][x] != FloorTile {
				t.Errorf("want tile at (%d,%d) to be FloorTile, got %v", x, y, game.Tiles[y][x])
			}
		}
	}

	// Verify tiles outside the room are still walls
	// Check a tile before the room
	if game.Tiles[3][4] != WallTile {
		t.Errorf("want tile at (4,3) to be WallTile, got %v", game.Tiles[3][4])
	}

	// Check a tile after the room
	if game.Tiles[3][13] != WallTile {
		t.Errorf("want tile at (13,3) to be WallTile, got %v", game.Tiles[3][13])
	}
}

func TestCreateCorridor(t *testing.T) {
	tests := []struct {
		name string
		x1   int
		y1   int
		x2   int
		y2   int
	}{
		{
			name: "corridor right and down",
			x1:   5,
			y1:   5,
			x2:   10,
			y2:   10,
		},
		{
			name: "corridor left and up",
			x1:   10,
			y1:   10,
			x2:   5,
			y2:   5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game := &Game{
				Width:  20,
				Height: 20,
				Tiles:  make([][]Tile, 20),
			}

			// Initialize all tiles as walls
			for y := range 20 {
				row := make([]Tile, 20)
				for x := range 20 {
					row[x] = WallTile
				}
				game.Tiles[y] = row
			}

			// Create the corridor
			game.CreateCorridor(tt.x1, tt.y1, tt.x2, tt.y2)

			// Verify horizontal segment is carved
			for x := min(tt.x1, tt.x2); x <= max(tt.x1, tt.x2); x++ {
				if game.Tiles[tt.y1][x] != FloorTile {
					t.Errorf("want horizontal corridor tile at (%d,%d) to be FloorTile, got %v", x, tt.y1, game.Tiles[tt.y1][x])
				}
			}

			// Verify vertical segment is carved
			for y := min(tt.y1, tt.y2); y <= max(tt.y1, tt.y2); y++ {
				if game.Tiles[y][tt.x2] != FloorTile {
					t.Errorf("want vertical corridor tile at (%d,%d) to be FloorTile, got %v", tt.x2, y, game.Tiles[y][tt.x2])
				}
			}

			// Verify a tile outside the corridor is still a wall
			if game.Tiles[0][0] != WallTile {
				t.Errorf("want tile at (0,0) to be WallTile, got %v", game.Tiles[0][0])
			}
		})
	}
}

func TestQuitConfirmation(t *testing.T) {
	game := NewGame()

	if game.IsConfirmingQuit() {
		t.Fatal("new game should not be confirming quit")
	}

	game.RequestQuit()

	if !game.IsConfirmingQuit() {
		t.Error("game should be confirming quit after RequestQuit()")
	}
}

func TestConfirmQuitExitsGame(t *testing.T) {
	game := NewGame()
	game.RequestQuit()

	shouldQuit := game.ConfirmQuit(true)

	if !shouldQuit {
		t.Error("ConfirmQuit(true) should return true to exit, got false")
	}

	if game.IsConfirmingQuit() {
		t.Error("game should not be confirming quit after confirmation")
	}
}

func TestCancelQuitReturnsToGame(t *testing.T) {
	game := NewGame()
	game.RequestQuit()

	shouldQuit := game.ConfirmQuit(false)

	if shouldQuit {
		t.Error("ConfirmQuit(false) should return false to stay in game")
	}

	if game.IsConfirmingQuit() {
		t.Error("game should not be confirming quit after cancellation")
	}
}
