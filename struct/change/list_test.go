package change_test

import (
	"testing"

	"github.com/belarte/GoGameOfLife/struct/change"
	"github.com/belarte/GoGameOfLife/struct/coord"
	"github.com/belarte/GoGameOfLife/struct/list"
)

func TestChangeListCompare(t *testing.T) {
	var table = []struct {
		left     change.List
		right    change.List
		expected bool
	}{
		{
			change.New(list.New(), list.New()),
			change.New(list.New(), list.New()),
			true,
		},
		{
			change.New(list.New(coord.New(1, 2)), list.New(coord.New(3, 4))),
			change.New(list.New(coord.New(5, 6)), list.New(coord.New(7, 8))),
			false,
		},
		{
			change.New(list.New(coord.New(1, 2)), list.New(coord.New(1, 2), coord.New(3, 4))),
			change.New(list.New(coord.New(3, 4)), list.New(coord.New(1, 2), coord.New(3, 4))),
			false,
		},
		{
			change.New(list.New(coord.New(1, 2), coord.New(3, 4)), list.New(coord.New(1, 2))),
			change.New(list.New(coord.New(1, 2), coord.New(3, 4)), list.New(coord.New(3, 4))),
			false,
		},
		{
			change.New(list.New(coord.New(1, 2), coord.New(3, 4)), list.New()),
			change.New(list.New(coord.New(1, 2), coord.New(3, 4)), list.New()),
			true,
		},
	}

	for _, entry := range table {
		result := change.Compare(entry.left, entry.right)
		if result != entry.expected {
			t.Errorf("Error expected=%t but result=%t\nleft=%v, right=%v",
				entry.expected, result, entry.left, entry.right)
		}
	}
}

func TestCover(t *testing.T) {
	var table = []struct {
		input  change.List
		output list.List
	}{
		{
			change.New(list.New(), list.New()),
			list.New(),
		},
		{
			change.New(list.New(coord.New(2, 2)), list.New()),
			list.New(coord.New(1, 1), coord.New(1, 2), coord.New(1, 3), coord.New(2, 1), coord.New(2, 2),
				coord.New(2, 3), coord.New(3, 1), coord.New(3, 2), coord.New(3, 3)),
		},
		{
			change.New(list.New(coord.New(2, 2)), list.New(coord.New(5, 5))),
			list.New(coord.New(1, 1), coord.New(1, 2), coord.New(1, 3), coord.New(2, 1), coord.New(2, 2),
				coord.New(2, 3), coord.New(3, 1), coord.New(3, 2), coord.New(3, 3),
				coord.New(4, 4), coord.New(4, 5), coord.New(4, 6), coord.New(5, 4), coord.New(5, 5),
				coord.New(5, 6), coord.New(6, 4), coord.New(6, 5), coord.New(6, 6)),
		},
		{
			change.New(list.New(coord.New(2, 2), coord.New(3, 2)), list.New()),
			list.New(coord.New(1, 1), coord.New(1, 2), coord.New(1, 3), coord.New(2, 1),
				coord.New(2, 2), coord.New(2, 3), coord.New(3, 1), coord.New(3, 2),
				coord.New(3, 3), coord.New(4, 1), coord.New(4, 2), coord.New(4, 3)),
		},
		{
			change.New(list.New(), list.New(coord.New(2, 2), coord.New(4, 2))),
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
