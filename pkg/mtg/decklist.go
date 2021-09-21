package mtg

import "strings"

type Decklist struct {
	Mainboard []*Card
	Sideboard []*Card
	renderer  CardRenderer
}

func NewDecklist(mainboard []*Card, sideboard []*Card, renderer CardRenderer) *Decklist {
	return &Decklist{
		Mainboard: mainboard,
		Sideboard: sideboard,
		renderer:  renderer,
	}
}

func (d Decklist) String() string {
	var result strings.Builder

	result.WriteString("Main Deck: \n")

	for _, c := range d.Mainboard {
		result.WriteString(d.renderer.Render(c))
	}

	if len(d.Sideboard) > 0 {
		result.WriteString("\nSideboard: \n")
	}

	for _, c := range d.Sideboard {
		result.WriteString(d.renderer.Render(c))
	}

	return result.String()
}
