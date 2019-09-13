package factor3

import (
	"github.com/sirupsen/logrus"
)

// Logger is a simple logger type
// that outputs JSON logs
// and an optional map of custom fields & values
type Logger interface {
	Info(string, map[string]interface{})
	Debug(string, map[string]interface{})
	Warn(string, map[string]interface{})
	Fatal(string, map[string]interface{})
}

type logger struct {
	*logrus.Logger
}

// NewLogger returns a new logger w/ default configs
func NewLogger() Logger {
	logrusLogger := logrus.New()
	logrusLogger.SetFormatter(&logrus.JSONFormatter{})
	logrusLogger.SetLevel(logrus.InfoLevel)

	return logger{
		logrusLogger,
	}
}

func (l logger) Debug(msg string, fields map[string]interface{}) {
	l.WithFields(fields).Debug(msg)
}

func (l logger) Info(msg string, fields map[string]interface{}) {
	l.WithFields(fields).Info(msg)
}

func (l logger) Warn(msg string, fields map[string]interface{}) {
	l.WithFields(fields).Warn(msg)
}

func (l logger) Fatal(msg string, fields map[string]interface{}) {
	l.WithFields(fields).Fatal(msg)
}
