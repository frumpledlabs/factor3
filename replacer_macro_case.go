package main

import (
	"strings"
	"unicode"
)

// This will attempt to convert straight-forward and conventional casing;
// it is not intended to replace complex or non-intuitive casing
func Replace(input string) string {
	if len(input) == 0 {
		return ""
	}

	var output []rune

	matchUppercase := unicode.IsUpper(rune(input[0]))
	var nextIsUppercase bool
	var nextNextIsUppercase bool

	for i := 0; i < len(input); i++ {
		r := rune(input[i])

		if !unicode.IsLetter(r) {
			output = append(output, '_')
			continue
		}

		// Determine case of current rune's successor:
		if i < len(input)-1 {
			nextIsUppercase = unicode.IsUpper(rune(input[i+1]))
		}

		// Determine case of next rune's successor:
		if i < len(input)-2 {
			nextNextIsUppercase = unicode.IsUpper(rune(input[i+2]))
		}

		// Account for Mixes uppercase:
		if unicode.IsUpper(r) && nextIsUppercase && !nextNextIsUppercase {
			output = append(output, unicode.ToUpper(r))
			output = append(output, '_')
			continue
		}

		// Account for upper to lower case change:
		if matchUppercase && !unicode.IsUpper(r) {
			output = append(output, unicode.ToUpper(r))
			matchUppercase = !matchUppercase
			continue
		}

		// Account for lower to upper case change:
		if !matchUppercase && unicode.IsUpper(r) {
			output = append(output, '_')
			output = append(output, unicode.ToUpper(r))
			matchUppercase = !matchUppercase
			continue
		}

		output = append(output, unicode.ToUpper(r))
	}

	return strings.Trim(string(output), "_")
}
