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

func TestWrap(t *testing.T) {
	grid := grid.New(4, 5)

	var table = []struct {
		input  coord.Coord
		output coord.Coord
	}{
		{coord.New(0, 0), coord.New(0, 0)},
		{coord.New(3, 4), coord.New(3, 4)},
		{coord.New(2, 3), coord.New(2, 3)},
		{coord.New(-1, 2), coord.New(3, 2)},
		{coord.New(1, -1), coord.New(1, 4)},
		{coord.New(4, 3), coord.New(0, 3)},
		{coord.New(3, 5), coord.New(3, 0)},
		{coord.New(4, 5), coord.New(0, 0)},
		{coord.New(5, 4), coord.New(1, 4)},
		{coord.New(-2, 0), coord.New(2, 0)},
	}

	for i, entry := range table {
		result := grid.Wrap(entry.input)
		if result != entry.output {
			t.Errorf("[%d] Error expected=%v but result=%v", i, entry.output, result)
		}
	}
}
