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

func Test_LoadVariablesOfAllSupportedTypesWithoutError(t *testing.T) {
	os.Setenv("STRING_VAR", "STRING")
	os.Setenv("INT_VAR", "1")
	os.Setenv("INT8VAR", "8")
	os.Setenv("INT16VAR", "16")
	os.Setenv("INT32VAR", "32")
	os.Setenv("INT64VAR", "64")
	os.Setenv("FLOAT32VAR", "32.32")
	os.Setenv("FLOAT64VAR", "64.64")

	defer os.Unsetenv("STRING_VAR")
	defer os.Unsetenv("INT_VAR")
	defer os.Unsetenv("INT8VAR")
	defer os.Unsetenv("INT16VAR")
	defer os.Unsetenv("INT32VAR")
	defer os.Unsetenv("INT64VAR")
	defer os.Unsetenv("FLOAT32VAR")
	defer os.Unsetenv("FLOAT64VAR")

	config := struct {
		StringVar  string
		Float32Var float32
		Float64Var float64
		IntVar     int
		Int8Var    int8
		Int16Var   int16
		Int32Var   int32
		Int64Var   int64
	}{}

	err := LoadEnvironment().
		Into(&config)

	require.Nil(t, err)

	assert.Equal(t, "STRING", config.StringVar)
	assert.Equal(t, int(1), config.IntVar)
	assert.Equal(t, int8(8), config.Int8Var)
	assert.Equal(t, int16(16), config.Int16Var)
	assert.Equal(t, int32(32), config.Int32Var)
	assert.Equal(t, int64(64), config.Int64Var)
	assert.Equal(t, float32(32.32), config.Float32Var)
	assert.Equal(t, float64(64.64), config.Float64Var)
}
