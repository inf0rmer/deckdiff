package parser

type IdentityParser struct{}

func NewIdentityParser() *IdentityParser {
	return &IdentityParser{}
}

func (p *IdentityParser) Parse(input string) string {
	return input
}
