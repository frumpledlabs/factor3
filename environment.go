package factor3

import (
	"github.com/frumpled/factor3/logger"
)

type environment struct {
	variablePrefix string
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
func (e environment) Into(configStruct interface{}) error {
	return readEnvironmentInto(
		e.variablePrefix,
		configStruct,
	)
}

// Debug reads the struct and logs the fields it finds with the sources for values it will use for each field.
func (e environment) Debug(configStruct interface{}) {
	log.WithLevel(logger.DebugLevel)

	output, err := parseEnvironmentToMap(
		e.variablePrefix,
		configStruct,
	)

	if err != nil {
		log.Fatal(
			"Error debugging struct",
			map[string]interface{}{
				"msg": err,
			},
		)
	}

	for key, value := range output {
		log.Debug(
			"",
			map[string]interface{}{
				"key":   key,
				"value": value,
			},
		)
	}
}
