# Factor 3

[![Build Status](https://travis-ci.org/frumpled/factor3.svg)](https://travis-ci.org/frumpled/factor3)
[![Go Report Card](https://goreportcard.com/badge/github.com/frumpled/factor3)](https://goreportcard.com/report/github.com/frumpled/factor3)
[![GoDoc](https://godoc.org/github.com/frumpled/factor3?status.svg)](https://godoc.org/github.com/frumpled/factor3)

## Overview
#### TLDR
This is an opinionated Golang package intended to load an environment into a config variable.  The intent is to strictly adhere to the principles outlined in Factor 3 of the [Twelve Factor App](https://12factor.net/) methodology.

#### Description
This is intended to ease the use of marshalling values from the environment into a struct for an app's config, similar to how the json/encoding package is doing thangs.  An intentionally small/restricted set of configuration flags are provided.  This package's functions should not need to be run frequently in an application's lifecycle, so ease-of-use has been considered higher in priorities than performance (within reason) here.

#### Example
```golang
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
		UndefinedVar string `envDefault:"Default value used"`
		DefinedVar   string `env:"required" envDefault:"Default value used"`
	}{}

	factor3.ReadEnvironmentInto("APP_EXAMPLE", &conf)

	// Pretty print the conf variable:
	jsonString, _ := json.Marshal(&conf)
	fmt.Println(string(jsonString))
}
```

## Notes
- All variable paths for the struct passed in are calculated with the given prefix and coverted to [MACRO_CASE](https://en.wikipedia.org/w/index.php?title=Naming_convention_(programming)#Delimiter-separated_words).
- Please not, this project is still in development (notice version 0).
	- [What this package will do / the problems it will solve](https://12factor.net/config) will remain the same
	- The interface for using / configuring this package may not be stable at this time.


This captures implementatino details and is strongly tied to the built in "reflect" package.
Testing this provides little value in comparison to the amount of effort required to do so and
relies on an already well tested/vetted package.  Further consideration is needed before any testing
will be pursued.

#### Potential future changes:
- Inject custom logger
- Clean up the codez.  I found dealing with reflection in Golang to be quite a nightmare; currently, the codebase "reflects" this D:  ( <-- dadjoke :D )

## Notes on test coverage

#### ./field.go

