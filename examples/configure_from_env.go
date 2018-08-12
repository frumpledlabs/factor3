package main

import (
	"fmt"
	"os"

	"gitlab.com/frumpled/factor3"
)

func main() {
	os.Setenv("STRING", "String value")
	os.Setenv("BOOL", "true")
	os.Setenv("INT", "42")
	os.Setenv("INT64", "64")
	os.Setenv("NESTED_STRING", "nestedString")
	os.Setenv("NESTED_BOOL", "true")
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
		DefaultedValues struct {
			DefaultFalse  bool   `envDefault:"false"`
			DefaultTrue   bool   `envDefualt:"true"`
			DefaultString string `envDefault:"default string value"`
		}
	}{}

	fmt.Println(conf)
	factor3.ReadEnvironmentInto(&conf)
	fmt.Println(conf)
}
