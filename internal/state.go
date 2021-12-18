package internal

import "fmt"

// State represents a state of the game.
type State struct {
	Instant uint
	Map     *Map
	Heroes  []*Position
}

// NewState creates a new state with the given instant, map and heroes.
func NewState(instant uint, m *Map, p []*Position) *State {
	return &State{
		Instant: instant,
		Map:     m,
		Heroes:  p,
	}
}

// String returns a string representation of the state.
func (s *State) String() string {
	return fmt.Sprintf(
		"State{\n\tInstant: %d,\n\tHeros: %v,\n}",
		s.Instant,
		s.Heroes,
	)
}
