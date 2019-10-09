package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/frumpled/factor3"
)

func main() {
	os.Setenv("APP_EXAMPLE_DEFINED_VAR", "PASSED")

	conf := struct {
		UndefinedVar string `env:"${UNDEFINED_VAR:-Default value used}"`
		DefinedVar   string `env:"${DEFINED_VAR:-Default value used},required"`
		RequiredVar  string `env:"required"`
	}{}

	os.Setenv("DEFINED_VAR", "PASSED")

	factor3.
		LoadEnvironment().
		Debug().
		WithVariablePrefix("APP_EXAMPLE").
		Into(&conf)

	// Pretty print the conf variable:
	jsonString, _ := json.Marshal(&conf)
	fmt.Println(string(jsonString))
}
