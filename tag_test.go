package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_SingleValueTagParsesKey(t *testing.T) {
	testTag, err := newTag(`key:"value"`)

	require.Nil(t, err)
	assert.Equal(t, "key", testTag.Key())
}

func Test_ParseSingleValue(t *testing.T) {
	testTag, err := newTag(`key:"value"`)

	require.Nil(t, err)
	assert.Equal(t, []string{"value"}, testTag.Values())
}

func Test_ParseMultipleValues(t *testing.T) {
	testTag, err := newTag(`key:"value"`)

	require.Nil(t, err)
	assert.Equal(t, []string{"value"}, testTag.Values())
}

func Test_ParseMultiValueTag_FindsAllValues(t *testing.T) {
	testTag, err := newTag(`key:"value,test"`)

	require.Nil(t, err)
	assert.Equal(t, 2, len(testTag.Values()))
}

func Test_ParseInvalidTagsResultsInError(t *testing.T) {

	invalidTags := []string{
		`:"value"`,
		`key:`,
		`:`,
		` :"value"`,
		`key:"value,"`,
	}

	for _, tagInput := range invalidTags {
		_, err := newTag(tagInput)
		assert.NotNil(t, err, fmt.Sprintf("Invalid tag input: '%s'", tagInput))
	}
}
