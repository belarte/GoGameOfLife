package life

import (
	"github.com/belarte/GoGameOfLife/struct/change"
	"github.com/belarte/GoGameOfLife/struct/list"
)

type Game interface {
	Simulate(iterations int)
	Step()
	NextChangeList() change.List
}

type game struct{}

func New() Game {
	return &game{}
}

func (g *game) Simulate(iterations int) {
}

func (g *game) Step() {
}

func (g *game) NextChangeList() change.List {
	return change.New(list.New(), list.New())
}
