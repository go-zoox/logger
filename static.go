package logger

import (
	"fmt"

	"github.com/go-zoox/logger/components/constants"
)

var logger = New()
var (
	// LevelDebug is Level Debug
	LevelDebug = constants.LevelDebug
	// LevelInfo is Level Info
	LevelInfo = constants.LevelInfo
	// LevelWarn is Level Warn
	LevelWarn = constants.LevelWarn
	// LevelError is Level Error
	LevelError = constants.LevelError
	// LevelFatal is Level Fatal
	LevelFatal = constants.LevelFatal
)

// SetLevel sets the level of the logger.
func SetLevel(level string) (err error) {
	switch level {
	// case "trace":
	// 	logger.SetLevel(logger.TraceLevel)
	case "debug", LevelDebug:
		logger.SetLevel(LevelDebug)
	case "info", LevelInfo:
		logger.SetLevel(LevelInfo)
	case "warn", "warning", LevelWarn:
		logger.SetLevel(LevelWarn)
	case "error", LevelError:
		logger.SetLevel(LevelError)
	case "fatal", LevelFatal:
		logger.SetLevel(LevelFatal)
	default:
		err = fmt.Errorf("not a valid logger Level: %s, available: %s", level, "debug,info,warn,error,fatal")
	}

	return err
}

// func Trace(format string, args ...interface{}) {
// 	New().Tracef(format, args...)
// }

// Debug logs a debug message.
func Debug(format string, args ...interface{}) {
	logger.Debug(format, args...)
}

// Info logs an info message.
func Info(format string, args ...interface{}) {
	logger.Info(format, args...)
}

// Warn logs a warning message.
func Warn(format string, args ...interface{}) {
	logger.Warn(format, args...)
}

// Error logs an error message.
func Error(format string, args ...interface{}) {
	logger.Error(format, args...)
}

// Fatal logs a fatal message.
func Fatal(format string, args ...interface{}) {
	logger.Fatal(format, args...)
}
