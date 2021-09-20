package main

import (
	"bufio"
	"flag"
	"fmt"
	"path"
	"regexp"
	"strconv"
	"strings"

	"github.com/hairyhenderson/go-fsimpl"
	"github.com/hairyhenderson/go-fsimpl/filefs"
	"github.com/hairyhenderson/go-fsimpl/httpfs"
)

type card struct {
	adjustment string
	quantity   int64
	name       string
}

type decklist struct {
	mainboard []card
	sideboard []card
}

func (c card) String() string {
	return fmt.Sprintf("%s%d %s \n", c.adjustment, c.quantity, c.name)
}

func (d decklist) String() string {
	var result strings.Builder

	result.WriteString("Main Deck: \n")

	for _, c := range d.mainboard {
		result.WriteString(c.String())
	}

	if len(d.sideboard) > 0 {
		result.WriteString("\nSideboard: \n")
	}

	for _, c := range d.sideboard {
		result.WriteString(c.String())
	}

	return result.String()
}

func main() {
	oldPtr := flag.String("old", "", "path to a decklist in MTGO format")
	newPtr := flag.String("new", "", "path to a decklist in MTGO format")

	flag.Parse()

	oldDeck, err := loadDeck(*oldPtr)
	check(err)

	newDeck, err := loadDeck(*newPtr)
	check(err)

	diff := diff(*oldDeck, *newDeck)

	fmt.Print(diff)
}

func loadDeck(p string) (deck *decklist, err error) {
	mux := fsimpl.NewMux()
	mux.Add(filefs.FS)
	mux.Add(httpfs.FS)

	if err != nil {
		return nil, err
	}

	fsys, err := mux.Lookup(strings.TrimSuffix(p, path.Base(p)))

	if err != nil {
		return nil, err
	}

	file, err := fsys.Open(path.Base(p))

	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	deck = &decklist{}
	var isSideboard bool = false

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 && !isSideboard {
			isSideboard = true
			continue
		}

		crd, err := parseLine(line)

		if err != nil {
			return nil, err
		}

		if !isSideboard {
			deck.mainboard = append(deck.mainboard, crd)
		} else {
			deck.sideboard = append(deck.sideboard, crd)
		}

	}

	return deck, nil
}

func parseLine(line string) (crd card, err error) {
	lineR := regexp.MustCompile(`(?m)(?P<Quantity>\d)\s(?P<Name>.*)`)

	results := lineR.FindStringSubmatch(line)

	if !lineR.MatchString(line) {
		err = fmt.Errorf("line is malformed: %s", line)

		return card{}, err
	}

	quantity, err := strconv.ParseInt(results[lineR.SubexpIndex("Quantity")], 0, 64)
	name := results[lineR.SubexpIndex("Name")]
	crd = card{quantity: quantity, name: name}

	return
}

func diff(a decklist, b decklist) (result decklist) {
	result = decklist{
		mainboard: diffList(a.mainboard, b.mainboard),
		sideboard: diffList(a.sideboard, b.sideboard),
	}

	return
}

func diffList(a []card, b []card) (result []card) {
	for _, c := range b {
		oldCard := findCard(c.name, a)

		if c.quantity != oldCard.quantity {
			newQuantity := c.quantity - oldCard.quantity
			adjustment := "+"

			if oldCard.quantity > newQuantity {
				adjustment = "-"
			}

			if newQuantity < 0 {
				newQuantity = -newQuantity
			}

			result = append(result, card{name: c.name, quantity: newQuantity, adjustment: adjustment})
		}
	}

	return
}

func findCard(name string, list []card) (result card) {
	result = card{name: name, quantity: 0}

	for _, c := range list {
		if c.name == name {
			result = c
			break
		}
	}

	return
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
