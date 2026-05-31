package main

import (
	"log"
	"image/color"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// colour constants
var RED = color.RGBA{255, 0, 0, 255}
var GREEN = color.RGBA{0, 255, 0, 255} 
var BLUE = color.RGBA{0, 0, 255, 255} 

type Game struct {
	ball *Ball
}

type Ball struct{
	x, y, rad float32
	name string
}

func (g *Game) Update() error {
	g.ball.x += 1.0
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	vector.FillCircle(screen, g.ball.x, g.ball.y, g.ball.rad, RED, true)

	ebitenutil.DebugPrint(screen, "Hello, World!")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")

	// initialize objects
	ballX := float32(50.0)
	ballY := float32(50.0)
	ballRadius := float32(20.0)
	game := &Game{
		ball: &Ball{ballX, ballY, ballRadius, "Dave"},
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

