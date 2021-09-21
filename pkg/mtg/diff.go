package mtg

func Diff(a Decklist, b Decklist, r CardRenderer) (result *Decklist) {
	return NewDecklist(diffList(a.Mainboard, b.Mainboard), diffList(a.Sideboard, b.Sideboard), r)
}

func diffList(a []*Card, b []*Card) (result []*Card) {
	for _, c := range b {
		oldCard := FindCard(c.Name, a)

		if c.Quantity != oldCard.Quantity {
			newQuantity := c.Quantity - oldCard.Quantity
			adjustment := Addition

			if oldCard.Quantity > newQuantity {
				adjustment = Subtraction
			}

			if newQuantity < 0 {
				newQuantity = -newQuantity
			}

			result = append(result, NewCard(c.Name, newQuantity, adjustment))
		}
	}

	for _, c := range a {
		newCard := FindCard(c.Name, b)

		if newCard.Quantity == 0 {
			result = append(result, NewCard(c.Name, c.Quantity, Subtraction))
		}
	}

	return
}
