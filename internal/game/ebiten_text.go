package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

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

// drawText is a helper to render text at pixel coordinates.
func (renderer *EbitenRenderer) drawText(screen *ebiten.Image, txt string, x, y float64, clr color.Color) {
	options := &text.DrawOptions{}
	options.GeoM.Translate(x, y)
	options.ColorScale.ScaleWithColor(clr)

	text.Draw(screen, txt, renderer.fontFace, options)
}

// centerText measures, centers, and draws at the given Y position.
func (renderer *EbitenRenderer) centerText(screen *ebiten.Image, txt string, y float64, clr color.Color) {
	screenWidthPixels := float64(renderer.game.Width * renderer.tileSize)
	textWidth := text.Advance(txt, renderer.fontFace)
	x := (screenWidthPixels - textWidth) / 2.0
	renderer.drawText(screen, txt, x, y, clr)
}
