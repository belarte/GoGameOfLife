package grid_test

import (
	"testing"

	"github.com/belarte/GoGameOfLife/struct/coord"
	"github.com/belarte/GoGameOfLife/struct/grid"
)

func TestGetAndSet(t *testing.T) {
	grid := grid.New(5, 2)

	grid.Set(coord.New(2, 0), 1)
	grid.Set(coord.New(2, 1), 4)
	grid.Set(coord.New(4, 1), 65536)

	var table = []struct {
		input  coord.Coord
		output int
	}{
		{coord.New(0, 0), 0},
		{coord.New(1, 0), 0},
		{coord.New(2, 0), 1},
		{coord.New(3, 0), 0},
		{coord.New(4, 0), 0},
		{coord.New(0, 1), 0},
		{coord.New(1, 1), 0},
		{coord.New(2, 1), 4},
		{coord.New(3, 1), 0},
		{coord.New(4, 1), 65536},
	}

	for i, entry := range table {
		result := grid.Get(entry.input)
		if result != entry.output {
			t.Errorf("[%d] expected=%d but result=%d", i, entry.output, result)
		}
	}
}
