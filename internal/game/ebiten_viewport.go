package game

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

// CalculatePlayerScreenPosition returns the player's screen coordinates
// relative to the viewport origin.
func (renderer *EbitenRenderer) CalculatePlayerScreenPosition() (int, int) {
	minX, minY, _, _ := renderer.CalculateViewportBounds()

	screenX := renderer.game.Player.X - minX
	screenY := renderer.game.Player.Y - minY

	return screenX, screenY
}
