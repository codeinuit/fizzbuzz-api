package logrus_test

import (
	"testing"

	logger "github.com/codeinuit/fizzbuzz-api/pkg/log"
	"github.com/codeinuit/fizzbuzz-api/pkg/log/logrus"
)

func TestLogrusImplementation(t *testing.T) {
	var logr logger.Logger = logrus.NewLogrusLogger()
	_ = logr
}
