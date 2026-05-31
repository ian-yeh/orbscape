package shared

type Ball struct {
	X, Y, Rad float32
	Name string
}

const ballSpeed = 1.5

func (b *Ball) Move(dx, dy float32) {
	b.X += dx * ballSpeed
	b.Y += dy * ballSpeed
}
