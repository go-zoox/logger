package logger

// Debug logs a debug message.
func Debugf(format string, args ...interface{}) {
	logger.Debug(format, args...)
}

// Info logs an info message.
func Infof(format string, args ...interface{}) {
	logger.Info(format, args...)
}

// Warn logs a warning message.
func Warnf(format string, args ...interface{}) {
	logger.Warn(format, args...)
}

// Error logs an error message.
func Errorf(format string, args ...interface{}) {
	logger.Error(format, args...)
}

// Fatal logs a fatal message.
func Fatalf(format string, args ...interface{}) {
	logger.Fatal(format, args...)
}
