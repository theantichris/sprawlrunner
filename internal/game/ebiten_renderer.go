package game

import (
	"fmt"
	"os"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
)

// EbitenRenderer handles rendering a Game using the Ebiten game engine.
type EbitenRenderer struct {
	screenWidth  int
	screenHeight int
	tileSize     int
	fontFace     font.Face
	fontSize     float64
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

// LoadFont loads a TrueType font from the given path at the specified size.
// The fontSize is in points. Returns an error if the font cannot be loaded.
func (renderer *EbitenRenderer) LoadFont(fontPath string, fontSize float64) error {
	fontData, err := os.ReadFile(fontPath)
	if err != nil {
		return fmt.Errorf("reading font file: %w", err)
	}

	ttfFont, err := truetype.Parse(fontData)
	if err != nil {
		return fmt.Errorf("parsing font: %w", err)
	}

	const dpi = 72
	renderer.fontFace = truetype.NewFace(ttfFont, &truetype.Options{
		Size:    fontSize,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})

	renderer.fontSize = fontSize

	return nil
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

// Layout returns the game's logical screen size. Required by ebiten.Game interface.
func (renderer *EbitenRenderer) Layout(outsideWidth, outsideHeight int) (int, int) {
	return renderer.screenWidth, renderer.screenHeight
}

// Close releases resources held by the renderer.
func (renderer *EbitenRenderer) Close() error {
	if renderer.fontFace != nil {
		return renderer.fontFace.Close()
	}

	return nil
}
