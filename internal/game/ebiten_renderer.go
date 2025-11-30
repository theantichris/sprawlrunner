package game

import (
	"fmt"
	"image/color"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

const (
	tileSize   = 16 // 16 pixels per tile
	fontGoMono = "../../assets/fonts/Go-Mono.ttf"
)

// EbitenRenderer handles rendering a Game using the Ebiten game engine.
type EbitenRenderer struct {
	screenWidth  int
	screenHeight int
	tileSize     int
	fontFace     *text.GoTextFace
	game         *Game
}

// NewEbitenRenderer creates a new Ebiten renderer for the given game.
// Screen dimensions are derived from the game's map size.
// fontPath specifies the TrueType font file to use and fontSize is in points.
// Returns an error if font cannot be loaded.
func NewEbitenRenderer(game *Game, fontPath string, fontSize float64) (*EbitenRenderer, error) {
	renderer := &EbitenRenderer{
		screenWidth:  game.Width * tileSize,
		screenHeight: game.Height * tileSize,
		tileSize:     tileSize,
		game:         game,
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
	// Up
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) || ebiten.IsKeyPressed(ebiten.KeyK) || ebiten.IsKeyPressed(ebiten.KeyNumpad8) {
		renderer.game.MovePlayer(0, -1)
	}

	// Down
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) || ebiten.IsKeyPressed(ebiten.KeyJ) || ebiten.IsKeyPressed(ebiten.KeyNumpad2) {
		renderer.game.MovePlayer(0, 1)
	}

	// Left
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) || ebiten.IsKeyPressed(ebiten.KeyH) || ebiten.IsKeyPressed(ebiten.KeyNumpad4) {
		renderer.game.MovePlayer(-1, 0)
	}

	// Right
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) || ebiten.IsKeyPressed(ebiten.KeyL) || ebiten.IsKeyPressed(ebiten.KeyNumpad6) {
		renderer.game.MovePlayer(1, 0)
	}

	// Up left
	if ebiten.IsKeyPressed(ebiten.KeyHome) || ebiten.IsKeyPressed(ebiten.KeyY) || ebiten.IsKeyPressed(ebiten.KeyNumpad7) {
		renderer.game.MovePlayer(-1, -1)
	}

	// Up right
	if ebiten.IsKeyPressed(ebiten.KeyPageUp) || ebiten.IsKeyPressed(ebiten.KeyU) || ebiten.IsKeyPressed(ebiten.KeyNumpad9) {
		renderer.game.MovePlayer(1, -1)
	}

	// Down left
	if ebiten.IsKeyPressed(ebiten.KeyEnd) || ebiten.IsKeyPressed(ebiten.KeyB) || ebiten.IsKeyPressed(ebiten.KeyNumpad1) {
		renderer.game.MovePlayer(-1, 1)
	}

	// Down right
	if ebiten.IsKeyPressed(ebiten.KeyPageDown) || ebiten.IsKeyPressed(ebiten.KeyN) || ebiten.IsKeyPressed(ebiten.KeyNumpad3) {
		renderer.game.MovePlayer(1, 1)
	}

	// Quit
	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		return ebiten.Termination
	}

	return nil
}

// Draw renders the game state to the screen. Required by ebiten.Game interface.
func (renderer *EbitenRenderer) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black) // Clear screen to black
	renderer.RenderMap(screen, renderer.game)
	renderer.RenderPlayer(screen, renderer.game.Player)
}

// Layout returns the game's logical screen size. Required by ebiten.Game interface.
func (renderer *EbitenRenderer) Layout(outsideWidth, outsideHeight int) (int, int) {
	return renderer.screenWidth, renderer.screenHeight
}

// renderGlyph draws a single character glyph at the specified position with the given color.
// This is a helper method used by RenderTile and RenderPlayer.
func (renderer *EbitenRenderer) renderGlyph(screen *ebiten.Image, glyph rune, tileX, tileY int, color color.Color) {
	// Convert tile coordinates to pixel coordinates
	pixelX := float64(tileX * renderer.tileSize)
	pixelY := float64(tileY * renderer.tileSize)

	// Draw the glyph
	glyphString := string(glyph)
	options := &text.DrawOptions{}
	options.GeoM.Translate(pixelX, pixelY)
	options.ColorScale.ScaleWithColor(color)

	text.Draw(screen, glyphString, renderer.fontFace, options)
}

// RenderTile draws a single tile glyph at the specified tile coordinates.
// tileX and tileY are in tile units which are converted to pixel coordinates.
func (renderer *EbitenRenderer) RenderTile(screen *ebiten.Image, tile Tile, tileX, tileY int) {
	renderer.renderGlyph(screen, tile.Glyph, tileX, tileY, tile.Color)
}

// RenderPlayer draws the player character at their current position.
func (renderer *EbitenRenderer) RenderPlayer(screen *ebiten.Image, player Player) {
	renderer.renderGlyph(screen, player.Glyph, player.X, player.Y, player.Color)
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
