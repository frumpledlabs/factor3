package factor3

import (
	"regexp"
	"strings"
)

type macroCaseReplacer struct {
	patternReplacements map[*regexp.Regexp]string
}

func NewMacroCaseReplacer() macroCaseReplacer {
	patternReplacements := map[*regexp.Regexp]string{
		regexp.MustCompile("([A-Z]+)([A-Z][a-z]+)"):       "${1}_${2}",
		regexp.MustCompile("([a-z]+)([A-Z]+)"):            "${1}_${2}",
		regexp.MustCompile("[\\._]+"):                     "_",
		regexp.MustCompile("([A-Z]+[a-z]*[0-9]+)([A-Z])"): "${1}_${2}",
	}

	return macroCaseReplacer{
		patternReplacements: patternReplacements,
	}
}

// This will attempt to convert straight-forward and conventional casing;
// it is not intended to replace complex or non-intuitive casing
func (m macroCaseReplacer) Replace(input string) string {
	result := input

	for pattern, replacement := range m.patternReplacements {
		result = pattern.ReplaceAllString(result, replacement)
	}

	result = strings.TrimSpace(result)
	result = strings.ToUpper(result)
	result = strings.Trim(result, "_")

	return result
}
