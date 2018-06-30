package life_test

import (
	"testing"

	"github.com/belarte/GoGameOfLife/life"
	"github.com/belarte/GoGameOfLife/struct/change"
	"github.com/belarte/GoGameOfLife/struct/coord"
	"github.com/belarte/GoGameOfLife/struct/list"
)

func TestNextChangeReturnsAddedAfterZeroStep(t *testing.T) {
	game := life.New()

	var table = []struct {
		in  change.List
		out change.List
	}{
		{change.New(nil, nil), change.New(nil, nil)},
		{change.New(nil, nil), change.New(list.New(), list.New())},
		{change.New(list.New(), list.New()), change.New(list.New(), list.New())},
		{change.New(list.New(), list.New()), change.New(nil, nil)},
		{
			change.New(list.New(coord.New(1, 2)), nil),
			change.New(list.New(coord.New(1, 2)), nil),
		},
		{
			change.New(list.New(coord.New(3, 4)), nil),
			change.New(list.New(coord.New(1, 2), coord.New(3, 4)), nil),
		},
		{
			change.New(nil, nil),
			change.New(list.New(coord.New(1, 2), coord.New(3, 4)), nil),
		},
		{
			change.New(list.New(coord.New(5, 6), coord.New(7, 8)), list.New(coord.New(13, 16))),
			change.New(list.New(coord.New(1, 2), coord.New(3, 4), coord.New(5, 6), coord.New(7, 8)), list.New(coord.New(13, 16))),
		},
	}

	for _, entry := range table {
		game.Add(entry.in)
		result := game.NextChangeList()
		if !change.Compare(result, entry.out) {
			t.Errorf("Error expected=%v but result=%v", entry.out, result)
		}
	}
}

func TestChangeListIsResetAfterAStep(t *testing.T) {
	game := life.New()
	game.Add(change.New(list.New(coord.New(1, 2)), list.New(coord.New(3, 4))))

	changes := game.NextChangeList()
	if change.Compare(changes, change.New(nil, nil)) {
		t.Error("Change list should not be empty")
	}

	game.Step()

	changes = game.NextChangeList()
	if !change.Compare(changes, change.New(nil, nil)) {
		t.Error("Change list should be empty")
	}
}
