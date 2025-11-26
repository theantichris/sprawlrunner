package game

import (
	"fmt"
	"image/color"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

// EbitenRenderer handles rendering a Game using the Ebiten game engine.
type EbitenRenderer struct {
	screenWidth  int
	screenHeight int
	tileSize     int
	fontFace     *text.GoTextFace
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
	fontData, err := os.Open(fontPath)
	if err != nil {
		return fmt.Errorf("reading font file: %w", err)
	}

	defer func() {
		_ = fontData.Close()
	}()

	fontSource, err := text.NewGoTextFaceSource(fontData)
	if err != nil {
		return fmt.Errorf("parsing font: %w", err)
	}

	renderer.fontFace = &text.GoTextFace{
		Source: fontSource,
		Size:   fontSize,
	}

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

// RenderTile draws a single tile glyph at the specified tile coordinates.
// tileX and tileY are in tile units which are converted to pixel coordinates.
func (renderer *EbitenRenderer) RenderTile(screen *ebiten.Image, tile Tile, tileX, tileY int) {
	// TODO: set fontFace in constructor
	if renderer.fontFace == nil {
		return // Can't render without a font
	}

	// Convert tile coordinates to pixel coordinates
	pixelX := float64(tileX * renderer.tileSize)
	pixelY := float64(tileY * renderer.tileSize)

	// Draw the glyph as gray text
	glyphString := string(tile.Glyph)
	options := &text.DrawOptions{}
	options.GeoM.Translate(pixelX, pixelY)
	options.ColorScale.ScaleWithColor(color.Gray{Y: 192})

	text.Draw(screen, glyphString, renderer.fontFace, options)
}
