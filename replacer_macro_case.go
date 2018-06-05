package main

import (
	"strings"
	"unicode"
)

// This will attempt to convert straight-forward and conventional casing;
// it is not intended to replace complex or non-intuitive casing
type replacerMacroCase struct{}

func newReplacerMacroCase() replacerMacroCase {
	return replacerMacroCase{}
}

func (r replacerMacroCase) Replace(input string) string {
	if len(input) == 0 {
		return ""
	}

	var output []rune

	matchUppercase := unicode.IsUpper(rune(input[0]))

	for _, r := range input {
		if !unicode.IsLetter(r) {
			output = append(output, '_')
			continue
		}

		if matchUppercase && !unicode.IsUpper(r) {
			output = append(output, unicode.ToUpper(r))
			matchUppercase = !matchUppercase
			continue
		}

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
