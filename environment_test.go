package factor3

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_LoadVariablesPopulatesExpectedValues(t *testing.T) {
	expectedStringVal := "PASSED"
	os.Setenv("STRING_VAR", expectedStringVal)
	defer os.Unsetenv("STRING_VAR")

	expectedIntVal := int64(2)
	os.Setenv("INT_VAR", "2")
	defer os.Unsetenv("INT_VAR")

	expectedFloatVal := float64(16.32)
	os.Setenv("FLOAT_VAR", "16.32")
	defer os.Unsetenv("FLOAT_VAR")

	config := struct {
		NonExistingValue string
		StringVar        string
		FloatVar         float64
		IntVar           int64
	}{}

	err := LoadEnvironment().
		Into(&config)

	require.Nil(t, err)

	assert.Equal(t, "", config.NonExistingValue)
	assert.Equal(t, expectedStringVal, config.StringVar)
	assert.Equal(t, expectedFloatVal, config.FloatVar)
	assert.Equal(t, expectedIntVal, config.IntVar)
}

func Test_LoadVariablesWithPrefixPopluatesExpectedValues(t *testing.T) {
	expectedStringVal := "PASSED"
	os.Setenv("TEST_STRING_VAR", expectedStringVal)
	defer os.Unsetenv("TEST_STRING_VAR")

	config := struct {
		StringVar string
	}{}

	err := LoadEnvironment().
		WithVariablePrefix("TEST").
		Into(&config)

	require.Nil(t, err)

	assert.Equal(t, expectedStringVal, config.StringVar)
}
