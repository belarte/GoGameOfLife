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
		{list.New(coord.New(1, 2)), list.New(coord.New(1, 2), coord.New(1, 2)), true},
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

func TestListConcatenation(t *testing.T) {
	var table = []struct {
		left     list.List
		right    list.List
		expected list.List
	}{
		{
			list.New(),
			list.New(),
			list.New(),
		},
		{
			list.New(),
			list.New(coord.New(1, 2)),
			list.New(coord.New(1, 2)),
		},
		{
			list.New(coord.New(1, 2)),
			list.New(),
			list.New(coord.New(1, 2)),
		},
		{
			list.New(coord.New(1, 2)),
			list.New(coord.New(1, 2)),
			list.New(coord.New(1, 2)),
		},
		{
			list.New(coord.New(1, 2)),
			list.New(coord.New(3, 4)),
			list.New(coord.New(1, 2), coord.New(3, 4)),
		},
		{
			list.New(coord.New(1, 2)),
			list.New(coord.New(1, 2), coord.New(3, 4)),
			list.New(coord.New(1, 2), coord.New(3, 4)),
		},
		{
			list.New(coord.New(1, 2), coord.New(3, 4)),
			list.New(coord.New(5, 6), coord.New(3, 4)),
			list.New(coord.New(1, 2), coord.New(3, 4), coord.New(5, 6)),
		},
	}

	for _, entry := range table {
		entry.left.Concat(entry.right)
		if !list.Compare(entry.left, entry.expected) {
			t.Errorf("Error expected=%v but result=%v",
				entry.expected, entry.left)
		}
	}
}

func TestCover(t *testing.T) {
	var table = []struct {
		input  list.List
		output list.List
	}{
		{
			list.New(),
			list.New(),
		},
		{
			list.New(coord.New(2, 2)),
			list.New(coord.New(1, 1), coord.New(1, 2), coord.New(1, 3), coord.New(2, 1), coord.New(2, 2),
				coord.New(2, 3), coord.New(3, 1), coord.New(3, 2), coord.New(3, 3)),
		},
		{
			list.New(coord.New(2, 2), coord.New(5, 5)),
			list.New(coord.New(1, 1), coord.New(1, 2), coord.New(1, 3), coord.New(2, 1), coord.New(2, 2),
				coord.New(2, 3), coord.New(3, 1), coord.New(3, 2), coord.New(3, 3),
				coord.New(4, 4), coord.New(4, 5), coord.New(4, 6), coord.New(5, 4), coord.New(5, 5),
				coord.New(5, 6), coord.New(6, 4), coord.New(6, 5), coord.New(6, 6)),
		},
		{
			list.New(coord.New(2, 2), coord.New(3, 2)),
			list.New(coord.New(1, 1), coord.New(1, 2), coord.New(1, 3), coord.New(2, 1),
				coord.New(2, 2), coord.New(2, 3), coord.New(3, 1), coord.New(3, 2),
				coord.New(3, 3), coord.New(4, 1), coord.New(4, 2), coord.New(4, 3)),
		},
		{
			list.New(coord.New(2, 2), coord.New(4, 2)),
			list.New(coord.New(1, 1), coord.New(1, 2), coord.New(1, 3), coord.New(2, 1),
				coord.New(2, 2), coord.New(2, 3), coord.New(3, 1), coord.New(3, 2),
				coord.New(3, 3), coord.New(4, 1), coord.New(4, 2), coord.New(4, 3),
				coord.New(5, 1), coord.New(5, 2), coord.New(5, 3)),
		},
	}

	for i, entry := range table {
		result := entry.input.Cover()
		if !list.Compare(result, entry.output) {
			t.Errorf("[%d] Error expected=%v\nbut result=%v", i, entry.output, result)
		}
	}
}
