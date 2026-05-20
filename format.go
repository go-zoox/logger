package logger

// Debugf logs a debug message.
func Debugf(format string, args ...interface{}) {
	logger.Debug(format, args...)
}

// Infof logs an info message.
func Infof(format string, args ...interface{}) {
	logger.Info(format, args...)
}

// Warnf logs a warning message (routes by level to configured sinks).
func Warnf(format string, args ...interface{}) {
	logger.Warn(format, args...)
}

// Errorf logs an error message (routes to sinks with levels: [error], etc.).
func Errorf(format string, args ...interface{}) {
	logger.Error(format, args...)
}

// Fatalf logs a fatal message.
func Fatalf(format string, args ...interface{}) {
	logger.Fatal(format, args...)
}
