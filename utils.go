package factor3

import (
	"regexp"
)

const varNameKey = ""
const defaultValueKey = "defaultValue"

type varMatch struct {
	varName      string
	defaultValue string
}

func newMatchVar(input string) varMatch {
	matches := parseToMap(input)
	vm := varMatch{
		varName:      matches[varNameKey],
		defaultValue: matches[defaultValueKey],
	}

	return vm
}

func parseToMap(input string) map[string]string {
	r := regexp.MustCompile(`\${(?P<` + varNameKey + `>.*)(?::-)(?P<` + defaultValueKey + `>.*)}`)
	match := r.FindStringSubmatch(input)
	paramsMap := make(map[string]string)
	for i, name := range r.SubexpNames() {
		if i > 0 && i <= len(match) {
			paramsMap[name] = match[i]
		}
	}

	return paramsMap
}