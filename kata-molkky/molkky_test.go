package molkky

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMolkkyStartWithScoreZero(t *testing.T) {
	m := startMolkky()
	// given
	expected := 0
	// than
	actual := m.getScore()
	// that
	assert.Equal(t, expected, actual)
}

func TestMolkkyScoreWithNoFallenQuellsl(t *testing.T) {
	m := startMolkky()
	m.setFallenQuells([]int{})
	// given
	expected := 0
	// than
	actual := m.getScore()
	// that
	assert.Equal(t, expected, actual)
}

func TestMolkkyScoreTwoPointsIfTwoFallenQuells(t *testing.T) {
	m := startMolkky()
	m.setFallenQuells([]int{3, 12})
	// given
	expected := 2
	// than
	actual := m.getScore()
	// that
	assert.Equal(t, expected, actual)
}

func TestMolkkyScoreQuellPointsIfOnlyOneQuell(t *testing.T) {
	m := startMolkky()
	m.setFallenQuells([]int{3})
	// given
	expected := 3
	// than
	actual := m.getScore()
	// that
	assert.Equal(t, expected, actual)
}
