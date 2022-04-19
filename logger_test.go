package logger

import "testing"

func TestLogger(t *testing.T) {
	Info("test info")
	Warn("test warn")
	Error("test error")

	err := SetLevel("trace")
	if err != nil {
		t.Error(err)
	}

	Trace("test trace")
	Debug("test debug")

	// Fatal("test fatal")
}
