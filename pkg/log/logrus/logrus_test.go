package logrus_test

import (
	logger "codeinuit/fizzbuzz-api/pkg/log"
	"codeinuit/fizzbuzz-api/pkg/log/logrus"
	"testing"
)

func TestLogrusImplementation(t *testing.T) {
	var logr logger.Logger = logrus.NewLogrusLogger()
	_ = logr
}
