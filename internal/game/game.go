// Package game contains core game state and logic independent of rendering.
package game

import "image/color"

const (
	mapWidth  = 80
	mapHeight = 24
)

// Game holds the current game state including map and entities.
type Game struct {
	Width          int      // Width describes the horizontal map dimensions in tiles
	Height         int      // Height describes the vertical map dimensions in tiles
	Tiles          [][]Tile // Tiles is a 2D grid of map tiles indexed as Tiles[y][x]
	Player         Player   // Player represents the runner controlled by the user
	CameraX        int      // CameraX is the camera's center position (horizontal)
	CameraY        int      // CameraY is the camera's center position (vertical)
	confirmingQuit bool     // confirmingQuit tracks whether the game is waiting for quit confirmation.
}

// NewGame creates a new Game with three rooms connected by corridors.
// The game map is fixed at 80x24 to match the hardcoded room layout.
func NewGame() *Game {
	game := &Game{
		Width:  mapWidth,
		Height: mapHeight,
		Tiles:  make([][]Tile, mapHeight),
		Player: Player{
			Glyph:  '@',
			Color:  color.White,
			Name:   "Decker",
			Level:  1,
			Health: 100,
		},
	}

	game.initializeMap(mapWidth, mapHeight)

	return game
}

// initializeMap creates 3 hardcoded rooms with corridors and starts the player
// in the center of room 1.
func (game *Game) initializeMap(width, height int) {
	// Initialize all tiles as walls
	for y := range height {
		row := make([]Tile, width)

		for x := range width {
			row[x] = WallTile
		}

		game.Tiles[y] = row
	}

	// Create rooms
	game.CreateRoom(10, 5, 15, 8)
	game.CreateRoom(35, 3, 12, 10)
	game.CreateRoom(55, 12, 18, 9)

	// Create corridors
	game.CreateCorridor(17, 9, 41, 8)
	game.CreateCorridor(41, 8, 64, 16)

	// Start player in center of first room
	game.Player.X = 17
	game.Player.Y = 9

	// Center camera on player
	game.CameraX = game.Player.X
	game.CameraY = game.Player.Y
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

	game.CameraX = newX
	game.CameraY = newY
}

// CreateRoom creates a room at x, y with the specified dimensions.
func (game *Game) CreateRoom(x, y, width, height int) {
	for yPos := y; yPos < y+height; yPos++ {
		for xPos := x; xPos < x+width; xPos++ {
			game.Tiles[yPos][xPos] = FloorTile
		}
	}
}

// CreateCorridor creates a corridor between two points horizontally then vertically.
func (game *Game) CreateCorridor(x1, y1, x2, y2 int) {
	// Horizontal segment
	for x := min(x1, x2); x <= max(x1, x2); x++ {
		game.Tiles[y1][x] = FloorTile
	}

	// Vertical segment
	for y := min(y1, y2); y <= max(y1, y2); y++ {
		game.Tiles[y][x2] = FloorTile
	}
}

// RequestQuit enters the quit confirmation state.
func (game *Game) RequestQuit() {
	game.confirmingQuit = true
}

// IsConfirmingQuit returns true if the game is waiting for quit confirmation.
func (game *Game) IsConfirmingQuit() bool {
	return game.confirmingQuit
}

// ConfirmQuit handles the user's response to the quit confirmation.
// Returns true if the game should exit, false otherwise.
func (game *Game) ConfirmQuit(confirmed bool) bool {
	game.confirmingQuit = false

	return confirmed
}
