# Factor 3

[![Build Status](https://travis-ci.org/frumpled/factor3.svg)](https://travis-ci.org/frumpled/factor3)
[![Go Report Card](https://goreportcard.com/badge/github.com/frumpled/factor3)](https://goreportcard.com/report/github.com/frumpled/factor3)
[![GoDoc](https://godoc.org/github.com/frumpled/factor3?status.svg)](https://godoc.org/github.com/frumpled/factor3)

## Overview
#### TLDR
Opinionated environment config loading for Golang.  Intended to strictly follow Factor 3 of the [Twelve Factor App](https://12factor.net/) methodology.

#### Description
This is intended to ease the use of marshalling values from the environment into a struct for an app's config, similar to how the json/encoding package is doing thangs.  An intentionally small/restricted set of configuration flags are provided.  This package's functions should not need to be run frequently in an application's lifecycle, so ease-of-use has been considered higher in priorities than performance (within reason) here.

#### Example
```go
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

	jsonString, _ := json.Marshal(&conf)
	fmt.Println(string(jsonString))
}
```

## Notes
- All variable paths are calculated with given prefix and coverted to [MACRO_CASE](https://en.wikipedia.org/w/index.php?title=Naming_convention_(programming)#Delimiter-separated_words).
- This project is still in development (notice version 0).
- What will remain stable:
  - [The problems this package will solve](https://12factor.net/config) will remain the same

- What may note remain stable:
  - The interface for using / configuring this package.

#### Potential future changes:
- Inject custom logger
- More complete Godocs
- Clean up the codez.  I found dealing with reflection in Golang to be quite a nightmare; currently, the codebase "reflects" this D:  ( <-- dadjoke :D )