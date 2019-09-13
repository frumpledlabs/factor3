package factor3

import (
	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

func init() {
	logrusLogger := logrus.New()
	logrusLogger.SetFormatter(&logrus.JSONFormatter{})
	logrusLogger.SetLevel(logrus.InfoLevel)

	log = logrusLogger
}
