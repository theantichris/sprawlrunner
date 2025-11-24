// Package game contains core game state and logic independent of rendering.
package game

// Tile represents a single map cell in the game world.
type Tile struct {
	Glyph    rune // Glyph is the rune used to render the tile in the terminal.
	Walkable bool // Walkable indicates whether entities can move onto this tile.
}

// Player represents the runner controlled by the user.
type Player struct {
	X     int  // X is the player's horizontal position in tile coordinates.
	Y     int  // Y is the player's vertical position in tile coordinates.
	Glyph rune // Glyph is the rune used to render the player in the terminal.
}

// Game holds the current game state including map and entities.
type Game struct {
	Width  int      // Width describes the horizontal map dimensions in tiles.
	Height int      // Height describes the vertical map dimensions in tiles.
	Tiles  [][]Tile // Tiles is a 2D grid of map tiles indexed as Tiles[y][x].
	Player Player   // Player represents the runner controlled by the user.
}

// NewGame creates a new Game with a single rectangular room and a player
// positioned near the center.
func NewGame(width, height int) *Game {
	game := &Game{
		Width:  width,
		Height: height,
		Tiles:  make([][]Tile, height),
		Player: Player{
			Glyph: '@',
		},
	}

	// Create room.
	for y := range height {
		row := make([]Tile, width)

		for x := range width {
			isBorder := x == 0 || y == 0 || x == width-1 || y == height-1

			if isBorder {
				// Wall
				row[x] = Tile{
					Glyph:    '#',
					Walkable: false,
				}
			} else {
				// Floor
				row[x] = Tile{
					Glyph:    '.',
					Walkable: true,
				}
			}
		}

		game.Tiles[y] = row
	}

	// Start player roughly in the middle.
	game.Player.X = width / 2
	game.Player.Y = height / 2

	return game
}

// MovePlayer attempts to move the player by (dx, dy). The move only succeeds
// if the target tile is inside the map and is walkable.
func (game *Game) MovePlayer(dx, dy int) {
	newX := game.Player.X + dx
	newY := game.Player.Y + dy

	// Prevent player from moving off the map.
	if newX < 0 || newX >= game.Width || newY < 0 || newY >= game.Height {
		return
	}

	// Prevent player from moving into walls.
	target := game.Tiles[newY][newX]
	if !target.Walkable {
		return
	}

	game.Player.X = newX
	game.Player.Y = newY
}
