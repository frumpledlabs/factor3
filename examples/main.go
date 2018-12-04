package main

import (
	"fmt"
	"os"

	"github.com/frumpled/factor3"
)

func main() {
	fmt.Println("vim-go")

	os.Setenv("EXISTING_VARIABLE_VALUE", "PASSED")
	defer os.Unsetenv("EXISTING_VARIABLE_VALUE")

	config := struct {
		ExistingVariableValue    string
		NonExistingVariableValue string
	}{
		ExistingVariableValue:    "",
		NonExistingVariableValue: "",
	}

	err := factor3.LoadEnvironment().
		WithVariablePrefix("").
		Into(&config)

	fmt.Println("Err:", err)
	fmt.Println("Config:", config)

}
