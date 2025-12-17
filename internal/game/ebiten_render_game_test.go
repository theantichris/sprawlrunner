package game

import (
	"testing"

	"github.com/hajimehoshi/ebiten/v2"
)

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

	t.Run("only renders viewport", func(t *testing.T) {
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
			t.Errorf("expected to render 1120 tiles, calculated bounds would render %d", expectedTileCount)
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

func TestRenderPlayer(t *testing.T) {
	t.Run("centers player in viewport", func(t *testing.T) {
		game := NewGame()

		renderer, err := NewEbitenRenderer(game, fontGoMono, 16.0)
		if err != nil {
			t.Fatalf("failed to create renderer: %v", err)
		}

		// Place player and camera at center of map
		game.Player.X = 40
		game.Player.Y = 12
		game.CameraX = 40
		game.CameraY = 12

		minX, minY, _, _ := renderer.CalculateViewportBounds()

		// Calculate where player should appear on screen
		screenX, screenY := renderer.CalculatePlayerScreenPosition()

		// Player would position (40, 12) - viewport minimum (12, 2) = screen (28, 10)
		// This is the center of a 56x20 viewport
		expectedX := game.Player.X - minX
		expectedY := game.Player.Y - minY

		if screenX != expectedX || screenY != expectedY {
			t.Errorf("expected player at screen (%d,%d), got (%d,%d)", expectedX, expectedY, screenX, screenY)
		}

		// Should be centered (28, 10)
		if screenX != 28 || screenY != 10 {
			t.Errorf("expected player at screen (28,10), got (%d,%d)", screenX, screenY)
		}
	})
}

func TestRenderStatsPanel(t *testing.T) {
	game := NewGame()

	game.Player.Name = "Decker"
	game.Player.Level = 1
	game.Player.Health = 15

	renderer, err := NewEbitenRenderer(game, fontGoMono, 16.0)
	if err != nil {
		t.Errorf("failed to create renderer: %v", err)
	}

	screen := ebiten.NewImage(renderer.screenWidth, renderer.screenHeight)

	// Should not panic
	renderer.RenderStatsPanel(screen)
}

func TestRenderMessageLog(t *testing.T) {
	game := NewGame()

	renderer, err := NewEbitenRenderer(game, fontGoMono, 16.0)
	if err != nil {
		t.Fatalf("failed to create renderer: %v", err)
	}

	screen := ebiten.NewImage(renderer.screenWidth, renderer.screenHeight)

	// Should not panic
	renderer.RenderMessageLog(screen)
}
