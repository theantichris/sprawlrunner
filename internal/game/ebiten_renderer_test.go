package game

import (
	"errors"
	"os"
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

func TestCalculateViewportBounds(t *testing.T) {
	game := NewGame()
	game.CameraX = 30
	game.CameraY = 12

	renderer, err := NewEbitenRenderer(game, fontGoMono, 16.0)
	if err != nil {
		t.Fatalf("failed to create renderer: %v", err)
	}

	minX, minY, maxX, maxY := renderer.CalculateViewportBounds()

	if minX != 2 || minY != 2 || maxX != 58 || maxY != 22 {
		t.Errorf("expected bounds 2, 2, 58, 22, got %d, %d, %d, %d", minX, minY, maxX, maxY)
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

func TestSinglePressMovement(t *testing.T) {
	game := NewGame()
	renderer, err := NewEbitenRenderer(game, fontGoMono, tileSize)
	if err != nil {
		t.Fatalf("failed to create renderer: %v", err)
	}

	initX := game.Player.X
	initY := game.Player.Y

	// Call Update() multiple times (simulating held key)
	// With proper single press handling the player should only move once
	// Note: Can't easily test keyboard state in unit tests but can verify the
	// logic structure exists

	// This test verifies that the renderer is ready for inpututil usage
	// Actual single press behavior will be verified by manual testing

	// Verify Update() doesn't panic
	err = renderer.Update()
	if err != nil {
		t.Errorf("update returned err: %v", err)
	}

	// Player shouldn't have moved (no key pressed in test)
	if renderer.game.Player.X != initX || renderer.game.Player.Y != initY {
		t.Errorf("player position changed unexpectedly: got (%d, %d), want (%d, %d)", renderer.game.Player.X, renderer.game.Player.Y, initX, initY)
	}
}

func TestFontFileNotFound(t *testing.T) {
	game := NewGame()

	_, err := NewEbitenRenderer(game, "nofont.ttf", tileSize)
	if err == nil {
		t.Fatal("expected error for nonexistent font file, got nil")
	}

	if !errors.Is(err, ErrFontNotFound) {
		t.Errorf("expected %v, got %v", ErrFontNotFound, err)
	}
}

func TestFontParseFailed(t *testing.T) {
	game := NewGame()

	tempFile := "/tmp/invalid_font.ttf"
	err := os.WriteFile(tempFile, []byte("not a valid font file"), 0644)
	if err != nil {
		t.Fatalf("failed to create test file: %v", err)
	}

	defer func() {
		_ = os.Remove(tempFile)
	}()

	_, err = NewEbitenRenderer(game, tempFile, tileSize)
	if err == nil {
		t.Fatal("expected error for invalid font file, got nil")
	}

	if !errors.Is(err, ErrFontParseFailed) {
		t.Errorf("expected %v, got %v", ErrFontParseFailed, err)
	}
}
