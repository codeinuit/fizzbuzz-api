package logrus

import (
	"os"

	"github.com/sirupsen/logrus"
)

const LOG_LEVEL string = "LOG_LEVEL"

type LogrusLogger struct {
	logger *logrus.Logger
}

// NewLogrusLogger returns a log instance from
// logger implementation
func NewLogrusLogger() *LogrusLogger {
	logger := logrus.New()

	lvl := os.Getenv(LOG_LEVEL)
	lvlParsed, err := logrus.ParseLevel(lvl)
	if err != nil {
		logger.WithError(err).Error("could not parse log level, defaulting to Info")
		lvlParsed = logrus.InfoLevel
	}
	logger.SetLevel(lvlParsed)

	return &LogrusLogger{
		logger: logger,
	}
}

// Debug prints logs at warn level
func (l LogrusLogger) Debug(v ...any) {
	l.logger.Debug(v...)
}

// Info prints logs at info level
func (l LogrusLogger) Info(v ...any) {
	l.logger.Info(v...)
}

// Error prints logs at error level
func (l LogrusLogger) Error(v ...any) {
	l.logger.Error(v...)
}

// Warn prints logs at warn level
func (l LogrusLogger) Warn(v ...any) {
	l.logger.Warn(v...)
}
