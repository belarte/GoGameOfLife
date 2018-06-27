package list_test

import (
	"testing"

	"github.com/belarte/GoGameOfLife/struct/coord"
	"github.com/belarte/GoGameOfLife/struct/list"
)

func TestCoordListEquality(t *testing.T) {
	var table = []struct {
		left     list.List
		right    list.List
		expected bool
	}{
		{list.New(), list.New(), true},
		{list.New(), list.New(coord.New(1, 2)), false},
		{list.New(coord.New(1, 2)), list.New(), false},
		{list.New(coord.New(1, 2)), list.New(coord.New(1, 2)), true},
		{list.New(coord.New(1, 2)), list.New(coord.New(3, 4)), false},
		{list.New(coord.New(1, 2)), list.New(coord.New(1, 2), coord.New(3, 4)), false},
		{list.New(coord.New(1, 2), coord.New(3, 4)), list.New(coord.New(1, 2), coord.New(3, 4)), true},
	}

	for _, entry := range table {
		result := list.Compare(entry.left, entry.right)
		if result != entry.expected {
			t.Errorf("Error expected=%t but result=%t\nleft=%v, right=%v",
				entry.expected, result, entry.left, entry.right)
		}
	}
}
