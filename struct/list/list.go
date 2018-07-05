package list

import (
	"fmt"
	"github.com/belarte/GoGameOfLife/struct/coord"
)

type List map[coord.Coord]bool

func New(coords ...coord.Coord) List {
	var list List = make(map[coord.Coord]bool)
	list.Append(coords...)
	return list
}

func Compare(left, right List) bool {
	if len(left) != len(right) {
		return false
	}

	for i := range left {
		if left[i] != right[i] {
			return false
		}
	}

	return true
}

func (list List) Append(coords ...coord.Coord) {
	for _, c := range coords {
		list[c] = true
	}
}

func (list List) Concat(other List) {
	for key, value := range other {
		list[key] = value
	}
}

func (list List) Cover() List {
	result := New()

	for c, value := range list {
		if value {
			result.Append(c)
			result.Concat(neighbours8(c))
		}
	}

	return result
}

func neighbours8(c coord.Coord) List {
	return New(
		coord.New(c.X-1, c.Y-1),
		coord.New(c.X-1, c.Y),
		coord.New(c.X-1, c.Y+1),
		coord.New(c.X, c.Y-1),
		coord.New(c.X, c.Y+1),
		coord.New(c.X+1, c.Y-1),
		coord.New(c.X+1, c.Y),
		coord.New(c.X+1, c.Y+1),
	)
}

func (list List) String() string {
	result := "["

	for key, value := range list {
		if value {
			result += fmt.Sprintf("%v, ", key)
		}
	}

	if len(result) > 1 {
		result = result[:len(result)-2]
	}

	return result + "]"
}
