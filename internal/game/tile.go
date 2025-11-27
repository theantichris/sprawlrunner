// Package game contains core game state and logic independent of rendering.
package game

import "image/color"

var (
	FloorTile = Tile{Glyph: '.', Color: color.Gray{Y: 192}, Walkable: true}
	WallTile  = Tile{Glyph: '#', Color: color.Gray{Y: 192}, Walkable: false}
)

// Tile represents a single map cell terrain in the game world.
type Tile struct {
	Glyph    rune        // Glyph is the rune used to render the tile.
	Color    color.Color // Color is the color used to render the tile (color.Gray{Y: 192}).
	Walkable bool        // Walkable indicates whether entities can move onto this tile.
}
