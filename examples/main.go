package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/frumpled/factor3"
)

func main() {
	os.Setenv("APP_EXAMPLE_DEFINED_VAR", "PASSED")
	os.Setenv("DEFINED_VAR", "PASSED")
	os.Setenv("APP_EXAMPLE_REQUIRED_VAR", "SET")

	conf := struct {
		UndefinedVar string `env:"${UNDEFINED_VAR:-Default value used}"`
		DefinedVar   string `env:"${DEFINED_VAR:-Default value used},required"`
		RequiredVar  string `env:"required"`
	}{}

	err := factor3.
		LoadEnvironment().
		Debug().
		WithVariablePrefix("APP_EXAMPLE").
		Into(&conf)

	if err != nil {
		panic("Unexpected error")
	}

	// Pretty print the conf variable:
	jsonString, _ := json.Marshal(&conf)
	fmt.Println(string(jsonString))
}
