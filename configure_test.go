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

	err := ReadEnvironmentInto(&conf)
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

	err := ReadEnvironmentInto(&conf)
	assert.Nil(t, err)
	assert.Equal(t, "PASS", conf.Another.Test)
}

func Test_UnsetRequiredVariableErrors(t *testing.T) {
	conf := struct {
		UnsetRequiredVar string `envRequired:"true"`
	}{}

	err := ReadEnvironmentInto(&conf)
	assert.NotNil(t, err)
}

func Test_SetRequiredVariableReadEnvironmentIntosWithoutError(t *testing.T) {
	os.Setenv("REQUIRED_VAR", "PASS")
	defer os.Unsetenv("REQUIRED_VAR")

	conf := struct {
		RequiredVar string `envRequired:"true"`
	}{}

	err := ReadEnvironmentInto(&conf)
	assert.Nil(t, err)
}

func Test_DefaultValueOverridden(t *testing.T) {
	os.Setenv("DEFAULT_KEY_EXISTS", "OVERRIDDEN")
	defer os.Unsetenv("DEFAULT_KEY_EXISTS")

	conf := struct {
		DefaultKeyExists string `envDefault:"DEFAULT"`
	}{}

	err := ReadEnvironmentInto(&conf)

	assert.Nil(t, err)
	assert.Equal(t, "OVERRIDDEN", conf.DefaultKeyExists)
}

func Test_DefaultValueIsOveriddenWhenEmptyValueSet(t *testing.T) {
	os.Setenv("DEFAULT_KEY_IS_EMPTY_STRING", "")
	defer os.Unsetenv("DEFAULT_KEY_IS_EMPTY_STRING")

	conf := struct {
		defaultKeyIsEmptyString string `envDefault:"EMPTY"`
	}{}

	ReadEnvironmentInto(&conf)

	assert.Equal(t, "", conf.defaultKeyIsEmptyString)
}

func Test_DefaultValuePersistsWhenEnvVariableNotSet(t *testing.T) {
	conf := struct {
		DefaultKeySet string `envDefault:"DEFAULT"`
		DefaultBool   bool   `envDefault:"true"`
	}{}

	ReadEnvironmentInto(&conf)

	assert.Equal(t, "DEFAULT", conf.DefaultKeySet)
	assert.Equal(t, true, conf.DefaultBool)
}

func Test_RequiredWithDefaultDoesNotErrorWhenNotSet(t *testing.T) {
	conf := struct {
		RequiredWithDefault string `env:"required" envDefault:"DEFAULT"`
	}{}

	err := ReadEnvironmentInto(&conf)

	assert.Nil(t, err)
}
