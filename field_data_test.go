package factor3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FieldDataRecognizesBuiltinTags(t *testing.T) {
	input := `required`
	fieldData := newFieldData(input)

	assert.True(t, fieldData.isRequired)
}

func Test_OverrideKeyIsParsed(t *testing.T) {
	input := `${SOMETHING_ELSE},required`
	fieldData := newFieldData(input)

	assert.Equal(t, "SOMETHING_ELSE", fieldData.overrideKey)
}

func Test_DefaultValueIsParsed(t *testing.T) {
	input := `${:-DEFAULT_VALUE}`
	fieldData := newFieldData(input)

	assert.Equal(t, "DEFAULT_VALUE", fieldData.defaultValue)
}

func Test_OverrideValueIsParsedWithoutDefaultValue(t *testing.T) {
	input := `${OVERRIDE:-},required`
	fieldData := newFieldData(input)

	assert.Equal(t, "OVERRIDE", fieldData.overrideKey)
}
