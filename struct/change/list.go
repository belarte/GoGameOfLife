package change

import "github.com/belarte/GoGameOfLife/struct/list"

type List struct {
	Alive list.List
	Dead  list.List
}

func New(alive, dead list.List) List {
	return List{Alive: alive, Dead: dead}
}

func Compare(left, right List) bool {
	return list.Compare(left.Alive, right.Alive) && list.Compare(left.Dead, right.Dead)
}

func (l List) Cover() list.List {
	result := l.Alive.Cover()
	result.Concat(l.Dead.Cover())
	return result
}
