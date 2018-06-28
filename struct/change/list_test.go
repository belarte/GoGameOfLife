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
