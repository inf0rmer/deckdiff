package parser

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/inf0rmer/deckdiff/pkg/mtg"
)

func toDecklist(input string) (decklist *mtg.Decklist, err error) {
	if !validateDecklist(input) {
		return nil, fmt.Errorf("Decklist is invalid: %s", input)
	}

	scanner := bufio.NewScanner(strings.NewReader(input))
	decklist = mtg.NewDecklist(make([]*mtg.Card, 0), make([]*mtg.Card, 0), nil)
	var isSideboard bool = false

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 && !isSideboard {
			isSideboard = true
			continue
		}

		crd, _ := parseLine(line)

		if !isSideboard {
			decklist.Mainboard = append(decklist.Mainboard, crd)
		} else {
			decklist.Sideboard = append(decklist.Sideboard, crd)
		}

	}

	return decklist, nil
}

func parseLine(line string) (crd *mtg.Card, err error) {
	lineR := regexp.MustCompile(`(?m)(?P<Quantity>\d)\s(?P<Name>.*)`)

	results := lineR.FindStringSubmatch(line)

	quantity, err := strconv.ParseInt(results[lineR.SubexpIndex("Quantity")], 0, 64)
	name := results[lineR.SubexpIndex("Name")]
	crd = mtg.NewCard(name, quantity, mtg.None)

	return
}
