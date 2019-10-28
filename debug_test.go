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
	defaultValue := "Passed - Default Value Used"
	overrideUnsetKey := "OVERRIDE_UNSET"

	os.Setenv(overrideSetKey, overrideSetValue)

	defer os.Unsetenv(overrideSetKey)

	input := struct {
		PlainField     string
		OverridenField string `env:"${OVERRIDE_SET}"`
		Embedded       struct {
			Field                           string
			FieldWithDefault                string `env:"${:-Passed - Default Value Used}"`
			UnsetOverriddenFieldWithDefault string `env:"${OVERRIDE_UNSET:-Passed - Default Value Used}"`
			SetOverriddenFieldWithDefault   string `env:"${OVERRIDE_SET:-Failed - Default value shouldn't be used}"`
		}
	}{}

	expectedOutput := map[string]fieldInfo{
		".PlainField": fieldInfo{
			EnvironmentVariable: "",
			DefaultValue:        "",
			CalculatedRawValue:  "",
		},
		".OverridenField": fieldInfo{
			EnvironmentVariable: overrideSetKey,
			DefaultValue:        "",
			CalculatedRawValue:  overrideSetValue,
		},
		".Embedded.Field": fieldInfo{
			EnvironmentVariable: "",
			DefaultValue:        "",
			CalculatedRawValue:  "",
		},
		".Embedded.FieldWithDefault": fieldInfo{
			EnvironmentVariable: "",
			DefaultValue:        defaultValue,
			CalculatedRawValue:  "",
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
	}

	var output map[string]fieldInfo
	output, err := debugFieldAndEnvironment("", &input)
	require.Nil(t, err)

	assert.Len(t, output, 6)

	for key := range expectedOutput {
		_, exists := output[key]
		assert.True(t, exists, key)
	}

	for key, value := range output {
		assert.Equal(t,
			expectedOutput[key].EnvironmentVariable,
			value.EnvironmentVariable,
		)
		// assert.Equal(t,
		// 	value.CalculatedRawValue,
		// 	expectedOutput[key].CalculatedRawValue,
		// )
		assert.Equal(t,
			expectedOutput[key].DefaultValue,
			value.DefaultValue,
		)
	}
}
