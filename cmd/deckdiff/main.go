package main

import (
	"bufio"
	"flag"
	"fmt"
	"path"
	"strings"

	"github.com/hairyhenderson/go-fsimpl"
	"github.com/hairyhenderson/go-fsimpl/filefs"
	"github.com/hairyhenderson/go-fsimpl/httpfs"
	"github.com/inf0rmer/deckdiff/pkg/mtg"
)

func main() {
	oldPtr := flag.String("old", "", "path to a decklist in MTGO format")
	newPtr := flag.String("new", "", "path to a decklist in MTGO format")

	flag.Parse()

	oldDeck, err := loadDeck(*oldPtr)
	check(err)

	newDeck, err := loadDeck(*newPtr)
	check(err)

	diff := mtg.Diff(*oldDeck, *newDeck)

	fmt.Print(diff)
}

func loadDeck(p string) (deck *mtg.Decklist, err error) {
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
	deck = &mtg.Decklist{}
	var isSideboard bool = false

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 && !isSideboard {
			isSideboard = true
			continue
		}

		crd, err := mtg.ParseLine(line)

		if err != nil {
			return nil, err
		}

		if !isSideboard {
			deck.Mainboard = append(deck.Mainboard, crd)
		} else {
			deck.Sideboard = append(deck.Sideboard, crd)
		}

	}

	return deck, nil
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
