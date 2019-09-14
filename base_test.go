package factor3

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// Setup:
	log.Disable()

	// Run Tests:
	code := m.Run()

	// Cleanup:

	// Finish:
	os.Exit(code)
}
