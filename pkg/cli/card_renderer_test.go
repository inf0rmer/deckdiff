package cli

import (
	"testing"

	"github.com/inf0rmer/deckdiff/pkg/mtg"
	"github.com/stretchr/testify/assert"
)

func TestCliCardRendererWithAddition(t *testing.T) {
	renderer := NewCliCardRenderer()

	card := mtg.NewCard("foo", 2, mtg.Addition)

	expected := "+2 foo"
	actual := renderer.Render(card)

	assert.Equal(t, expected, actual)
}

func TestCliCardRendererWithSubtraction(t *testing.T) {
	renderer := NewCliCardRenderer()

	card := mtg.NewCard("foo", 2, mtg.Subtraction)

	expected := "-2 foo"
	actual := renderer.Render(card)

	assert.Equal(t, expected, actual)
}
