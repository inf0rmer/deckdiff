package mtg

import "strings"

type Decklist struct {
	Mainboard []Card
	Sideboard []Card
}

func (d Decklist) String() string {
	var result strings.Builder

	result.WriteString("Main Deck: \n")

	for _, c := range d.Mainboard {
		result.WriteString(c.String())
	}

	if len(d.Sideboard) > 0 {
		result.WriteString("\nSideboard: \n")
	}

	for _, c := range d.Sideboard {
		result.WriteString(c.String())
	}

	return result.String()
}
