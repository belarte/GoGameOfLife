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

func (list List) String() string {
	result := "["
	for key, value := range list {
		if value {
			result += fmt.Sprintf("%v, ", key)
		}
	}
	return result[:len(result)-2] + "]"
}
