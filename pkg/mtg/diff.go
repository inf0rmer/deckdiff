package mtg

func Diff(a Decklist, b Decklist) (result Decklist) {
	result = Decklist{
		Mainboard: diffList(a.Mainboard, b.Mainboard),
		Sideboard: diffList(a.Sideboard, b.Sideboard),
	}

	return
}

func diffList(a []Card, b []Card) (result []Card) {
	for _, c := range b {
		oldCard := FindCard(c.Name, a)

		if c.Quantity != oldCard.Quantity {
			newQuantity := c.Quantity - oldCard.Quantity
			adjustment := "+"

			if oldCard.Quantity > newQuantity {
				adjustment = "-"
			}

			if newQuantity < 0 {
				newQuantity = -newQuantity
			}

			result = append(result, Card{Name: c.Name, Quantity: newQuantity, Adjustment: adjustment})
		}
	}

	for _, c := range a {
		newCard := FindCard(c.Name, b)

		if newCard.Quantity == 0 {
			result = append(result, Card{Name: c.Name, Quantity: c.Quantity, Adjustment: "-"})
		}
	}

	return
}
