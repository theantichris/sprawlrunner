package game

import (
	"errors"
	"fmt"
	"image/color"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
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
		return nil, fmt.Errorf("%w", errors.Join(ErrFontNotFound, err))
	}

	defer func() {
		_ = fontData.Close()
	}()

	fontSource, err := text.NewGoTextFaceSource(fontData)
	if err != nil {
		return nil, fmt.Errorf("%w", errors.Join(ErrFontParseFailed, err))
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
	// Quit
	if ebiten.IsKeyPressed(ebiten.KeyShift) && ebiten.IsKeyPressed(ebiten.KeyQ) {
		renderer.game.RequestQuit()
	}

	// Handle quit confirmation if active
	if renderer.game.IsConfirmingQuit() {
		if inpututil.IsKeyJustPressed(ebiten.KeyY) {
			if renderer.game.ConfirmQuit(true) {
				return ebiten.Termination
			}
		}

		if inpututil.IsKeyJustPressed(ebiten.KeyN) {
			renderer.game.ConfirmQuit(false)
		}

		return nil
	}

	// Up
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) || inpututil.IsKeyJustPressed(ebiten.KeyK) || inpututil.IsKeyJustPressed(ebiten.KeyNumpad8) {
		renderer.game.MovePlayer(0, -1)
	}

	// Down
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) || inpututil.IsKeyJustPressed(ebiten.KeyJ) || inpututil.IsKeyJustPressed(ebiten.KeyNumpad2) {
		renderer.game.MovePlayer(0, 1)
	}

	// Left
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) || inpututil.IsKeyJustPressed(ebiten.KeyH) || inpututil.IsKeyJustPressed(ebiten.KeyNumpad4) {
		renderer.game.MovePlayer(-1, 0)
	}

	// Right
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) || inpututil.IsKeyJustPressed(ebiten.KeyL) || inpututil.IsKeyJustPressed(ebiten.KeyNumpad6) {
		renderer.game.MovePlayer(1, 0)
	}

	// Up left
	if inpututil.IsKeyJustPressed(ebiten.KeyHome) || inpututil.IsKeyJustPressed(ebiten.KeyY) || inpututil.IsKeyJustPressed(ebiten.KeyNumpad7) {
		renderer.game.MovePlayer(-1, -1)
	}

	// Up right
	if inpututil.IsKeyJustPressed(ebiten.KeyPageUp) || inpututil.IsKeyJustPressed(ebiten.KeyU) || inpututil.IsKeyJustPressed(ebiten.KeyNumpad9) {
		renderer.game.MovePlayer(1, -1)
	}

	// Down left
	if inpututil.IsKeyJustPressed(ebiten.KeyEnd) || inpututil.IsKeyJustPressed(ebiten.KeyB) || inpututil.IsKeyJustPressed(ebiten.KeyNumpad1) {
		renderer.game.MovePlayer(-1, 1)
	}

	// Down right
	if inpututil.IsKeyJustPressed(ebiten.KeyPageDown) || inpututil.IsKeyJustPressed(ebiten.KeyN) || inpututil.IsKeyJustPressed(ebiten.KeyNumpad3) {
		renderer.game.MovePlayer(1, 1)
	}

	return nil
}

// Draw renders the game state to the screen. Required by ebiten.Game interface.
func (renderer *EbitenRenderer) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black) // Clear screen to black
	renderer.RenderMap(screen, renderer.game)
	renderer.RenderPlayer(screen, renderer.game.Player)

	if renderer.game.IsConfirmingQuit() {
		renderer.RenderQuitPrompt(screen)
	}
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

// RenderQuitPrompt draws the quit confirmation dialog.
func (renderer *EbitenRenderer) RenderQuitPrompt(screen *ebiten.Image) {
	prompt := "Really quit? (Y/N)"

	// Center horizontally, place near top
	promptX := float64(renderer.screenWidth/2 - len(prompt)*renderer.tileSize/4)
	promptY := float64(renderer.tileSize * 2)

	options := &text.DrawOptions{}
	options.GeoM.Translate(promptX, promptY)
	options.ColorScale.ScaleWithColor(color.RGBA{R: 255, G: 255, B: 0, A: 255}) // Yellow

	text.Draw(screen, prompt, renderer.fontFace, options)
}
