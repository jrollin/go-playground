package molkky

// Molkky game
type Molkky struct {
	score int
	rolls []int
}

func startMolkky() Molkky {
	return Molkky{score: 0}
}

func (m *Molkky) getScore() int {
	return m.score
}

func (m *Molkky) setFallenQuells(quells []int) {
	m.rolls = quells
	if len(quells) == 0 {
		m.score += 0
	} else if len(quells) > 1 {
		m.score += len(quells)
	} else {
		m.score += quells[0]
	}
}
