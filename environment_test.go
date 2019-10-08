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
	os.Setenv("INT8_VAR", "8")
	os.Setenv("INT16_VAR", "16")
	os.Setenv("INT32_VAR", "32")
	os.Setenv("INT64_VAR", "64")
	os.Setenv("FLOAT32_VAR", "32.32")
	os.Setenv("FLOAT64_VAR", "64.64")
	os.Setenv("BOOL_VAR", "true")

	defer os.Unsetenv("STRING_VAR")
	defer os.Unsetenv("INT_VAR")
	defer os.Unsetenv("INT8_VAR")
	defer os.Unsetenv("INT16_VAR")
	defer os.Unsetenv("INT32_VAR")
	defer os.Unsetenv("INT64_VAR")
	defer os.Unsetenv("FLOAT32_VAR")
	defer os.Unsetenv("FLOAT64_VAR")
	defer os.Unsetenv("BOOL_VAR")

	config := struct {
		StringVar  string
		Float32Var float32
		Float64Var float64
		IntVar     int
		Int8Var    int8
		Int16Var   int16
		Int32Var   int32
		Int64Var   int64
		BoolVar    bool
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
	assert.Equal(t, true, config.BoolVar)
}

func Test_GivenInvalidInput_ThenErrorIsReturnedWhenLoadingVariables(t *testing.T) {
	os.Setenv("BOOL_VAR", "trueeeee")
	os.Setenv("FLOAT32_VAR", "thirty-two")
	os.Setenv("FLOAT64_VAR", "sixty-four")
	os.Setenv("INT_VAR", "1.1")
	os.Setenv("INT8_VAR", "8.1")
	os.Setenv("INT16_VAR", "16.1")
	os.Setenv("INT32_VAR", "32.1")
	os.Setenv("INT64_VAR", "64.1")

	defer os.Unsetenv("BOOL_VAR")
	defer os.Unsetenv("FLOAT32_VAR")
	defer os.Unsetenv("FLOAT64_VAR")
	defer os.Unsetenv("INT_VAR")
	defer os.Unsetenv("INT8_VAR")
	defer os.Unsetenv("INT16_VAR")
	defer os.Unsetenv("INT32_VAR")
	defer os.Unsetenv("INT64_VAR")

	var err error

	err = LoadEnvironment().
		Into(&struct {
			BoolVar bool
		}{})
	assert.NotNil(t, err)

	err = LoadEnvironment().
		Into(&struct {
			Float32Var float32
		}{})
	assert.NotNil(t, err)

	err = LoadEnvironment().
		Into(&struct {
			Float64Var float64
		}{})
	assert.NotNil(t, err)

	err = LoadEnvironment().
		Into(&struct {
			IntVar int
		}{})
	assert.NotNil(t, err)

	err = LoadEnvironment().
		Into(&struct {
			Int8Var int8
		}{})
	assert.NotNil(t, err)

	err = LoadEnvironment().
		Into(&struct {
			Int16Var int16
		}{})
	assert.NotNil(t, err)

	err = LoadEnvironment().
		Into(&struct {
			Int32Var int32
		}{})
	assert.NotNil(t, err)

	err = LoadEnvironment().
		Into(&struct {
			Int64Var int64
		}{})
	assert.NotNil(t, err)

	err = LoadEnvironment().
		Into(&struct {
			BoolVar bool
		}{})
	assert.NotNil(t, err)
}

func Test_EndToEnd(t *testing.T) {
	os.Setenv("APP_EXAMPLE_DEFINED_VAR", "PASSED")

	conf := struct {
		UndefinedVar string `env:"${UNDEFINED_VAR:-Default value used}"`
		// DefinedVar   string `env:"${DEFINED_VAR:-Default value used},required"`
		// RequiredVar  string `env:"required"`
	}{}

	LoadEnvironment().
		Debug().
		WithVariablePrefix("APP_EXAMPLE").
		Into(&conf)
}
