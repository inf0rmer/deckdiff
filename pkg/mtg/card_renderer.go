package mtg

import "fmt"

type CardRenderer interface {
	Render(c *Card) string
}

type DefaultCardRenderer struct{}

func NewDefaultCardRenderer() CardRenderer {
	return DefaultCardRenderer{}
}

func (r DefaultCardRenderer) Render(c *Card) string {
	adjustment := ""

	switch c.Adjustment {
	case Addition:
		adjustment = "+"
	case Subtraction:
		adjustment = "-"
	}

	return fmt.Sprintf("%s%d %s", adjustment, c.Quantity, c.Name)
}
