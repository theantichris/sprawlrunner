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
	tileSize          = 16
	mapViewportWidth  = 56
	mapViewportHeight = 20
	hudPanelWidth     = 24
	messageLogHeight  = 4
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

// CalculateViewportBounds returns the tile coordinates visible in the viewport.
func (renderer *EbitenRenderer) CalculateViewportBounds() (int, int, int, int) {
	// Calculate viewport bounds centered on camera
	minX := renderer.game.CameraX - mapViewportWidth/2
	minY := renderer.game.CameraY - mapViewportHeight/2
	maxX := minX + mapViewportWidth
	maxY := minY + mapViewportHeight

	// Clamp to map bounds
	if minX < 0 {
		minX = 0
		maxX = mapViewportWidth
	}

	if minY < 0 {
		minY = 0
		maxY = mapViewportHeight
	}

	if maxX > renderer.game.Width {
		maxX = renderer.game.Width
		minX = maxX - mapViewportWidth
	}

	if maxY > renderer.game.Height {
		maxY = renderer.game.Height
		minY = maxY - mapViewportHeight
	}

	return minX, minY, maxX, maxY
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
	renderer.RenderStatsPanel(screen)
	renderer.RenderMessageLog(screen)
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

// CalculatePlayerScreenPosition returns the player's screen coordinates
// relative to the viewport origin.
func (renderer *EbitenRenderer) CalculatePlayerScreenPosition() (int, int) {
	minX, minY, _, _ := renderer.CalculateViewportBounds()

	screenX := renderer.game.Player.X - minX
	screenY := renderer.game.Player.Y - minY

	return screenX, screenY
}

// RenderPlayer draws the player character at their viewport relative position.
func (renderer *EbitenRenderer) RenderPlayer(screen *ebiten.Image, player Player) {
	screenX, screenY := renderer.CalculatePlayerScreenPosition()
	renderer.renderGlyph(screen, player.Glyph, screenX, screenY, player.Color)
}

// RenderMap draws all the tiles from the game map that are visible in the viewport.
func (renderer *EbitenRenderer) RenderMap(screen *ebiten.Image, game *Game) {
	minX, minY, maxX, maxY := renderer.CalculateViewportBounds()

	for y := minY; y < maxY; y++ {
		for x := minX; x < maxX; x++ {
			tile := game.Tiles[y][x]

			// Render at screen position offset by viewport origin
			screenX := x - minX
			screenY := y - minY

			renderer.RenderTile(screen, tile, screenX, screenY)
		}
	}
}

// RenderStatsPanel draws the player stats in the right panel (24 columns).
func (renderer *EbitenRenderer) RenderStatsPanel(screen *ebiten.Image) {
	// Panel stats at x=56 (after viewport), top of screen
	panelX := float64(mapViewportWidth * renderer.tileSize)
	startY := 0.0
	lineHeight := float64(renderer.tileSize)

	// Draw panel title
	renderer.drawText(screen, "== Runner ==", panelX, startY, colorYellow)

	// Draw player name
	nameY := startY + lineHeight*2
	renderer.drawText(screen, renderer.game.Player.Name, panelX, nameY, colorWhite)

	// Draw level
	levelY := nameY + lineHeight*2
	levelText := fmt.Sprintf("Level: %d", renderer.game.Player.Level)
	renderer.drawText(screen, levelText, panelX, levelY, colorWhite)

	// Draw health
	healthY := levelY + lineHeight
	healthText := fmt.Sprintf("Health: %d", renderer.game.Player.Health)
	renderer.drawText(screen, healthText, panelX, healthY, color.White)
}

// RenderMessageLog draws the message log area at the bottom of (4 lines).
func (renderer *EbitenRenderer) RenderMessageLog(screen *ebiten.Image) {
	// Message log starts below the viewport (20 tiles down)
	logY := mapViewportHeight

	// Draw separator line
	for x := 0; x < renderer.game.Width; x++ {
		renderer.renderGlyph(screen, '=', x, logY, colorYellow)
	}

	// If quit confirmation is active show it in the message log
	if renderer.game.IsConfirmingQuit() {
		promptX := 1.0 * float64(renderer.tileSize)
		promptY := float64((logY + 1) * renderer.tileSize)
		renderer.drawText(screen, "Really quit? (Y/N)", promptX, promptY, colorYellow)
	}
}

// drawText is a helper to render text at pixel coordinates.
func (renderer *EbitenRenderer) drawText(screen *ebiten.Image, txt string, x, y float64, clr color.Color) {
	options := &text.DrawOptions{}
	options.GeoM.Translate(x, y)
	options.ColorScale.ScaleWithColor(clr)

	text.Draw(screen, txt, renderer.fontFace, options)
}
