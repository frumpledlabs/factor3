package factor3

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
