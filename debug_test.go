package factor3

import (
	"os"

	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_debug(t *testing.T) {
	os.Setenv("OVERRIDE", "OVERRIDE_PASSED")
	os.Setenv("EMBEDDED_OVERRIDE_ENV", "EMBEDDED_OVERRIDE_ENV_PASSED")

	defer os.Unsetenv("OVERRIDE")
	defer os.Unsetenv("EMBEDDED_OVERRIDE_ENV")

	prefix := "PREFIX"
	input := struct {
		PlainField     string
		OverridenField string `env:"${OVERRIDE}"`
		Embedded       struct {
			Field          string
			OverridenField string `env:"${EMBEDDED_OVERRIDE_ENV}"`
			DefaultField   string `env:"${:-DEFAULT_VALUE}"`
		}
	}{}

	output, err := debugEnvironmentInto(prefix, &input)

	assert.Nil(t, err)

	println("Output:")
	for key, value := range output {
		println(key, '-', value)
	}
}
