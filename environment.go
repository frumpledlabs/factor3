package factor3

import (
	"github.com/frumpled/factor3/logger"
)

// Environment accumulates meta-data used to populate input with values when loading
type Environment struct {
	variablePrefix string
	fields         map[string]fieldInfo
}

// LoadEnvironment initializes the environemnt w/ default configuration
func LoadEnvironment() Environment {
	return Environment{}
}

// WithVariablePrefix sets a prefix used when defining default environment variables
func (e Environment) WithVariablePrefix(environmentVariablePrefix string) Environment {
	e.variablePrefix = environmentVariablePrefix
	log.Info("Using environment variable prefix.",
		map[string]interface{}{
			"prefix": environmentVariablePrefix,
		})

	return e
}

// Into reads local environment into this "environment" instance
//  Traverses the passed in config object and populates it's fields w/ environment variable data
func (e Environment) Into(input interface{}) error {

	e.readFields(input)
	// setFields(input, e.fields)

	return readEnvironmentInto(
		e.variablePrefix,
		input,
	)
}

// Debug reads the struct and logs the fields it finds with the sources for values it will use for each field.
func (e Environment) Debug(input interface{}) {
	log.WithLevel(logger.DebugLevel)

	e.readFields(input)

	for key, value := range e.fields {
		log.Debug(
			"",
			map[string]interface{}{
				"key":   key,
				"value": value,
			},
		)
	}
}

func (e Environment) readFields(input interface{}) {
	fields, err := readEnvironmentFor(
		e.variablePrefix,
		input,
	)

	e.fields = fields

	if err != nil {
		log.Fatal(
			"Error reading fields",
			map[string]interface{}{
				"msg": err.Error(),
			},
		)
	}
}
