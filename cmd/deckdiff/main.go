package main

import (
	"flag"
	"fmt"

	"github.com/inf0rmer/deckdiff/pkg/mtg"
)

func main() {
	oldPtr := flag.String("old", "", "path to a decklist in MTGO format")
	newPtr := flag.String("new", "", "path to a decklist in MTGO format")

	flag.Parse()

	oldDeck, err := mtg.LoadDeck(*oldPtr)
	check(err)

	newDeck, err := mtg.LoadDeck(*newPtr)
	check(err)

	diff := mtg.Diff(*oldDeck, *newDeck)

	fmt.Print(diff)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
