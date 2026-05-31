package src

import (
	"log"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	ball *Ball
}

func (g *Game) Update() error {
	g.ball.Update()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.ball.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 600, 400
}

func NewGame() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Orbscape!")

	// initialize objects
	ballX := float32(48.0)
	ballY := float32(48.0)
	ballRadius := float32(8.0)
	game := &Game{
		ball: &Ball{ballX, ballY, ballRadius, "Dave"},
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

