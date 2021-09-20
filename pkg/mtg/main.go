package mtg

import (
	"fmt"
	"regexp"
	"strconv"
)

func Diff(a Decklist, b Decklist) (result Decklist) {
	result = Decklist{
		Mainboard: diffList(a.Mainboard, b.Mainboard),
		Sideboard: diffList(a.Sideboard, b.Sideboard),
	}

	return
}

func ParseLine(line string) (crd Card, err error) {
	lineR := regexp.MustCompile(`(?m)(?P<Quantity>\d)\s(?P<Name>.*)`)

	results := lineR.FindStringSubmatch(line)

	if !lineR.MatchString(line) {
		err = fmt.Errorf("line is malformed: %s", line)

		return Card{}, err
	}

	quantity, err := strconv.ParseInt(results[lineR.SubexpIndex("Quantity")], 0, 64)
	name := results[lineR.SubexpIndex("Name")]
	crd = Card{Quantity: quantity, Name: name}

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
