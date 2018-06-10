package factor3

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_SingleRootLevelVariableIsLoaded(t *testing.T) {
	os.Setenv("TEST", "PASS")
	defer os.Unsetenv("PASS")

	conf := struct {
		Test string
	}{}

	err := Load(&conf)
	assert.Nil(t, err)
	assert.Equal(t, "PASS", conf.Test)
}

func Test_NestedVariableIsLoaded(t *testing.T) {
	os.Setenv("ANOTHER_TEST", "PASS")
	defer os.Unsetenv("ANOTHER_TEST")

	conf := struct {
		Another struct {
			Test string
		}
	}{}

	err := Load(&conf)
	assert.Nil(t, err)
	assert.Equal(t, "PASS", conf.Another.Test)
}

func Test_UnsetRequiredVariableErrors(t *testing.T) {
	conf := struct {
		UnsetRequiredVar string `envRequired:"true"`
	}{}

	err := Load(&conf)
	assert.NotNil(t, err)
}

func Test_SetRequiredVariableLoadsWithoutError(t *testing.T) {
	os.Setenv("REQUIRED_VAR", "PASS")
	defer os.Unsetenv("REQUIRED_VAR")

	conf := struct {
		RequiredVar string `envRequired:"true"`
	}{}

	err := Load(&conf)
	assert.Nil(t, err)
}

func Test_DefaultValueOverridden(t *testing.T) {
	os.Setenv("DEFAULT_KEY_EXISTS", "OVERRIDDEN")
	defer os.Unsetenv("DEFAULT_KEY_EXISTS")

	conf := struct {
		DefaultKeyExists string `envDefault:"DEFAULT"`
	}{}

	err := Load(&conf)

	assert.Nil(t, err)
	assert.Equal(t, "OVERRIDDEN", conf.DefaultKeyExists)
}

func Test_DefaultValueIsOveriddenWhenEmptyValueSet(t *testing.T) {
	os.Setenv("DEFAULT_KEY_IS_EMPTY_STRING", "")
	defer os.Unsetenv("DEFAULT_KEY_IS_EMPTY_STRING")

	conf := struct {
		defaultKeyIsEmptyString string `envDefault:"EMPTY"`
	}{}

	Load(&conf)

	assert.Equal(t, "", conf.defaultKeyIsEmptyString)
}

func Test_DefaultValuePersistsWhenEnvVariableNotSet(t *testing.T) {
	conf := struct {
		DefaultKeySet string `envDefault:"DEFAULT"`
	}{}

	Load(&conf)

	assert.Equal(t, "DEFAULT", conf.DefaultKeySet)
}

func Test_RequiredWithDefaultDoesNotErrorWhenNotSet(t *testing.T) {
	conf := struct {
		RequiredWithDefault string `env:"required" envDefault:"DEFAULT"`
	}{}

	err := Load(&conf)

	assert.Nil(t, err)
}

func Test_Keys(t *testing.T) {
	conf := struct {
		A      string
		Nested struct {
			C string
			D string
		}
	}{}

	keys, err := Keys(conf)

	require.Nil(t, err)
	assert.Len(t, keys, 3)
}
