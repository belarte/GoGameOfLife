package list

import "github.com/belarte/GoGameOfLife/struct/coord"

type List []coord.Coord

func New(coords ...coord.Coord) List {
	list := List{}
	list = append(list, coords...)
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
