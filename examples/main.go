package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/frumpled/factor3"
)

func main() {
	os.Setenv("APP_EXAMPLE_DEFINED_VAR", "SUP")

	conf := struct {
		UndefinedVar string `envDefault:"PASSED: Default value used"`
		DefinedVar   string `env:"required" envDefault:"PASSED: Default value used"`
		RequiredVar  string `env:"required"`
	}{}

	os.Setenv("APP_EXAMPLE_REQUIRED_VAR", "PASSED:  Required var loaded from env")

	factor3.
		LoadEnvironment().
		WithVariablePrefix("APP_EXAMPLE").
		Into(&conf)

	// Pretty print the conf variable:
	jsonString, _ := json.Marshal(&conf)
	fmt.Println(string(jsonString))
}
