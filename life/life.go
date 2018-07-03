package life

import (
	"github.com/belarte/GoGameOfLife/struct/change"
	"github.com/belarte/GoGameOfLife/struct/coord"
	"github.com/belarte/GoGameOfLife/struct/list"
)

type Game interface {
	Add(changes change.List)
	Simulate(iterations int)
	Step()
	NextChangeList() change.List
}

type game struct {
	changes change.List
}

func New() Game {
	list := change.New(list.New(), list.New())
	return &game{changes: list}
}

func (g *game) Add(changes change.List) {
	g.changes.Alive.Concat(changes.Alive)
	g.changes.Dead.Concat(changes.Dead)
}

func (g *game) Simulate(iterations int) {
}

func (g *game) Step() {
	g.changes.Alive = make(map[coord.Coord]bool)
	g.changes.Dead = make(map[coord.Coord]bool)
}

func (g *game) NextChangeList() change.List {
	return g.changes
}
