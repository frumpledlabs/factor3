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
	type testCase struct {
		envVarName  string
		envVarValue string
		testStruct  *interface{}
	}

	testCases := []struct {
		envVarName  string
		envVarValue string
		testStruct  interface{}
	}{
		{
			envVarName:  "BOOL_VAR",
			envVarValue: "not a bool",
			testStruct: &struct {
				BoolVar bool
			}{},
		},
		{
			envVarName:  "FLOAT32_VAR",
			envVarValue: "not-float-32",
			testStruct: &struct {
				Float32Var float32
			}{},
		},
		{
			envVarName:  "FLOAT64_VAR",
			envVarValue: "sixty-four",
			testStruct: &struct {
				Float64Var float64
			}{},
		},
		{
			envVarName:  "INT_VAR",
			envVarValue: "43.2",
			testStruct: &struct {
				IntVar int
			}{},
		},
		{
			envVarName:  "INT8_VAR",
			envVarValue: "543.2",
			testStruct: &struct {
				Int8Var int8
			}{},
		},
		{
			envVarName:  "INT16_VAR",
			envVarValue: "32.1",
			testStruct: &struct {
				Int16Var int16
			}{},
		},
		{
			envVarName:  "INT32_VAR",
			envVarValue: "32.1",
			testStruct: &struct {
				Int32Var int32
			}{},
		},
		{
			envVarName:  "INT64_VAR",
			envVarValue: "64.1",
			testStruct: &struct {
				Int64Var int64
			}{},
		},
	}

	for _, tc := range testCases {
		os.Setenv(tc.envVarName, tc.envVarValue)
		defer os.Unsetenv(tc.envVarName)

		err := LoadEnvironment().
			Into(tc.testStruct)
		assert.NotNil(t, err)
	}
}

func Test_EndToEnd(t *testing.T) {
	os.Setenv("APP_EXAMPLE_DEFINED_VAR", "PASSED")

	conf := struct {
		UndefinedVar string `env:"${UNDEFINED_VAR:-Default value used}"`
		DefinedVar   string `env:"${DEFINED_VAR:-Default value used},required"`
		// RequiredVar  string `env:"required"`
	}{}

	LoadEnvironment().
		WithVariablePrefix("APP_EXAMPLE").
		Into(&conf)
}

func Test_SingleRootLevelVariableIsReadEnvironment(t *testing.T) {
	os.Setenv("TEST", "PASS")
	defer os.Unsetenv("PASS")

	conf := struct {
		Test string
	}{}

	err := LoadEnvironment().Into(&conf)
	assert.Nil(t, err)
	assert.Equal(t, "PASS", conf.Test)
}

func Test_NestedVariableIsRead(t *testing.T) {
	os.Setenv("ANOTHER_DEEPLY_NESTED_FIELD", "PASS")
	defer os.Unsetenv("ANOTHER_DEEPLY_NESTED_FIELD")

	conf := struct {
		Another struct {
			Deeply struct {
				Nested struct {
					Field string
				}
			}
		}
	}{}

	err := LoadEnvironment().Into(&conf)

	assert.Nil(t, err)
	assert.Equal(t, "PASS", conf.Another.Deeply.Nested.Field)
}

func Test_UnsetRequiredVariableErrors(t *testing.T) {
	conf := struct {
		UnsetRequiredVar string `env:"required"`
	}{}

	err := LoadEnvironment().Into(&conf)
	assert.NotNil(t, err)
}

func Test_SetRequiredVariableReadEnvironmentIntoWithoutError(t *testing.T) {
	os.Setenv("REQUIRED_VAR", "PASS")
	defer os.Unsetenv("REQUIRED_VAR")

	conf := struct {
		RequiredVar string `env:"required"`
	}{}

	err := LoadEnvironment().Into(&conf)
	assert.Nil(t, err)
}

func Test_DefaultValueOverridden(t *testing.T) {
	os.Setenv("DEFAULT_KEY_EXISTS", "OVERRIDDEN")
	defer os.Unsetenv("DEFAULT_KEY_EXISTS")

	conf := struct {
		Default struct {
			KeyExists string `env:"${:-DEFAULT}"`
		}
	}{}

	err := LoadEnvironment().Into(&conf)

	assert.Nil(t, err)
	assert.Equal(t, "OVERRIDDEN", conf.Default.KeyExists)
}

