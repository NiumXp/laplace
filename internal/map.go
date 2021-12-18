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

// NewMap creates a new map with the given width and height.
func NewMap(width, height uint) *Map {
	m := make(Map, height)

	for y := range m {
		m[y] = make([]uint, width)
	}

	return &m
}

// RandomOpenPosition returns a random open position (NULL) on the map.
func (m *Map) RandomOpenPosition() *Position {
	t := (*m)

	w := len(t[0])
	h := len(t)

	for {
		x := uint(rand.Intn(w))
		y := uint(rand.Intn(h))

		if t[y][x] == NULL {
			return NewPosition(x, y)
		}
	}
}

// RandomOpenPositions returns a slice of `n` random open positions (NULL) on the map.
func (m *Map) RandomOpenPositions(n int) []*Position {
	positions := make([]*Position, n)

	for i := 0; i < n; i++ {
		positions[i] = m.RandomOpenPosition()
	}

	return positions
}
