package mtg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDecklist(t *testing.T) {
	mainboard := []*Card{NewCard("mainboard 1", 2, None), NewCard("mainboard 2", 4, None)}
	sideboard := []*Card{NewCard("sideboard", 3, None)}

	decklist := NewDecklist(mainboard, sideboard, NewDefaultCardRenderer())

	assert.Equal(t, mainboard, decklist.Mainboard)
	assert.Equal(t, sideboard, decklist.Sideboard)
}

func TestDecklistString(t *testing.T) {
	mainboard := []*Card{NewCard("mainboard 1", 2, None), NewCard("mainboard 2", 4, None)}
	sideboard := []*Card{NewCard("sideboard", 3, None)}

	decklist := NewDecklist(mainboard, sideboard, NewDefaultCardRenderer())

	actual := decklist.String()
	expected := `Main Deck:
2 mainboard 1
4 mainboard 2

Sideboard:
3 sideboard
`
	assert.Equal(t, expected, actual)
}
