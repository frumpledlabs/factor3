package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MacroCaseReplacer(t *testing.T) {
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
		"ApiKey":        "API_KEY",
		"API_Key":       "API_KEY",
		"AnotherAPIKey": "ANOTHER_API_KEY",
		"anotherAPIKey": "ANOTHER_API_KEY",
		"someJSON":      "SOME_JSON",
		"moarJSONData":  "MOAR_JSON_DATA",
		"JSONPi":        "JSON_PI",
	}

	for input, expected := range testCases {
		actual := Replace(input)
		assert.Equal(t, expected, actual, input)
	}
}

func Test_SpecialCase(t *testing.T) {
	testCases := map[string]string{
		"anotherAPIKey": "ANOTHER_API_KEY",
	}

	for input, expected := range testCases {
		actual := Replace(input)
		assert.Equal(t, expected, actual, input)
	}
}
