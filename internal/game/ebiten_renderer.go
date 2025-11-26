package game

import "github.com/hajimehoshi/ebiten/v2"

// EbitenRenderer handles rendering a Game using the Ebiten game engine.
type EbitenRenderer struct {
	screenWidth  int
	screenHeight int
	tileSize     int
}

// NewEbitenRenderer creates a new Ebiten based renderer for the given map
// dimensions.
// mapWidth and mapHeight are in tiles, which are converted to pixels based on
// tileSize.
func NewEbitenRenderer(mapWidth, mapHeight int) *EbitenRenderer {
	const tileSize = 16 // pixels per tile

	return &EbitenRenderer{
		screenWidth:  mapWidth * tileSize,
		screenHeight: mapHeight * tileSize,
		tileSize:     tileSize,
	}
}

// Update updates the game state. Required by ebiten.Game interface.
// Returns error if the game should terminate.
func (renderer *EbitenRenderer) Update() error {
	// TODO: Handle input and update game state.

	return nil
}

// Draw renders the game state to the screen. Required by ebiten.Game interface.
func (renderer *EbitenRenderer) Draw(screen *ebiten.Image) {
	// TODO: Render tiles and player
}

// Layout returns the game's logical screensize. Required by ebiten.Game interface.
func (renderer *EbitenRenderer) Layout(outsideWidth, outsideHeight int) (int, int) {
	return renderer.screenWidth, renderer.screenHeight
}
