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
}

// NewEbitenRenderer creates a new Ebiten based renderer for the given map
// dimensions.
// mapWidth and mapHeight are in tiles, which are converted to pixels based on
// tileSize.
// fontPath specifies the TrueType font file to use and fontSize is in points.
// Returns an error if the font cannot be loaded.
func NewEbitenRenderer(mapWidth, mapHeight int, fontPath string, fontSize float64) (*EbitenRenderer, error) {
	// TODO: make package constants
	const tileSize = 16 // pixels per tile

	renderer := &EbitenRenderer{
		screenWidth:  mapWidth * tileSize,
		screenHeight: mapHeight * tileSize,
		tileSize:     tileSize,
	}

	fontData, err := os.Open(fontPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open font file: %w", err)
	}

	defer func() {
		_ = fontData.Close()
	}()

	fontSource, err := text.NewGoTextFaceSource(fontData)
	if err != nil {
		return nil, fmt.Errorf("failed to parse font: %w", err)
	}

	renderer.fontFace = &text.GoTextFace{
		Source: fontSource,
		Size:   fontSize,
	}

	return renderer, nil
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

// RenderMap draws all the tiles from the game map onto the screen.
func (renderer *EbitenRenderer) RenderMap(screen *ebiten.Image, game *Game) {
	for y := range game.Height {
		for x := range game.Width {
			tile := game.Tiles[y][x]
			renderer.RenderTile(screen, tile, x, y)
		}
	}
}

// RenderPlayer draws the player character at their current position.
func (renderer *EbitenRenderer) RenderPlayer(screen *ebiten.Image, player Player) {
	// Convert tile coordinates to pixel coordinates
	pixelX := float64(player.X * renderer.tileSize)
	pixelY := float64(player.Y * renderer.tileSize)

	glyphString := string(player.Glyph)
	options := &text.DrawOptions{}
	options.GeoM.Translate(pixelX, pixelY)
	options.ColorScale.ScaleWithColor(color.White)

	text.Draw(screen, glyphString, renderer.fontFace, options)
}
