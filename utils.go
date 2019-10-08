package factor3

import (
	"regexp"
)

type varMatch struct {
	varName      string
	defaultValue string
}

func newMatchVar(input string) varMatch {
	matches := parseToMap(input)
	vm := varMatch{
		varName:      matches["name"],
		defaultValue: matches["default"],
	}

	return vm
}

func parseToMap(input string) map[string]string {
	r := regexp.MustCompile(`\${(?P<name>.*)(?::-)(?P<default>.*)}`)
	match := r.FindStringSubmatch(input)
	paramsMap := make(map[string]string)
	for i, name := range r.SubexpNames() {
		if i > 0 && i <= len(match) {
			paramsMap[name] = match[i]
		}
	}

	return paramsMap
}
