package game

import (
	"github.com/ian-yeh/orbscape/shared"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func DrawBall(screen *ebiten.Image, b *shared.Ball) {
	vector.FillCircle(screen, b.X, b.Y, b.Rad, shared.RED, true)
}

func DrawGrid(screen *ebiten.Image, cellSize float32, grid [][]shared.TileType) {
	for row := range grid {
		for col := range grid[row] {
			x := float32(col) * cellSize
			y := float32(row) * cellSize

			switch grid[row][col] {
			case shared.Wall:
				vector.FillRect(screen, x, y, cellSize, cellSize, shared.WHITE, false)
			case shared.Goal:
				vector.FillRect(screen, x, y, cellSize, cellSize, shared.GREEN, false)
			}

			vector.StrokeRect(screen, x, y, cellSize, cellSize, 1, shared.BLACK, false)
		}
	}
}


