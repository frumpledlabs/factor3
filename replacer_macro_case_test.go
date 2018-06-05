package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MacroCaseReplacer(t *testing.T) {
	r := newReplacerMacroCase()

	testCases := map[string]string{
		// Documentation tests:
		"":           "",
		"A":          "A",
		"lowercase":  "LOWERCASE",
		"UPPERCASE":  "UPPERCASE",
		"camelCase":  "CAMEL_CASE",
		"snake_case": "SNAKE_CASE",
		"oOpS_CaSe":  "O_OP_S_CA_SE",

		// Practical tests:
		"APIKey":        "API_KEY",
		"anotherAPIKey": "ANOTHER_API_KEY",
		"someJSON":      "SOME_JSON",
		"JSONData":      "JSON_DATA",
	}

	for input, expected := range testCases {
		actual := r.Replace(input)
		assert.Equal(t, expected, actual, input)
	}
}
