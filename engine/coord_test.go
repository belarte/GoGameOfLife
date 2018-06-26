package engine_test

import (
	"testing"

	"github.com/belarte/GoGameOfLife/engine"
)

func TestCoordEquality(t *testing.T) {
	var table = []struct {
		left     engine.Coord
		right    engine.Coord
		expected bool
	}{
		{engine.New(0, 0), engine.New(0, 0), true},
		{engine.New(1, 2), engine.New(1, 2), true},
		{engine.New(65536, 16), engine.New(65536, 16), true},
		{engine.New(0, 4096), engine.New(0, 4096), true},
		{engine.New(37, 1234567), engine.New(37, 1234567), true},
		{engine.New(0, 0), engine.New(0, 1), false},
		{engine.New(1, 0), engine.New(0, 0), false},
		{engine.New(16, 32), engine.New(16, 33), false},
		{engine.New(64, 128), engine.New(63, 128), false},
		{engine.New(1, 2), engine.New(3, 4), false},
	}

	for _, entry := range table {
		result := entry.left == entry.right
		if result != entry.expected {
			t.Errorf("Error expected=%t but result=%t", entry.expected, result)
		}
	}
}
