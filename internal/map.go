package internal

import (
	"math/rand"
)

const (
	NULL = iota
	WALL

	// Chest colors
	WHITE
	BROWN
	PURPLE
	YELLOW
	BLUE
)

type Map [][]uint

func NewMap(width, height uint) *Map {
	m := make(Map, height)

	for y := range m {
		m[y] = make([]uint, width)
	}

	return &m
}

func (m *Map) RandomOpenPosition() *Position {
	t := (*m)

	w := len(t[0])
	h := len(t)

	for {
		x := uint(rand.Intn(w))
		y := uint(rand.Intn(h))

		if t[y][x] == NULL {
			return &Position{x, y}
		}
	}
}

func (m *Map) RandomOpenPositions(n int) []*Position {
	positions := make([]*Position, n)

	for i := range positions {
		positions[i] = m.RandomOpenPosition()
	}

	return positions
}
