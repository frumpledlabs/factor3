package factor3

import (
	"github.com/frumpled/factor3/logger"
)

var log logger.Logger

func init() {
	log = logger.
		New().
		WithLevel(logger.InfoLevel)

	// log.Info("Initializing env loader...", nil)
}
