package internal

import "fmt"

type State struct {
	Instant uint
	Map     *Map
	Heros   []*Position
}

func NewState(instant uint, m *Map, p []*Position) *State {
	return &State{
		Instant: instant,
		Map:     m,
		Heros:   p,
	}
}

func (s *State) String() string {
	return fmt.Sprintf("State{Instant: %d}", s.Instant)
}