func Test_DefaultValuePersistsWhenEnvVariableNotSet(t *testing.T) {
	conf := struct {
		DefaultKeySet string `env:"${:-DEFAULT}"`
		DefaultBool   bool   `env:"${:-true}"`
	}{}

	LoadEnvironment().Into(&conf)

	assert.Equal(t, "DEFAULT", conf.DefaultKeySet)
	assert.Equal(t, true, conf.DefaultBool)
}

func Test_RequiredWithDefaultDoesNotErrorWhenNotSet(t *testing.T) {
	conf := struct {
		RequiredWithDefault string `env:"${:-DEFAULT},required"`
	}{}

	err := LoadEnvironment().Into(&conf)

	assert.Nil(t, err)
}

func ExampleLoadEnvironment() {
	os.Setenv("BOOL", "true")
	os.Setenv("DEFAULTED_VALUES_OVERIDDEN_STRING", "WAS_OVERRIDEN")
	os.Setenv("INT", "42")
	os.Setenv("INT64", "64")
	os.Setenv("NESTED_BOOL", "true")
	os.Setenv("NESTED_INT", "42")
	os.Setenv("NESTED_INT64", "64")
	os.Setenv("NESTED_STRING", "nestedString")
	os.Setenv("STRING", "String value")

	defer os.Unsetenv("BOOL")
	defer os.Unsetenv("DEFAULTED_VALUES_OVERIDDEN_STRING")
	defer os.Unsetenv("INT")
	defer os.Unsetenv("INT64")
	defer os.Unsetenv("NESTED_BOOL")
	defer os.Unsetenv("NESTED_INT")
	defer os.Unsetenv("NESTED_INT64")
	defer os.Unsetenv("NESTED_STRING")
	defer os.Unsetenv("STRING")

	conf := struct {
		String string
		Bool   bool
		Int    int
		Int64  int64
		Nested struct {
			String string
			Bool   bool
			Int    int
			Int64  int64
		}
		DefaultedValues struct {
			DefaultFalse    bool   `envDefault:"false"`
			DefaultTrue     bool   `envDefault:"true"`
			DefaultString   string `envDefault:"default string value"`
			OveriddenString string `envDefault:"NOT_OVERRIDEN"`
		}
	}{}

	LoadEnvironment().Into(&conf)
}

func TestLoadFieldWithoutRequiredValueFails(t *testing.T) {
	conf := struct {
		RequiredValue string `envOpts:"required"`
	}{}

	err := LoadEnvironment().Into(&conf)
	assert.Nil(t, err)
}

func Test_LoadRequiredFieldWithValueSucceeds(t *testing.T) {
	expected := "PASSED"
	os.Setenv("REQUIRED_VALUE", expected)
	defer os.Unsetenv("REQUIRED_VALUE")

	conf := struct {
		RequiredValue string `env:"required"`
	}{}

	err := LoadEnvironment().Into(&conf)
	assert.Nil(t, err)

	assert.Equal(t, expected, conf.RequiredValue)
}

func Test_LoadRequiredFieldWithoutValueFails(t *testing.T) {
	conf := struct {
		RequiredValue string `env:"required"`
	}{}

	err := LoadEnvironment().Into(&conf)
	assert.NotNil(t, err)
}

func Test_LoadRequiredFieldWithDefaultValueSucceeds(t *testing.T) {
	conf := struct {
		RequiredValue string `env:"${:-DEFAULT_VALUE},required"`
	}{}

	err := LoadEnvironment().Into(&conf)
	assert.Nil(t, err)
	assert.Equal(t, "DEFAULT_VALUE", conf.RequiredValue)
}

func TestLoadingFieldWithOverrideNameLoads(t *testing.T) {
	expected := "PASSED"
	os.Setenv("SomeOtherFieldName", expected)
	defer os.Unsetenv("SomeOtherFieldName")

	conf := struct {
		Field string `env:"${SomeOtherFieldName}"`
	}{}

	err := LoadEnvironment().Into(&conf)
	assert.Nil(t, err)

	assert.Equal(t, expected, conf.Field)
}

func TestLoadFieldWithMultipleConfigLoads(t *testing.T) {
	expected := "PASSED"
	os.Setenv("SomeOtherFieldName", expected)

	conf := struct {
		Field string `env:"${SomeOtherFieldName},required"`
	}{}

	err := LoadEnvironment().Into(&conf)
	assert.Nil(t, err)
	assert.Equal(t, expected, conf.Field)
}
