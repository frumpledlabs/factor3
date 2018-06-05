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
