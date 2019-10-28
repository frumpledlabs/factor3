package factor3

import (
	"os"

	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_debug(t *testing.T) {
	overrideSetKey := "OVERRIDE_SET"
	overrideSetValue := "Passed - Override Set"
	defaultValue := "Default Value Used"
	overrideUnsetKey := "OVERRIDE_UNSET"

	os.Setenv(overrideSetKey, overrideSetValue)

	defer os.Unsetenv(overrideSetKey)

	input := struct {
		PlainField     string
		OverridenField string `env:"${OVERRIDE_SET}"`
		Embedded       struct {
			Field                           string
			FieldWithDefault                string `env:"${:-Default Value Used}"`
			UnsetOverriddenFieldWithDefault string `env:"${OVERRIDE_UNSET:-Default Value Used}"`
			SetOverriddenFieldWithDefault   string `env:"${OVERRIDE_SET:-Default Value Used}"`
			Deeply                          struct {
				NestedField string
			}
		}
	}{}

	expectedOutput := map[string]fieldInfo{
		".PlainField": fieldInfo{
			EnvironmentVariable: "PREFIX_PLAIN_FIELD",
			DefaultValue:        "",
			CalculatedRawValue:  nil,
		},
		".OverridenField": fieldInfo{
			EnvironmentVariable: overrideSetKey,
			DefaultValue:        "",
			CalculatedRawValue:  overrideSetValue,
		},
		".Embedded.Field": fieldInfo{
			EnvironmentVariable: "PREFIX_EMBEDDED_FIELD",
			DefaultValue:        "",
			CalculatedRawValue:  nil,
		},
		".Embedded.FieldWithDefault": fieldInfo{
			EnvironmentVariable: "PREFIX_EMBEDDED_FIELD_WITH_DEFAULT",
			DefaultValue:        defaultValue,
			CalculatedRawValue:  defaultValue,
		},
		".Embedded.UnsetOverriddenFieldWithDefault": fieldInfo{
			EnvironmentVariable: overrideUnsetKey,
			DefaultValue:        defaultValue,
			CalculatedRawValue:  defaultValue,
		},
		".Embedded.SetOverriddenFieldWithDefault": fieldInfo{
			EnvironmentVariable: overrideSetKey,
			DefaultValue:        defaultValue,
			CalculatedRawValue:  overrideSetValue,
		},
		".Embedded.Deeply.NestedField": fieldInfo{
			EnvironmentVariable: "PREFIX_EMBEDDED_DEEPLY_NESTED_FIELD",
			DefaultValue:        "",
			CalculatedRawValue:  nil,
		},
	}

	var output map[string]fieldInfo
	output, err := debugReadEnvironmentInto("PREFIX", &input)
	require.Nil(t, err)

	assert.Len(t, output, 7)

	for key := range expectedOutput {
		_, exists := output[key]
		assert.True(t, exists, key)
	}

	// for key := range output {
	// 	println("TEST:", key)
	// }

	for key, value := range output {
		assert.Equal(t,
			expectedOutput[key].EnvironmentVariable,
			value.EnvironmentVariable,
			key,
		)
		assert.Equal(t,
			expectedOutput[key].CalculatedRawValue,
			value.CalculatedRawValue,
			key,
		)
		assert.Equal(t,
			expectedOutput[key].DefaultValue,
			value.DefaultValue,
			key,
		)
	}
}
