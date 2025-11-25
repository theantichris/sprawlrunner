// Package game contains core game state and logic independent of rendering.
package game

// Player represents the runner controlled by the user.
type Player struct {
	X     int  // X is the player's horizontal position in tile coordinates.
	Y     int  // Y is the player's vertical position in tile coordinates.
	Glyph rune // Glyph is the rune used to render the player in the terminal.
}
