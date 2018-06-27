package engine

type Coord struct {
	X, Y int
}

func New(x, y int) Coord {
	return Coord{X: x, Y: y}
}

type List []Coord

func NewList(coords ...Coord) List {
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
