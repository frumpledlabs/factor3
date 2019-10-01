package factor3

import (
	"regexp"
	"strings"
)

const labelIsRequired = "required"

var valuesDefinitionPattern = regexp.MustCompile(`\${.*(:-)?.*}`)

type tagSet struct {
	definition      string
	defaultValue    string
	overrideKey     string
	isRequired      bool
	keyIsOverriden  bool
	hasDefaultValue bool
}

func newTagSet(input string) tagSet {
	ts := tagSet{
		definition: input,
	}

	tags := strings.Split(input, ",")
	if len(tags) == 0 {
		return ts
	}

	ts.parseValuesDefinition()
	ts.parseKnownLabels()

	return ts
}

func (t *tagSet) parseValuesDefinition() {
	valueDefinition := strings.Split(t.definition, ",")[0]

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
		t.overrideKey = override
		t.keyIsOverriden = true

		values = values[1:]
	}

	if hasSuffix {
		t.defaultValue = values[0]
		t.hasDefaultValue = true
	}
}

func (t *tagSet) parseKnownLabels() {
	for _, tag := range strings.Split(t.definition, ",") {
		switch tag {
		case labelIsRequired:
			t.isRequired = true
		}
	}
}
