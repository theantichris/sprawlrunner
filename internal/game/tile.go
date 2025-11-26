// Package game contains core game state and logic independent of rendering.
package game

var (
	FloorTile = Tile{Glyph: '.', Walkable: true}
	WallTile  = Tile{Glyph: '#', Walkable: false}
)

// Tile represents a single map cell terrain in the game world.
type Tile struct {
	Glyph    rune // Glyph is the rune used to render the tile in the terminal.
	Walkable bool // Walkable indicates whether entities can move onto this tile.
}
