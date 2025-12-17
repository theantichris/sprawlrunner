package game

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

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

// RenderTile draws a single tile glyph at the specified tile coordinates.
// tileX and tileY are in tile units which are converted to pixel coordinates.
func (renderer *EbitenRenderer) RenderTile(screen *ebiten.Image, tile Tile, tileX, tileY int) {
	renderer.renderGlyph(screen, tile.Glyph, tileX, tileY, tile.Color)
}

// RenderPlayer draws the player character at their viewport relative position.
func (renderer *EbitenRenderer) RenderPlayer(screen *ebiten.Image, player Player) {
	screenX, screenY := renderer.CalculatePlayerScreenPosition()
	renderer.renderGlyph(screen, player.Glyph, screenX, screenY, player.Color)
}

// RenderStatsPanel draws the player stats in the right panel (24 columns).
func (renderer *EbitenRenderer) RenderStatsPanel(screen *ebiten.Image) {
	// Panel starts at x=56 (after viewport), top of screen
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

// RenderMessageLog draws the message log area at the bottom of the screen
// (4 lines high).
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
