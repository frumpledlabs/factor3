package main

// Loads configuration into the given variable, assuming all values are
// optional, unless otherwise tagged.
func Load(configuration interface{}) error {
	return nil
}

// Returns a map of field names and there calculated env variable name for
// convenience.
func PrintKeys(configuration interface{}) map[string]string {
	return make(map[string]string)
}
