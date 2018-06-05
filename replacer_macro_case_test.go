package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MacroCaseReplacer(t *testing.T) {
	r := newReplacerMacroCase()

	testCases := map[string]string{
		"A":          "A",
		"lowercase":  "LOWERCASE",
		"UPPERCASE":  "UPPERCASE",
		"camelCase":  "CAMEL_CASE",
		"snake_case": "SNAKE_CASE",
		"oOpS_CaSe":  "O_OP_S_CA_SE",

		"aBC":      "A_BC",
		"aBcD":     "A_BC_D",
		"aa_bC_De": "AA_B_C_DE",
	}

	for input, expected := range testCases {
		actual := r.Replace(input)
		assert.Equal(t, expected, actual, input)
	}
}
