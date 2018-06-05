package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SingleRootLevelVariableIsLoaded(t *testing.T) {
	os.Setenv("TEST", "PASS")
	defer os.Unsetenv("TEST")

	conf := struct {
		test string
	}{}

	Load(conf)

	assert.Equal(t, "PASS", conf.test)
}

func Test_NestedVariableIsLoaded(t *testing.T) {
	os.Setenv("ANOTHER_TEST", "PASS")
	defer os.Unsetenv("ANOTHER_TEST")

	conf := struct {
		another struct {
			test string
		}
	}{}

	Load(conf)

	assert.Equal(t, "PASS", conf.another.test)
}

func Test_UnsetRequiredVariableErrors(t *testing.T) {
	conf := struct {
		test string `env:"required"`
	}{}

	err := Load(conf)
	assert.NotNil(t, err)
}

func Test_SetRequiredVariableLoadsWithoutError(t *testing.T) {
	os.Setenv("REQUIRED_VAR", "PASS")
	defer os.Unsetenv("REQUIRED_VAR")

	conf := struct {
		requiredVar string `env:"required"`
	}{}

	err := Load(conf)
	assert.Nil(t, err)
}

func Test_DefaultValueOverriden(t *testing.T) {
	os.Setenv("DEFAULT_KEY_EXISTS", "OVERRIDDEN")
	os.Unsetenv("DEFAULT_KEY_EXISTS")

	conf := struct {
		defaultKeyExists string `envDefault:"MISSING"`
	}{}

	Load(conf)

	assert.Equal(t, "OVERRIDDEN", conf.defaultKeyExists)
}

func Test_DefaultValueIsOveriddenWhenEmptyValueSet(t *testing.T) {
	os.Setenv("DEFAULT_KEY_IS_EMPTY_STRING", "")
	defer os.Unsetenv("DEFAULT_KEY_IS_EMPTY_STRING")

	conf := struct {
		defaultKeyIsEmptyString string `envDefault:"EMPTY"`
	}{}

	Load(conf)

	assert.Equal(t, "", conf.defaultKeyIsEmptyString)
}

func Test_DefaultValuePersistsWhenEnvVariableNotSet(t *testing.T) {
	conf := struct {
		defaultKeySet string `envDefault:"DEFAULT"`
	}{}

	Load(conf)

	assert.Equal(t, "DEFAULT", conf.defaultKeySet)
}

func Test_RequiredWithDefaultDoesNotErrorWhenNotSet(t *testing.T) {
	conf := struct {
		requiredWithDefault string `env:"required" envDefault:"DEFAULT"`
	}{}

	err := Load(conf)

	assert.Nil(t, err)
}

func Test_PrintKeys(t *testing.T) {
	assert.Fail(t, "Not yet implemented")
}
