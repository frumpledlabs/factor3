package factor3

import (
	"regexp"
	"strings"
)

const labelIsRequired = "required"

var valuesDefinitionPattern = regexp.MustCompile(`\${.*(:-)?.*}`)

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
	valueDefinition := strings.Split(fd.definition, ",")[0]

	if !valuesDefinitionPattern.MatchString(valueDefinition) {
		return
	}

	valueDefinition = strings.TrimLeft(valueDefinition, "${")
	valueDefinition = strings.TrimRight(valueDefinition, "}")

	valuesDefinition := strings.TrimLeft(valueDefinition, ".*:-")

	values := strings.Split(valuesDefinition, ",")

	hasPrefix := regexp.MustCompile(".+:-").
		MatchString(valueDefinition)
	if !hasPrefix {
		hasPrefix = !strings.Contains(valueDefinition, ":-") && valueDefinition != ""
	}

	hasSuffix := regexp.MustCompile(":-.+").MatchString(valueDefinition)

	if hasPrefix {
		override := values[0]
		override = strings.TrimRight(override, ":-")
		fd.overrideKey = override
		fd.keyIsOverriden = true

		values = values[1:]
	}

	if hasSuffix {
		fd.defaultValue = values[0]
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
