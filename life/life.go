package life

import (
	"github.com/belarte/GoGameOfLife/struct/change"
	"github.com/belarte/GoGameOfLife/struct/coord"
	"github.com/belarte/GoGameOfLife/struct/grid"
	"github.com/belarte/GoGameOfLife/struct/list"
)

type Game interface {
	Add(changes change.List)
	Simulate(iterations int)
	Step()
	NextChangeList() change.List
}

type game struct {
	grid.Grid
	changes change.List
}

func New(width, height int) Game {
	list := change.New(list.New(), list.New())
	grid := grid.New(width, height)
	return &game{changes: list, Grid: grid}
}

func (g *game) Add(changes change.List) {
	g.changes.Alive.Concat(changes.Alive)
	g.changes.Dead.Concat(changes.Dead)
}

func (g *game) Simulate(iterations int) {
}

func (g *game) Step() {
	cover := g.changes.Cover()

	for c, _ := range g.changes.Alive {
		g.Set(c, 1)
	}

	for c, _ := range g.changes.Dead {
		g.Set(c, 0)
	}

	g.changes.Alive = make(map[coord.Coord]bool)
	g.changes.Dead = make(map[coord.Coord]bool)

	for c, _ := range cover {
		wrappedCoord := g.Wrap(c)
		state := g.Get(wrappedCoord)
		count := g.countAliveNeighbours(wrappedCoord)
		if state == 0 && count == 3 {
			g.changes.Alive.Append(wrappedCoord)
		}
		if state != 0 && !(count == 2 || count == 3) {
			g.changes.Dead.Append(wrappedCoord)
		}
	}
}

func (g *game) NextChangeList() change.List {
	return g.changes
}

func (g game) countAliveNeighbours(c coord.Coord) int {
	return g.Get(g.Wrap(coord.New(c.X-1, c.Y-1))) +
		g.Get(g.Wrap(coord.New(c.X, c.Y-1))) +
		g.Get(g.Wrap(coord.New(c.X+1, c.Y-1))) +
		g.Get(g.Wrap(coord.New(c.X-1, c.Y))) +
		g.Get(g.Wrap(coord.New(c.X+1, c.Y))) +
		g.Get(g.Wrap(coord.New(c.X-1, c.Y+1))) +
		g.Get(g.Wrap(coord.New(c.X, c.Y+1))) +
		g.Get(g.Wrap(coord.New(c.X+1, c.Y+1)))
}
