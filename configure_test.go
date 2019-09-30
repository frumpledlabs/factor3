package factor3

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SingleRootLevelVariableIsReadEnvironmentIntoed(t *testing.T) {
	os.Setenv("TEST", "PASS")
	defer os.Unsetenv("PASS")

	conf := struct {
		Test string
	}{}

	err := LoadEnvironment().Into(&conf)
	assert.Nil(t, err)
	assert.Equal(t, "PASS", conf.Test)
}

func Test_NestedVariableIsReadEnvironmentIntoed(t *testing.T) {
	os.Setenv("ANOTHER_TEST", "PASS")
	defer os.Unsetenv("ANOTHER_TEST")

	conf := struct {
		Another struct {
			Test string
		}
	}{}

	err := LoadEnvironment().Into(&conf)

	assert.Nil(t, err)
	assert.Equal(t, "PASS", conf.Another.Test)
}

func Test_UnsetRequiredVariableErrors(t *testing.T) {
	conf := struct {
		UnsetRequiredVar string `envRequired:"true"`
	}{}

	err := LoadEnvironment().Into(&conf)
	assert.NotNil(t, err)
}

func Test_SetRequiredVariableReadEnvironmentIntoWithoutError(t *testing.T) {
	os.Setenv("REQUIRED_VAR", "PASS")
	defer os.Unsetenv("REQUIRED_VAR")

	conf := struct {
		RequiredVar string `envRequired:"true"`
	}{}

	err := LoadEnvironment().Into(&conf)
	assert.Nil(t, err)
}

func Test_DefaultValueOverridden(t *testing.T) {
	os.Setenv("DEFAULT_KEY_EXISTS", "OVERRIDDEN")
	defer os.Unsetenv("DEFAULT_KEY_EXISTS")

	conf := struct {
		Default struct {
			KeyExists string `envDefault:"DEFAULT"`
		}
	}{}

	err := LoadEnvironment().Into(&conf)

	assert.Nil(t, err)
	assert.Equal(t, "OVERRIDDEN", conf.Default.KeyExists)
}

func Test_DefaultValueIsOveriddenWhenEmptyValueSet(t *testing.T) {
	os.Setenv("DEFAULT_KEY_IS_EMPTY_STRING", "")
	defer os.Unsetenv("DEFAULT_KEY_IS_EMPTY_STRING")

	conf := struct {
		defaultKeyIsEmptyString string `envDefault:"EMPTY"`
	}{}

	LoadEnvironment().Into(&conf)

	assert.Equal(t, "", conf.defaultKeyIsEmptyString)
}

func Test_DefaultValuePersistsWhenEnvVariableNotSet(t *testing.T) {
	conf := struct {
		DefaultKeySet string `envDefault:"DEFAULT"`
		DefaultBool   bool   `envDefault:"true"`
	}{}

	LoadEnvironment().Into(&conf)

	assert.Equal(t, "DEFAULT", conf.DefaultKeySet)
	assert.Equal(t, true, conf.DefaultBool)
}

func Test_RequiredWithDefaultDoesNotErrorWhenNotSet(t *testing.T) {
	conf := struct {
		RequiredWithDefault string `env:"required" envDefault:"DEFAULT"`
	}{}

	err := LoadEnvironment().Into(&conf)

	assert.Nil(t, err)
}

func ExampleLoadEnvironment() {
	os.Setenv("STRING", "String value")
	os.Setenv("BOOL", "true")
	os.Setenv("INT", "42")
	os.Setenv("INT64", "64")
	os.Setenv("NESTED_STRING", "nestedString")
	os.Setenv("NESTED_BOOL", "true")
	os.Setenv("NESTED_INT", "42")
	os.Setenv("NESTED_INT64", "64")
	os.Setenv("DEFAULTED_VALUES_OVERIDDEN_STRING", "WAS_OVERRIDEN")

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
	// println(conf)

	// Import "encoding/json" to pretty print:
	// jsonString, _ := json.Marshal(&conf)
	// log.Info(string(jsonString))
}

func TestLoadFieldWithoutRequiredValueFails(t *testing.T) {
	conf := struct {
		RequiredValue string `envOpts:"required"`
	}{}

	err := LoadEnvironment().Into(&conf)
	assert.NotNil(t, err)
}

func TestLoadRequiredFieldWithValueSucceeds(t *testing.T) {
	expected := "PASSED"
	os.Setenv("REQUIRED_VALUE", expected)

	conf := struct {
		RequiredValue string `envOpts:"required"`
	}{}

	err := LoadEnvironment().Into(&conf)
	assert.Nil(t, err)

	assert.Equal(t, expected, conf.RequiredValue)
}

func TestLoadingFieldWithOverrideNameLoads(t *testing.T) {
	expected := "PASSED"
	os.Setenv("SomeOtherFieldName", expected)

	conf := struct {
		Field string `env:"SomeOtherFieldName"`
	}{}

	err := LoadEnvironment().Into(&conf)
	assert.Nil(t, err)

	assert.Equal(t, expected, conf.Field)
}
