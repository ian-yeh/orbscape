package game

import (
	"github.com/ian-yeh/orbscape/shared"
	"github.com/hajimehoshi/ebiten/v2"
)

func HandleBallInput(b *shared.Ball) {
	switch {
		case ebiten.IsKeyPressed(ebiten.KeyArrowLeft): b.Move(-1, 0)
		case ebiten.IsKeyPressed(ebiten.KeyArrowRight): b.Move(1, 0)
		case ebiten.IsKeyPressed(ebiten.KeyArrowUp): b.Move(0, -1)
		case ebiten.IsKeyPressed(ebiten.KeyArrowDown): b.Move(0, 1)
	}
}

func ShouldQuit() bool {
	return ebiten.IsKeyPressed(ebiten.KeyEscape)
}
