package game

import (
	"testing"

	"github.com/hajimehoshi/ebiten/v2"
)

func TestNewEbitenRenderer(t *testing.T) {
	renderer, err := NewEbitenRenderer(80, 24, "../../assets/fonts/Go-Mono.ttf", 16)
	if err != nil {
		t.Fatalf("failed to create renderer: %v", err)
	}

	var _ ebiten.Game = renderer

	if renderer.screenWidth != 80*16 {
		t.Errorf("expected screenWidth 1280, got %d", renderer.screenWidth)
	}

	if renderer.screenHeight != 24*16 {
		t.Errorf("expected screenHeight 384, got %d", renderer.screenHeight)
	}

	if renderer.fontFace == nil {
		t.Error("expected fontFace to be set, got nil")
	}
}

func TestRenderTile(t *testing.T) {
	renderer, err := NewEbitenRenderer(80, 24, "../../assets/fonts/Go-Mono.ttf", 16)
	if err != nil {
		t.Fatalf("failed to create renderer: %v", err)
	}

	testImage := ebiten.NewImage(renderer.screenWidth, renderer.screenHeight)

	// Can't verify pixel perfect rendering in a test.
	// Verify the method exists and doesn't panic.
	renderer.RenderTile(testImage, FloorTile, 0, 0)
}
