package internal

// Hero represents a hero in the game.
type Hero struct {
	// Power is the damage of a bomb.
	Power uint8
	Speed uint8
	// Stamina is the level of stamina of the Hero.
	Stamina uint8
	// Bombs is the amount of bombs can be dropped.
	Bombs uint8
	// Range is the range of the bombs.
	Range uint8

	// Baterry is if the Hero has a baterry.
	Baterry bool
	// Mana is if the Hero has the mana powerup.
	Mana bool
}

func (h *Hero) CanMoveAt(instant uint) bool
