package main

import (
	"flag"
	"fmt"
	"net/url"

	"github.com/inf0rmer/deckdiff/pkg/cli"
	"github.com/inf0rmer/deckdiff/pkg/mtg"
	"github.com/inf0rmer/deckdiff/pkg/parser"
)

func main() {
	oldPtr := flag.String("old", "", "path to a decklist in MTGO format")
	newPtr := flag.String("new", "", "path to a decklist in MTGO format")
	flag.Parse()

	oldUrl, err := url.Parse(*oldPtr)

	check(err)

	newUrl, err := url.Parse(*newPtr)

	check(err)

	oldDeck, err := mtg.LoadDeck(oldUrl, getParser(oldUrl))
	check(err)

	newDeck, err := mtg.LoadDeck(newUrl, getParser(newUrl))
	check(err)

	diff := mtg.Diff(oldDeck, newDeck, cli.NewCliCardRenderer())

	fmt.Print(diff)
}

func getParser(u *url.URL) parser.DecklistParser {
	switch u.Host {
	case "deckbox.org":
		return parser.NewDeckboxParser()
	}

	return parser.NewIdentityParser()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
