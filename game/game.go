package game

import (
	"github.com/ian-yeh/orbscape/shared"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	*shared.GameState
}
func (g *Game) Update() error {
	if ShouldQuit() {
		return ebiten.Termination
	}

	HandleBallInput(g.GameState.Ball)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	DrawGrid(screen, 50)
	DrawBall(screen, g.GameState.Ball)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1600, 900
}

func NewGame() error {
	ebiten.SetWindowSize(1600, 900)
	ebiten.SetWindowTitle("Orbscape!")

	game := &Game{
		GameState: &shared.GameState{
			Ball: &shared.Ball{X: 48, Y: 48, Rad: 44, Name: "Dave"},
		},
	}

	return ebiten.RunGame(game)
}

