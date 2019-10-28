package factor3

import (
	// "fmt"
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
			DefaultValueField               string `env:"${:-Passed Default Value}"`
			UnsetOverriddenFieldWithDefault string `env:"${OVERRIDE_UNSET:-Passed - Default Value}"`
			SetOverriddenFieldWithDefault   string `env:"${OVERRIDE_SET:-Failed - Default value shouldn't be used}"`
		}
	}{}

	expectedOutput := map[string]fieldInfo{
		"PlainField": fieldInfo{
			EnvironmentVariable: "",
			DefaultValue:        "",
			CalculatedValue:     "",
		},
		"OverridenField": fieldInfo{
			EnvironmentVariable: overrideSetKey,
			DefaultValue:        "",
			CalculatedValue:     overrideSetValue,
		},
		"Embedded.Field": fieldInfo{
			EnvironmentVariable: "",
			DefaultValue:        "",
			CalculatedValue:     "",
		},
		"Embedded.DefaultValueField": fieldInfo{
			EnvironmentVariable: "",
			DefaultValue:        defaultValue,
			CalculatedValue:     "",
		},
		"Embedded.UnsetOverriddenFieldWithDefault": fieldInfo{
			EnvironmentVariable: overrideUnsetKey,
			DefaultValue:        defaultValue,
			CalculatedValue:     defaultValue,
		},
		"Embedded.SetOverriddenFieldWithDefault": fieldInfo{
			EnvironmentVariable: overrideSetKey,
			DefaultValue:        defaultValue,
			CalculatedValue:     overrideSetValue,
		},
	}

	var output map[string]fieldInfo
	output, err := debugFieldAndEnvironment("", &input)
	require.Nil(t, err)

	// assert.NotEqual(t, input, prefix)
	assert.NotEqual(t, output, expectedOutput)

	assert.Len(t, output, 6)
	// for key, value := range output {
	// 	// assert.Equal(
	// 	// 	t, value.CalculatedValue, expectedOutput[key].CalculatedValue,
	// 	// )
	// 	assert.Equal(
	// 		t, value.DefaultValue, expectedOutput[key].DefaultValue,
	// 	)
	// 	// assert.Equal(
	// 	// 	t, value.EnvironmentVariable, expectedOutput[key].EnvironmentVariable,
	// 	// )
	// }
}
