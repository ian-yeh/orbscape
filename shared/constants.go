package shared

import (
	"image/color"
)

var RED = color.RGBA{255, 0, 0, 255}
var GREEN = color.RGBA{0, 255, 0, 255} 
var BLUE = color.RGBA{0, 0, 255, 255} 
var BLACK = color.RGBA{0, 0, 0, 0}
var WHITE = color.RGBA{255, 255, 255, 255}

type TileType int

const (
	Empty TileType = 0
	Wall TileType = 1
	Goal TileType = 2
)

const (
	CellSize float32 = 50
	GridCols int = 32
	GridRows int = 18
)

