package factor3

import (
	"github.com/frumpled/factor3/logger"
)

// Environment accumulates meta-data used to populate field data into an input from the environment
type Environment struct {
	variablePrefix string
}

// LoadEnvironment initializes the environemnt w/ default configuration
func LoadEnvironment() Environment {
	return Environment{}
}

// WithVariablePrefix sets a prefix to determine default environment variable names
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
func (e Environment) Into(configStruct interface{}) error {
	return setFields(
		e.variablePrefix,
		configStruct,
	)
}

// Debug enables debug level logging
func (e Environment) Debug() Environment {
	log.WithLevel(logger.DebugLevel)

	return e
}
