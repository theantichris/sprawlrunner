// Package game contains core game state and logic independent of rendering.
package game

// Tile represents a single map cell in the game world.
type Tile struct {
	Glyph    rune // Glyph is the rune used to render the tile in the terminal.
	Walkable bool // Walkable indicates whether entities can move onto this tile.
}
