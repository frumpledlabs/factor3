package main

import (
	"strings"
	"unicode"
)

type replacerMacroCase struct{}

func newReplacerMacroCase() replacerMacroCase {
	return replacerMacroCase{}
}

func (r replacerMacroCase) OLD_Replace(input string) string {
	var replacement string

	fieldFunc := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c) && string(c) != "_"
	}

	//words := strings.Fields(input)
	words := strings.FieldsFunc(input, fieldFunc)
	for _, word := range words {
		replacement += strings.ToUpper(word) + "_"
	}

	replacement = strings.Trim(replacement, "_")

	return replacement
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
			output = append(output, '_')
			output = append(output, unicode.ToUpper(r))
			matchUppercase = false
			continue
		}

		if !matchUppercase && unicode.IsUpper(r) {
			output = append(output, '_')
			output = append(output, unicode.ToUpper(r))
			matchUppercase = true
			continue
		}

		output = append(output, unicode.ToUpper(r))
	}

	return strings.Trim(string(output), "_")
}
