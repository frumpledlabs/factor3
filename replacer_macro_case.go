package main

import (
	//"fmt"
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

		// Determine casing:
		currentIsUppercase := unicode.IsUpper(r)

		if i < len(input)-1 {
			nextIsUppercase = unicode.IsUpper(rune(input[i+1]))
		}

		if i < len(input)-2 {
			nextNextIsUppercase = unicode.IsUpper(rune(input[i+2]))
		}

		// Account for Mixes uppercase:
		if i < len(input)-2 && !unicode.IsLetter(rune(input[i+1])) {
			output = append(output, unicode.ToUpper(r))
			continue
		}

		if !currentIsUppercase && nextIsUppercase {
			output = append(output, unicode.ToUpper(r))
			output = append(output, '_')
			continue
		}

		if currentIsUppercase && nextIsUppercase && !nextNextIsUppercase {
			if i < len(input)-2 {
				if !unicode.IsLetter(rune(input[i+2])) {
					output = append(output, unicode.ToUpper(r))
					continue
				}
			}

			output = append(output, unicode.ToUpper(r))
			output = append(output, '_')
			continue
		}

		// Account for upper to lower case change:
		if matchUppercase && !currentIsUppercase {
			output = append(output, unicode.ToUpper(r))
			matchUppercase = !matchUppercase
			continue
		}

		// Account for lower to upper case change:
		if !matchUppercase && currentIsUppercase {
			output = append(output, '_')
			output = append(output, unicode.ToUpper(r))
			matchUppercase = !matchUppercase
			continue
		}

		output = append(output, unicode.ToUpper(r))
	}

	return strings.Replace(strings.Trim(string(output), "_"), "__", "_", -1)
}
