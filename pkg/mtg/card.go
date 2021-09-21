package mtg

import "fmt"

type Card struct {
	Adjustment Adjustment
	Quantity   int64
	Name       string
}

type Adjustment int64

const (
	Addition Adjustment = iota
	Subtraction
)

func (a Adjustment) String() string {
	switch a {
	case Addition:
		return "+"
	case Subtraction:
		return "-"
	}
	return ""
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
