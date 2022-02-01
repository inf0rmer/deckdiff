package mtg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultCardRendererWithAddition(t *testing.T) {
	renderer := NewDefaultCardRenderer()

	card := NewCard("foo", 2, Addition)

	expected := "+2 foo"
	actual := renderer.Render(card)

	assert.Equal(t, expected, actual)
}

func TestDefaultCardRendererWithSubtraction(t *testing.T) {
	renderer := NewDefaultCardRenderer()

	card := NewCard("foo", 2, Subtraction)

	expected := "-2 foo"
	actual := renderer.Render(card)

	assert.Equal(t, expected, actual)
}
