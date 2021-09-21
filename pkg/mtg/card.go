package mtg

type Card struct {
	Adjustment Adjustment
	Quantity   int64
	Name       string
}

type Adjustment int64

const (
	Addition Adjustment = iota
	Subtraction
	None
)

func NewCard(name string, quantity int64, adjustment Adjustment) *Card {
	return &Card{
		Name:       name,
		Quantity:   quantity,
		Adjustment: adjustment,
	}
}

func FindCard(name string, list []*Card) (result *Card) {
	result = NewCard("", 0, None)

	for _, c := range list {
		if c.Name == name {
			result = c
			break
		}
	}

	return
}
