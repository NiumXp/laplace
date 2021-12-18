package internal

type Game struct {
	Map   *Map
	Heros []*Hero

	states map[uint]*State
}

func NewGame(m *Map, h []*Hero) *Game {
	return &Game{
		Map:   m,
		Heros: h,

		states: make(map[uint]*State),
	}
}

func (g *Game) StateAt(instant uint) *State {
	if s, ok := g.states[instant]; ok {
		return s
	}

	m := g.Map
	h := g.Heros

	if instant == 0 {
		state := NewState(
			instant, m,
			m.RandomOpenPositions(len(h)),
		)

		g.states[instant] = state
		return state
	}

	state := NewState(instant, m, nil)

	return state
}
