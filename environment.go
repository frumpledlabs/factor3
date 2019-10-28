package factor3

import (
	"github.com/frumpled/factor3/logger"
)

type environment struct {
	variablePrefix string
	fields         map[string]fieldInfo
}

// LoadEnvironment initializes the environemnt w/ default configuration
func LoadEnvironment() environment {
	return environment{}
}

// Set a prefix to use when fetchnig environment variables
func (e environment) WithVariablePrefix(environmentVariablePrefix string) environment {
	e.variablePrefix = environmentVariablePrefix
	log.Info("Using environment variable prefix.",
		map[string]interface{}{
			"prefix": environmentVariablePrefix,
		})

	return e
}

// Into reads local environment into this "environment" instance
//  Traverses the passed in config object and populates it's fields w/ environment variable data
func (e environment) Into(input interface{}) error {

	e.readFields(input)
	// setFields(input, e.fields)

	return readEnvironmentInto(
		e.variablePrefix,
		input,
	)
}

// Debug reads the struct and logs the fields it finds with the sources for values it will use for each field.
func (e environment) Debug(input interface{}) {
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

func (e environment) readFields(input interface{}) {
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
