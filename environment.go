package factor3

type environment struct {
	variablePrefix string
}

// Initialize config for the environemnt to load
func LoadEnvironment() environment {
	return environment{}
}

// Set a prefix to use when fetchnig environment variables
func (e environment) WithVariablePrefix(environmentVariablePrefix string) environment {
	e.variablePrefix = environmentVariablePrefix

	return e
}

// Read environment into given struct
//  Traverses the passed in config object and populates it's fields w/ environment variable data
func (e environment) Into(configStruct interface{}) error {
	return readEnvironmentInto(
		e.variablePrefix,
		configStruct,
	)
}
