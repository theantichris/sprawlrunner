package game

import "testing"

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
