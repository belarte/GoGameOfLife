package grid

import (
	"fmt"

	"github.com/belarte/GoGameOfLife/struct/coord"
)

type Grid struct {
	buffer        [][]int
	width, height int
}

func New(width, height int) Grid {
	buffer := make([][]int, height)
	for i := range buffer {
		buffer[i] = make([]int, width)
	}
	return Grid{buffer: buffer, width: width, height: height}
}

func (g Grid) Set(c coord.Coord, value int) {
	g.buffer[c.Y][c.X] = value
}

func (g Grid) Get(c coord.Coord) int {
	return g.buffer[c.Y][c.X]
}

func (g Grid) Wrap(c coord.Coord) coord.Coord {
	return coord.New((c.X+g.width)%g.width, (c.Y+g.height)%g.height)
}

func (g Grid) String() string {
	result := ""
	for _, line := range g.buffer {
		result += fmt.Sprintf("%v\n", line)
	}
	return result
}
