package game

import (
	"testing"

	"github.com/hajimehoshi/ebiten/v2"
)

func TestEbitenGameCreation(t *testing.T) {
	// Create a new Ebiten renderer
	renderer := NewEbitenRenderer(80, 24)

	// Verify it implements ebiten.Game interface
	var _ ebiten.Game = renderer

	// Verify dimensions are set correctly
	if renderer.screenWidth != 80*16 {
		t.Errorf("expected screenWidth 1280, got %d", renderer.screenWidth)
	}

	if renderer.screenHeight != 24*16 {
		t.Errorf("expected screenHeight 384, got %d", renderer.screenHeight)
	}
}
