package log

import (
	"os"
	"testing"
)

func TestLogrusLogger(t *testing.T) {
	logger := NewLogrusLogger(os.Stdout, WithLevel(4))

	Debug(logger).Print("log", "test debug")
	Info(logger).Print("log", "test info")
	Warn(logger).Print("log", "test warn")
	Error(logger).Print("log", "test error")
}
