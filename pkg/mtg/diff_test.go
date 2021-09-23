package mtg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecklistDiff(t *testing.T) {
	a := NewDecklist(
		[]*Card{NewCard("Blood Crypt", 2, None), NewCard("Blackcleave Cliffs", 4, None)},
		[]*Card{NewCard("Lurrus of the Dream Den", 1, None)},
		NewDefaultCardRenderer(),
	)

	b := NewDecklist(
		[]*Card{NewCard("Blood Crypt", 1, None), NewCard("Blackcleave Cliffs", 4, None)},
		[]*Card{NewCard("Lurrus of the Dream Den", 1, None), NewCard("Fatal Push", 4, None)},
		NewDefaultCardRenderer(),
	)

	actual := Diff(a, b, NewDefaultCardRenderer())
	expected := NewDecklist(
		[]*Card{NewCard("Blood Crypt", 1, Subtraction)},
		[]*Card{NewCard("Fatal Push", 4, Addition)},
		NewDefaultCardRenderer(),
	)

	assert.Equal(t, expected, actual)
}
