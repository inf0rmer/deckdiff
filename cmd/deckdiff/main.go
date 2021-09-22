package main

import (
	"flag"
	"fmt"

	"github.com/inf0rmer/deckdiff/pkg/cli"
	"github.com/inf0rmer/deckdiff/pkg/mtg"
	"github.com/inf0rmer/deckdiff/pkg/parser"
)

func main() {
	oldPtr := flag.String("old", "", "path to a decklist in MTGO format")
	newPtr := flag.String("new", "", "path to a decklist in MTGO format")

	flag.Parse()

	oldDeck, err := mtg.LoadDeck(*oldPtr, parser.NewIdentityParser())
	check(err)

	newDeck, err := mtg.LoadDeck(*newPtr, parser.NewIdentityParser())
	check(err)

	diff := mtg.Diff(*oldDeck, *newDeck, cli.NewCliRenderer())

	fmt.Print(diff)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
