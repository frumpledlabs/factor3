# Factor 3

[![Build Status](https://travis-ci.org/frumpled/factor3.svg)](https://travis-ci.org/frumpled/factor3)
[![Go Report Card](https://goreportcard.com/badge/github.com/frumpled/factor3)](https://goreportcard.com/report/github.com/frumpled/factor3)
[![GoDoc](https://godoc.org/github.com/frumpled/factor3?status.svg)](https://godoc.org/github.com/frumpled/factor3)

## Overview
### TLDR

[What this project intends to solve/make simpler](https://12factor.net/config) stems from the configuration portion of 12 Factor Apps.

### Example
```golang
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
```

## Description

This is an opinionated Golang package intended to load an environment into a config variable.  The intent is to strictly adhere to the principles outlined in Factor 3 of the [Twelve Factor App](https://12factor.net/) methodology.


This is intended to ease the use of marshalling values from the environment into a struct for an app's config, similar to how the json/encoding package is doing thangs.  An intentionally small/restricted set of configuration flags are provided.  This package's functions should not need to be run frequently in an application's lifecycle, so ease-of-use has been considered higher in priorities than performance (within reason) here.


## Notes
- All variable paths for the struct passed in are calculated with the given prefix and coverted to [MACRO_CASE](https://en.wikipedia.org/w/index.php?title=Naming_convention_(programming)#Delimiter-separated_words).
- Please note that this project is still in development (notice version 0).
- This captures implementatino details and is strongly tied to the built in "reflect" package.