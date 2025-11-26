package game

import "testing"

func TestTileTypes(t *testing.T) {
	tests := []struct {
		name     string
		tile     Tile
		glyph    rune
		walkable bool
	}{
		{
			name:     "floor tile",
			tile:     FloorTile,
			glyph:    '.',
			walkable: true,
		},
		{
			name:     "wall tile",
			tile:     WallTile,
			glyph:    '#',
			walkable: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.tile.Glyph != tt.glyph {
				t.Errorf("expected glyph '%c', got '%c'", tt.glyph, tt.tile.Glyph)
			}

			if tt.tile.Walkable != tt.walkable {
				t.Errorf("expected walkable to be %v, got %v", tt.walkable, tt.tile.Walkable)
			}
		})
	}
}
