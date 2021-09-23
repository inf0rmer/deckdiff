package mtg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCard(t *testing.T) {
	name := "foo"
	quantity := int64(3)
	adjustment := Addition

	card := NewCard(name, quantity, adjustment)

	assert.Equal(t, card.Name, name)
	assert.Equal(t, card.Quantity, quantity)
	assert.Equal(t, card.Adjustment, adjustment)
}

func TestFindCard(t *testing.T) {
	list := []*Card{NewCard("foo", 1, None), NewCard("bar", 2, None)}
	card := FindCard("foo", list)

	assert.Equal(t, card, list[0])
}
