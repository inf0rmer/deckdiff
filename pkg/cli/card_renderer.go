package cli

import (
	"github.com/fatih/color"

	"github.com/inf0rmer/deckdiff/pkg/mtg"
)

type CliCardRenderer struct{}

func NewCliRenderer() mtg.CardRenderer {
	return CliCardRenderer{}
}

func (r CliCardRenderer) Render(c *mtg.Card) string {
	adjustment := ""
	var d *color.Color

	switch c.Adjustment {
	case mtg.Addition:
		adjustment = "+"
		d = color.New(color.FgGreen)
	case mtg.Subtraction:
		adjustment = "-"
		d = color.New(color.FgRed)
	}

	return d.Sprintf("%s%d %s \n", adjustment, c.Quantity, c.Name)
}
