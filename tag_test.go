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
	assert.Equal(t, "value", testTag.Value())
}

func Test_ParseMultipleValues(t *testing.T) {
	testTag, err := newTag(`key:"value"`)

	require.Nil(t, err)
	assert.Equal(t, "value", testTag.Value())
}

func Test_ParseInvalidTagsResultsInError(t *testing.T) {

	invalidTags := []string{
		`:"value"`,
		`key:`,
		`:`,
		` :"value"`,
		//		`key:"value,"`,
	}

	for _, tagInput := range invalidTags {
		_, err := newTag(tagInput)
		assert.NotNil(t, err, fmt.Sprintf("Invalid tag input: '%s'", tagInput))
	}
}
