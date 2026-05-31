package src

import (
	"github.com/ian-yeh/orbscape/shared"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/hajimehoshi/ebiten/v2"
)

type Ball struct {
	x, y, rad float32
	name string
}

func (b *Ball) Update() {
	if (ebiten.IsKeyPressed(ebiten.KeyArrowLeft)) {
		b.x -= 1.5
	} else if (ebiten.IsKeyPressed(ebiten.KeyArrowRight)) {
		b.x += 1.5
	} else if (ebiten.IsKeyPressed(ebiten.KeyArrowUp)) {
		b.y -= 1.5
	} else if (ebiten.IsKeyPressed(ebiten.KeyArrowDown)) {
		b.y += 1.5
	}
}

func (b *Ball) Draw(screen *ebiten.Image) {
	vector.FillCircle(screen, b.x, b.y, b.rad, shared.RED, true)
}

