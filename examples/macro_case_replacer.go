package main

import (
	"fmt"

	"gitlab.com/frumpled/factor3"
)

func main() {
	r := factor3.NewMacroCaseReplacer()

	for _, s := range []string{
		"apiKeys",
		"dataJSON",
		"UpperCamelCase",
		"Dot.Case",
	} {
		fmt.Print(s, "\t->\t")
		fmt.Println(r.Replace(s))
	}
}
