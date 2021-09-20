package mtg

import "fmt"

type Card struct {
	Adjustment string
	Quantity   int64
	Name       string
}

func (c Card) String() string {
	return fmt.Sprintf("%s%d %s \n", c.Adjustment, c.Quantity, c.Name)
}

func FindCard(name string, list []Card) (result Card) {
	result = Card{Name: name, Quantity: 0}

	for _, c := range list {
		if c.Name == name {
			result = c
			break
		}
	}

	return
}
