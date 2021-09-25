package parser

import (
	"fmt"
	"testing"

	"github.com/inf0rmer/deckdiff/pkg/mtg"
	"github.com/stretchr/testify/assert"
)

func TestIdentityParser(t *testing.T) {
	input := "3 Blood Crypt"

	expected := mtg.NewDecklist([]*mtg.Card{mtg.NewCard("Blood Crypt", 3, mtg.None)}, []*mtg.Card{}, nil)
	actual, err := NewIdentityParser().Parse(input)

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func TestIdentityParserErrorWithInvalidDecklist(t *testing.T) {
	input := "foobar"

	_, err := NewIdentityParser().Parse(input)

	assert.Errorf(t, err, fmt.Sprintf("Decklist is invalid: %s", input))
}

func TestIdentityParserErrorWithEmptyDecklist(t *testing.T) {
	input := ""

	_, err := NewIdentityParser().Parse(input)

	assert.Errorf(t, err, fmt.Sprintf("Decklist is invalid: %s", input))
}
