package factor3

import (
	"regexp"
)

const varNameKey = "variableName"
const defaultValueKey = "defaultValue"

type envVarDefinition struct {
	varName      string
	defaultValue string
}

func newEnvVarDefinition(input string) envVarDefinition {
	matches := parseToMap(input)
	definition := envVarDefinition{
		varName:      matches[varNameKey],
		defaultValue: matches[defaultValueKey],
	}

	return definition
}

func parseToMap(input string) map[string]string {
	r := regexp.MustCompile(`^\${(?P<` + varNameKey + `>[A-z_-]*)(?::-)?(?P<` + defaultValueKey + `>.*)}`)
	match := r.FindStringSubmatch(input)
	paramsMap := make(map[string]string)
	for i, name := range r.SubexpNames() {
		if i > 0 && i <= len(match) {
			paramsMap[name] = match[i]
		}
	}

	return paramsMap
}
