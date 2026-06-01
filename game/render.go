package game

import (
	"github.com/ian-yeh/orbscape/shared"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func DrawBall(screen *ebiten.Image, b *shared.Ball) {
	vector.FillCircle(screen, b.X, b.Y, b.Rad, shared.RED, true)
}

func DrawGrid(screen *ebiten.Image, cellSize float32) {
	w, h := float32(screen.Bounds().Dx()), float32(screen.Bounds().Dy())
	for x := float32(0); x <= w; x += cellSize {
		vector.StrokeLine(screen, x, 0, x, h, 1, shared.GREEN, false)
	}
	for y := float32(0); y <= h; y += cellSize {
		vector.StrokeLine(screen, 0, y, w, y, 1, shared.GREEN, false)
	}
}


