package life_test

import (
	"testing"

	"github.com/belarte/GoGameOfLife/life"
	"github.com/belarte/GoGameOfLife/struct/change"
	"github.com/belarte/GoGameOfLife/struct/coord"
	"github.com/belarte/GoGameOfLife/struct/list"
)

func TestNextChangeReturnsAddedAfterZeroStep(t *testing.T) {
	game := life.New(32, 32)

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

func TestStepProducesCorrectChangeList(t *testing.T) {
	game := life.New(5, 5)
	game.Add(change.New(list.New(coord.New(2, 1), coord.New(3, 2), coord.New(1, 3), coord.New(2, 3), coord.New(3, 3)), list.New()))

	var table = []struct {
		expected change.List
	}{
		{change.New(
			list.New(coord.New(1, 2), coord.New(2, 4)),
			list.New(coord.New(2, 1), coord.New(1, 3)))},
		{change.New(
			list.New(coord.New(1, 3), coord.New(3, 4)),
			list.New(coord.New(1, 2), coord.New(2, 3)))},
		{change.New(
			list.New(coord.New(2, 2), coord.New(4, 3)),
			list.New(coord.New(3, 2), coord.New(1, 3)))},
		{change.New(
			list.New(coord.New(3, 2), coord.New(4, 4)),
			list.New(coord.New(2, 2), coord.New(3, 3)))},
	}

	for i, entry := range table {
		game.Step()
		result := game.NextChangeList()

		if !change.Compare(result, entry.expected) {
			t.Errorf("[%d] Error expected=%v but result=%v\n%v", i, entry.expected, result, game)
		}
	}
}
