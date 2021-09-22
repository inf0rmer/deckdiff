package mtg

import (
	"bufio"
	"bytes"
	"fmt"
	"net/url"
	"path"
	"regexp"
	"strconv"
	"strings"

	"github.com/hairyhenderson/go-fsimpl"
	"github.com/hairyhenderson/go-fsimpl/filefs"
	"github.com/hairyhenderson/go-fsimpl/httpfs"
	"github.com/inf0rmer/deckdiff/pkg/parser"
)

func LoadDeck(u *url.URL, prs parser.DecklistParser) (deck *Decklist, err error) {
	p := u.String()

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

	buf := new(bytes.Buffer)
	buf.ReadFrom(file)
	contents := prs.Parse(buf.String())

	scanner := bufio.NewScanner(strings.NewReader(contents))
	deck = NewDecklist(make([]*Card, 0), make([]*Card, 0), nil)
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
			deck.Mainboard = append(deck.Mainboard, crd)
		} else {
			deck.Sideboard = append(deck.Sideboard, crd)
		}

	}

	return deck, nil
}

func parseLine(line string) (crd *Card, err error) {
	lineR := regexp.MustCompile(`(?m)(?P<Quantity>\d)\s(?P<Name>.*)`)

	results := lineR.FindStringSubmatch(line)

	if !lineR.MatchString(line) {
		err = fmt.Errorf("line is malformed: %s", line)

		return nil, err
	}

	quantity, err := strconv.ParseInt(results[lineR.SubexpIndex("Quantity")], 0, 64)
	name := results[lineR.SubexpIndex("Name")]
	crd = NewCard(name, quantity, None)

	return
}
