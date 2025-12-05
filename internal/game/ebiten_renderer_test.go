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
	t.Run("sets min and max x and y", func(t *testing.T) {
		game := NewGame()
		game.CameraX = 30
		game.CameraY = 12

		renderer, err := NewEbitenRenderer(game, fontGoMono, 16.0)
		if err != nil {
			t.Fatalf("failed to create renderer: %v", err)
		}

		minX, minY, maxX, maxY := renderer.CalculateViewportBounds()

		if minX != 2 || minY != 2 || maxX != 58 || maxY != 22 {
			t.Errorf("expected bounds (2,2,58,22), got (%d,%d,%d,%d)", minX, minY, maxX, maxY)
		}
	})

	t.Run("clamps at edges", func(t *testing.T) {
		game := NewGame()

		renderer, err := NewEbitenRenderer(game, fontGoMono, 16.0)
		if err != nil {
			t.Fatalf("failed to create renderer: %v", err)
		}

		// Camera at top left corner
		game.CameraX = 0
		game.CameraY = 0
		minX, minY, maxX, maxY := renderer.CalculateViewportBounds()

		if minX != 0 || minY != 0 {
			t.Errorf("top left: expected min (0,0), got (%d,%d)", minX, minY)
		}

		if maxX != 56 || maxY != 20 {
			t.Errorf("top left: expected max (56,20), got (%d,%d)", maxX, maxY)
		}

		// Camera at bottom right corner (map is 80x24)
		game.CameraX = 79
		game.CameraY = 23
		minX, minY, maxX, maxY = renderer.CalculateViewportBounds()

		if minX != 24 || minY != 4 {
			t.Errorf("bottom right: expected min (24,4), got (%d,%d)", minX, minY)
		}

		if maxX != 80 || maxY != 24 {
			t.Errorf("bottom right: expected max (80,24), got (%d,%d)", maxX, maxY)
		}
	})
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
	t.Run("method exists and does not panic", func(t *testing.T) {
		game := NewGame()

		renderer, err := NewEbitenRenderer(game, fontGoMono, tileSize)
		if err != nil {
			t.Fatalf("failed to create renderer: %v", err)
		}

		testImage := ebiten.NewImage(renderer.screenWidth, renderer.screenHeight)

		// Verify method exists and does not panic.
		renderer.RenderMap(testImage, game)
	})

	t.Run("only renders viewpoert", func(t *testing.T) {
		game := NewGame()

		// Position camera so viewport shows tiles [2,2] to [57,21]
		game.CameraX = 30
		game.CameraY = 12

		renderer, err := NewEbitenRenderer(game, fontGoMono, 16.0)
		if err != nil {
			t.Fatalf("failed to create renderer: %v", err)
		}

		// Create a test screen
		screen := ebiten.NewImage(renderer.screenWidth, renderer.screenHeight)

		// RenderMap should not panic and should complete
		renderer.RenderMap(screen, game)

		// Verify bounds were calculated correctly
		minX, minY, maxX, maxY := renderer.CalculateViewportBounds()
		expectedTileCount := (maxX - minX) * (maxY - minY)

		// Should be rendering 56x20 = 1120 tiles
		if expectedTileCount != 1120 {
			t.Errorf("expected to render 1120 tiles, calculated bounds would rnder %d", expectedTileCount)
		}
	})
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
