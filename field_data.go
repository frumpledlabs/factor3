package factor3

import (
	"regexp"
	"strings"
)

const labelIsRequired = "required"

var envVariablePattern = regexp.MustCompile(`\${.*(?::-)?.*}`)

type fieldData struct {
	definition      string
	defaultValue    string
	overrideKey     string
	isRequired      bool
	keyIsOverriden  bool
	hasDefaultValue bool
}

func newFieldData(input string) fieldData {
	fd := fieldData{
		definition: input,
	}

	tags := strings.Split(input, ",")
	if len(tags) == 0 {
		return fd
	}

	fd.parseValuesDefinition()
	fd.parseKnownLabels()

	return fd
}

func (fd *fieldData) parseValuesDefinition() {
	values := strings.Split(fd.definition, ",")
	if len(values) == 0 {
		return
	}

	value := values[0]
	values = values[1:]

	if !envVariablePattern.MatchString(value) {
		return
	}

	envVarDefinition := newEnvVarDefinition(value)

	if envVarDefinition.varName != "" {
		fd.overrideKey = envVarDefinition.varName
		fd.keyIsOverriden = true
	}

	if envVarDefinition.defaultValue != "" {
		fd.defaultValue = envVarDefinition.defaultValue
		fd.hasDefaultValue = true
	}
}

func (fd *fieldData) parseKnownLabels() {
	for _, tag := range strings.Split(fd.definition, ",") {
		switch tag {
		case labelIsRequired:
			fd.isRequired = true
		}
	}
}
