package internal

import "fmt"

// Position represents a position in the game.
type Position [2]uint

// NewPosition creates a new position with the given x and y.
func NewPosition(x, y uint) *Position {
	return &Position{x, y}
}

// String returns a string representation of the position.
func (p *Position) String() string {
	return fmt.Sprintf("(%d, %d)", p[0], p[1])
}
