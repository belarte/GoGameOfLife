package engine

type Coord struct {
	X, Y int
}

func New(x, y int) Coord {
	return Coord{X: x, Y: y}
}
