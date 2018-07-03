package coord

import "fmt"

type Coord struct {
	X, Y int
}

func New(x, y int) Coord {
	return Coord{X: x, Y: y}
}

func (c Coord) String() string {
	return fmt.Sprintf("(%d, %d)", c.X, c.Y)
}
