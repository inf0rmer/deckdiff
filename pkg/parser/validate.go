package parser

import (
	"bufio"
	"regexp"
	"strings"
)

func validateDecklist(list string) (valid bool) {
	valid = true
	scanner := bufio.NewScanner(strings.NewReader(list))

	if len(list) == 0 {
		valid = false
		return
	}

	for scanner.Scan() {
		line := scanner.Text()
		valid = validateLine(line)

		if !valid {
			break
		}
	}

	return
}

func validateLine(line string) bool {
	lineR := regexp.MustCompile(`(?m)(?P<Quantity>\d)\s(?P<Name>.*)`)

	return lineR.MatchString(line) || len(line) == 0
}
