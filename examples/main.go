package main

import (
	"fmt"
	"os"
)

func main() {
	os.Setenv("STRING", "String")
	os.Setenv("BOOL", "true")
	os.Setenv("INT", "42")
	os.Setenv("INT64", "64")
	os.Setenv("NESTED_STRING", "nestedString")
	os.Setenv("NESTED_BOOL", "false")
	os.Setenv("NESTED_INT", "42")
	os.Setenv("NESTED_INT64", "64")

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
	}{}

	fmt.Println(conf)
	Load(&conf)
	fmt.Println(conf)
}
