package factor3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MacroCaseReplacer(t *testing.T) {
	macroCaser := newMacroCaseReplacer()

	testCases := map[string]string{
		// Documentation tests:
		"":           "",
		"A":          "A",
		"oOpS_C":     "O_OP_S_C",
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

		// Mixed Character Tests:
		"Int64Var": "INT64_VAR",
	}

	for input, expected := range testCases {
		actual := macroCaser.Replace(input)
		assert.Equal(t, expected, actual, input)
	}
}
