package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MacroCaseReplacer(t *testing.T) {
	r := newReplacerMacroCase()

	testCases := map[string]string{
		"a":   "A",
		"a_b": "A_B",
		"aB":  "A_B",
		"abC": "AB_C",
		"aBC": "A_BC",
	}

	for input, expected := range testCases {
		actual := r.Replace(input)
		assert.Equal(t, expected, actual, input)
	}
}
