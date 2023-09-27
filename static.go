package logger

import (
	"fmt"
	"os"

	"github.com/go-zoox/logger/components/constants"
	cst "github.com/go-zoox/logger/components/transport"
)

// DefaultLevel ...
var DefaultLevel = constants.LevelInfo

// LogLevelEnv ...
var LogLevelEnv = "LOG_LEVEL"

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

func init() {
	envLogLevel := os.Getenv(LogLevelEnv)
	if envLogLevel != "" {
		DefaultLevel = envLogLevel
		logger.SetLevel(DefaultLevel)
	}
}

// SetLevel sets the level of the logger.
func SetLevel(level string) (err error) {
	switch level {
	// case "trace":
	// 	logger.SetLevel(logger.TraceLevel)
	case "debug", LevelDebug:
		err = logger.SetLevel(LevelDebug)
	case "info", LevelInfo:
		err = logger.SetLevel(LevelInfo)
	case "warn", "warning", LevelWarn:
		err = logger.SetLevel(LevelWarn)
	case "error", LevelError:
		err = logger.SetLevel(LevelError)
	case "fatal", LevelFatal:
		err = logger.SetLevel(LevelFatal)
	default:
		err = fmt.Errorf("not a valid logger Level: %s, available: %s", level, "debug,info,warn,error,fatal")
	}

	return err
}

// GetLevel returns the level of the logger.
func GetLevel() string {
	return logger.GetLevel()
}

// SetTransports sets the transports of the logger.
//
//	it will overrides system transport, if you only want append transport, just use AppendTransports.
func SetTransports(transports map[string]cst.Transport) error {
	return logger.SetTransports(transports)
}

// AppendTransports appends the transports of the logger.
func AppendTransports(transports map[string]cst.Transport) error {
	return logger.AppendTransports(transports)
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

// IsDebugLevel returns true if the logger level is debug.
func IsDebugLevel() bool {
	return logger.IsDebugLevel()
}

// IsInfoLevel returns true if the logger level is info.
func IsInfoLevel() bool {
	return logger.IsInfoLevel()
}

// IsWarnLevel returns true if the logger level is warn.
func IsWarnLevel() bool {
	return logger.IsWarnLevel()
}

// IsErrorLevel returns true if the logger level is error.
func IsErrorLevel() bool {
	return logger.IsErrorLevel()
}

// IsFatalLevel returns true if the logger level is fatal.
func IsFatalLevel() bool {
	return logger.IsFatalLevel()
}
