package game

import "github.com/hajimehoshi/ebiten/v2"

var titleScreenArt = []string{
	"  _________                          .__                                          ",
	" /   _____/_________________ __  _  _|  |_______ __ __  ____   ____   ___________ ",
	" \\_____  \\\\____ \\_  __ \\__  \\\\ \\/ \\/ /  |\\_  __ \\  |  \\/    \\ /    \\_/ __ \\_  __ \\",
	" /        \\  |_> >  | \\// __ \\\\     /|  |_|  | \\/  |  /   |  \\   |  \\  ___/|  | \\/",
	"/_______  /   __/|__|  (____  /\\/\\_/ |____/__|  |____/|___|  /___|  /\\___  >__|   ",
	"        \\/|__|              \\/                             \\/     \\/     \\/       ",
}

const (
	titleScreenSubtitle    = "A Cyberpunk Roguelike"
	titleScreenCopyright   = "Copyright 2025"
	titleScreenInstruction = "Press SPACE to start or Q to quit"
)

// RenderTitleScreen draws the title screen with ASCII art and instructions.
func (renderer *EbitenRenderer) RenderTitleScreen(screen *ebiten.Image) {
	screen.Fill(colorBlack) // Clear screen to black

	lineHeight := float64(renderer.tileSize)
	startY := 5.0 * lineHeight

	// Draw title ASCII art
	for i, line := range titleScreenArt {
		y := startY + float64(i)*lineHeight
		renderer.centerText(screen, line, y, colorYellow)
	}

	metaY := startY + float64(len(titleScreenArt)+3)*lineHeight
	renderer.centerText(screen, titleScreenSubtitle, metaY, colorWhite)
	renderer.centerText(screen, titleScreenCopyright, metaY+lineHeight*2, colorWhite)

	instructionsY := float64(renderer.game.Height-3) * lineHeight
	renderer.centerText(screen, titleScreenInstruction, instructionsY, colorYellow)
}
