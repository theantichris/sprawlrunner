package game

import (
	"testing"

	"github.com/hajimehoshi/ebiten/v2"
)

func TestEbitenGameCreation(t *testing.T) {
	renderer := NewEbitenRenderer(80, 24)

	var _ ebiten.Game = renderer

	if renderer.screenWidth != 80*16 {
		t.Errorf("expected screenWidth 1280, got %d", renderer.screenWidth)
	}

	if renderer.screenHeight != 24*16 {
		t.Errorf("expected screenHeight 384, got %d", renderer.screenHeight)
	}
}

func TestFontLoading(t *testing.T) {
	renderer := NewEbitenRenderer(80, 24)

	err := renderer.LoadFont("../../assets/fonts/Go-Mono.ttf", 16)
	if err != nil {
		t.Fatalf("failed to load font: %v", err)
	}

	if renderer.fontFace == nil {
		t.Errorf("expected fontFace to be set, got nil")
	}

	if renderer.fontSize != 16 {
		t.Errorf("expected fontSize 16, got %f", renderer.fontSize)
	}
}

func TestRenderTileGlyph(t *testing.T) {
	renderer := NewEbitenRenderer(80, 24)

	err := renderer.LoadFont("../../assets/fonts/Go-Mono.ttf", 16)
	if err != nil {
		t.Fatalf("failed to load font: %v", err)
	}

	testImage := ebiten.NewImage(renderer.screenWidth, renderer.screenHeight)

	// Can't verify pixel perfect rendering in a test.
	// Verify the method exists and doesn't panic.
	renderer.RenderTile(testImage, FloorTile, 0, 0)
}
