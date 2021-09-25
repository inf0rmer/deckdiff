package parser

import "github.com/inf0rmer/deckdiff/pkg/mtg"

type DecklistParser interface {
	Parse(input string) (*mtg.Decklist, error)
}
