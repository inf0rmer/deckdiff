package parser

import "github.com/inf0rmer/deckdiff/pkg/mtg"

type IdentityParser struct{}

func NewIdentityParser() *IdentityParser {
	return &IdentityParser{}
}

func (p *IdentityParser) Parse(input string) (*mtg.Decklist, error) {
	return toDecklist(input)
}
