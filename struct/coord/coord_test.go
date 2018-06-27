package coord_test

import (
	"testing"

	"github.com/belarte/GoGameOfLife/struct/coord"
)

func TestCoordEquality(t *testing.T) {
	var table = []struct {
		left     coord.Coord
		right    coord.Coord
		expected bool
	}{
		{coord.New(0, 0), coord.New(0, 0), true},
		{coord.New(1, 2), coord.New(1, 2), true},
		{coord.New(65536, 16), coord.New(65536, 16), true},
		{coord.New(0, 4096), coord.New(0, 4096), true},
		{coord.New(37, 1234567), coord.New(37, 1234567), true},
		{coord.New(0, 0), coord.New(0, 1), false},
		{coord.New(1, 0), coord.New(0, 0), false},
		{coord.New(16, 32), coord.New(16, 33), false},
		{coord.New(64, 128), coord.New(63, 128), false},
		{coord.New(1, 2), coord.New(3, 4), false},
	}

	for _, entry := range table {
		result := entry.left == entry.right
		if result != entry.expected {
			t.Errorf("Error expected=%t but result=%t\nleft=%v, right=%v",
				entry.expected, result, entry.left, entry.right)
		}
	}
}
