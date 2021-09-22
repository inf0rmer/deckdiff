package parser

type DecklistParser interface {
	Parse(input string) string
}
