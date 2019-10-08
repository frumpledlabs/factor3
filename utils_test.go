package factor3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MatchVar(t *testing.T) {
	input := "${SOMETHING:-interesting},required"

	match := newMatchVar(input)

	assert.Equal(t, "SOMETHING", match.varName)
	assert.Equal(t, "interesting", match.defaultValue)
}

func Test_NonMatchingVar(t *testing.T) {
	match := newMatchVar(`$$$${PASSED:-}`)

	assert.Equal(t, "", match.varName)
	assert.Equal(t, "", match.defaultValue)
}

func Test_MatchOnlyVarNameValue(t *testing.T) {
	match := newMatchVar(`${PASSED:-}`)

	assert.Equal(t, "PASSED", match.varName)
}

func Test_MatchOnlyDefaultValue(t *testing.T) {
	match := newMatchVar(`${:-PASSED}`)

	assert.Equal(t, "PASSED", match.defaultValue)
}
