package src

import (
	"github.com/ian-yeh/orbscape/shared"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	*shared.GameState
}

func (g *Game) Update() error {
	HandleBallInput(g.GameState.Ball)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	DrawBall(screen, g.GameState.Ball)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 600, 400
}

func NewGame() error {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Orbscape!")

	// initialize objects
	game := &Game{
		GameState: &shared.GameState{
			Ball: &shared.Ball{X: 48, Y: 48, Rad: 8, Name: "Dave"},
		},
	}

	return ebiten.RunGame(game)
}

