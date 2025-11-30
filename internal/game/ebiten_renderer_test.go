package game

import (
	"testing"

	"github.com/hajimehoshi/ebiten/v2"
)

func TestNewEbitenRenderer(t *testing.T) {
	game := NewGame()

	renderer, err := NewEbitenRenderer(game, fontGoMono, tileSize)
	if err != nil {
		t.Fatalf("failed to create renderer: %v", err)
	}

	var _ ebiten.Game = renderer

	if renderer.screenWidth != mapWidth*tileSize {
		t.Errorf("expected screenWidth 1280, got %d", renderer.screenWidth)
	}

	if renderer.screenHeight != mapHeight*tileSize {
		t.Errorf("expected screenHeight 384, got %d", renderer.screenHeight)
	}

	if renderer.fontFace == nil {
		t.Error("expected fontFace to be set, got nil")
	}
}

func TestRenderTile(t *testing.T) {
	game := NewGame()

	renderer, err := NewEbitenRenderer(game, fontGoMono, tileSize)
	if err != nil {
		t.Fatalf("failed to create renderer: %v", err)
	}

	testImage := ebiten.NewImage(renderer.screenWidth, renderer.screenHeight)

	// Verify method exists and does not panic.
	renderer.RenderTile(testImage, FloorTile, 0, 0)
}

func TestRenderMap(t *testing.T) {
	game := NewGame()

	renderer, err := NewEbitenRenderer(game, fontGoMono, tileSize)
	if err != nil {
		t.Fatalf("failed to create renderer: %v", err)
	}

	testImage := ebiten.NewImage(renderer.screenWidth, renderer.screenHeight)

	// Verify method exists and does not panic.
	renderer.RenderMap(testImage, game)
}

func TestRenderPlayer(t *testing.T) {
	game := NewGame()

	renderer, err := NewEbitenRenderer(game, fontGoMono, tileSize)
	if err != nil {
		t.Fatalf("failed to create renderer: %v", err)
	}

	testImage := ebiten.NewImage(renderer.screenWidth, renderer.screenHeight)

	// Verify method exists and does not panic.
	renderer.RenderPlayer(testImage, game.Player)
}

func TestHandleInput(t *testing.T) {
	game := NewGame()

	renderer, err := NewEbitenRenderer(game, fontGoMono, tileSize)
	if err != nil {
		t.Fatalf("failed to create renderer: %v", err)
	}

	renderer.game = game

	initX := game.Player.X
	initY := game.Player.Y

	// Simulate pressing the right arrow key
	// Note: Can't easily simulate keypresses in unit tests, so this test verifies
	// the game reference is stored correctly

	if renderer.game == nil {
		t.Error("expected game to be set, got nil")
	}

	if renderer.game.Player.X != initX || renderer.game.Player.Y != initY {
		t.Errorf("player position changed unexpectedly: got (%d, %d), want (%d, %d)", renderer.game.Player.X, renderer.game.Player.Y, initX, initY)
	}
}
