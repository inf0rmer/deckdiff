package parser

import (
	"log"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type DeckboxParser struct{}

func NewDeckboxParser() *DeckboxParser {
	return &DeckboxParser{}
}

func (p DeckboxParser) Parse(input string) (result string) {

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(input))
	if err != nil {
		log.Fatal(err)
	}

	result = doc.Find("body").Text()

	re := regexp.MustCompile(`(?m)[a-zA-Z]{0}(\d{1})`)
	result = re.ReplaceAllString(result, "\n${1}")

	re = regexp.MustCompile(`(?m).*Sideboard:`)
	result = re.ReplaceAllString(result, "")

	re = regexp.MustCompile(`(?m)\n\s*\n`)
	result = re.ReplaceAllString(result, "\n\n")

	result = strings.Trim(result, "\n")

	return
}
