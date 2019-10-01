package factor3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_TagSetRecognizesBuiltinTags(t *testing.T) {
	input := `required`
	tagSet := newTagSet(input)

	assert.True(t, tagSet.isRequired)
}

func Test_OverrideValueIsParsed(t *testing.T) {
	input := `${SOMETHING_ELSE},required`
	tagSet := newTagSet(input)

	assert.Equal(t, "SOMETHING_ELSE", tagSet.overrideKey)
}

func Test_DefaultValueIsParsed(t *testing.T) {
	input := `${:-DEFAULT_VALUE}`
	tagSet := newTagSet(input)

	assert.Equal(t, "DEFAULT_VALUE", tagSet.defaultValue)
}

func Test_OverrideValueIsParsedWithoutDefaultValue(t *testing.T) {
	input := `${OVERRIDE:-},required`
	tagSet := newTagSet(input)

	assert.Equal(t, "OVERRIDE", tagSet.overrideKey)
}

// func Test_MalformedInputsReturnError(t *testing.T) {
// 	testInputs := []string{
// 		`${A:B},required`, // invalid value/default separator
// 		`{A}`,             // Missing leading $
// 		`${:-},`,          // trailing comma / empty tag
// 	}

// 	for _, input := range testInputs {
// 		_, err := newTagSet(input)
// 		assert.Nil(t, err)
// 	}
// }
