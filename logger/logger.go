package logger

import (
	"io/ioutil"

	"github.com/sirupsen/logrus"
)

// Logger is a simple logger type
// that outputs JSON logs
// and an optional map of custom fields & values
type Logger interface {
	Debug(string, map[string]interface{})
	Info(string, map[string]interface{})
	Error(string, map[string]interface{})
	Warn(string, map[string]interface{})
	Fatal(string, map[string]interface{})

	Disable()
	WithLevel(logrus.Level) Logger
}

const (
	// DebugLevel sets logger to log Debug lvl events and higher
	DebugLevel = logrus.DebugLevel

	// InfoLevel sets logger to log Info lvl events and higher
	InfoLevel = logrus.InfoLevel

	// ErrorLevel sets logger to log Error lvl events and higher
	ErrorLevel = logrus.ErrorLevel

	// WarnLevel sets logger to log Warn lvl events and higher
	WarnLevel = logrus.WarnLevel

	// FatalLevel sets logger to log Fatal lvl events and higher
	FatalLevel = logrus.FatalLevel
)

type logger struct {
	logrusLogger *logrus.Logger
}

// New returns a new Logger w/ default configs
func New() Logger {
	logrusLogger := logrus.New()
	logrusLogger.SetFormatter(&logrus.JSONFormatter{})
	logrusLogger.SetLevel(logrus.InfoLevel)

	return logger{
		logrusLogger: logrusLogger,
	}
}

func (l logger) Debug(msg string, fields map[string]interface{}) {
	l.logrusLogger.WithFields(fields).Debug(msg)
}

func (l logger) Info(msg string, fields map[string]interface{}) {
	l.logrusLogger.WithFields(fields).Info(msg)
}

func (l logger) Error(msg string, fields map[string]interface{}) {
	l.logrusLogger.WithFields(fields).Error(msg)
}

func (l logger) Warn(msg string, fields map[string]interface{}) {
	l.logrusLogger.WithFields(fields).Warn(msg)
}

func (l logger) Fatal(msg string, fields map[string]interface{}) {
	l.logrusLogger.WithFields(fields).Fatal(msg)
}

func (l logger) WithLevel(lvl logrus.Level) Logger {
	l.logrusLogger.SetLevel(lvl)
	return l
}

func (l logger) Disable() {
	l.logrusLogger.Out = ioutil.Discard
}
