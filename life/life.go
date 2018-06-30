package life

import (
	"github.com/belarte/GoGameOfLife/struct/change"
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
	return &game{}
}

func (g *game) Add(changes change.List) {
	g.changes.Alive = append(g.changes.Alive, changes.Alive...)
	g.changes.Dead = append(g.changes.Dead, changes.Dead...)
}

func (g *game) Simulate(iterations int) {
}

func (g *game) Step() {
}

func (g *game) NextChangeList() change.List {
	return g.changes
}
