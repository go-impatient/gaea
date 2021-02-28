package log

import (
	"testing"
)

func TestLogger(t *testing.T) {
	logger := DefaultLogger
	Debug(logger).Print("log", "test debug")
	Info(logger).Print("log", "test info")
	Warn(logger).Print("log", "test warn")
	Error(logger).Print("log", "test error")

	log := LogrusLogger
	Debug(log).Print("log", "test debug")
	Info(log).Print("log", "test info")
	Warn(log).Print("log", "test warn")
	Error(log).Print("log", "test error")
}
