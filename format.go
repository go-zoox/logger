package logger

// Debugf logs a debug message.
func Debugf(format string, args ...interface{}) {
	logger.Debug(format, args...)
}

// Infof logs an info message.
func Infof(format string, args ...interface{}) {
	logger.Info(format, args...)
}

// Warnf logs a warning message.
func Warnf(format string, args ...interface{}) {
	logger.Warn(format, args...)
}

// Errorf logs an error message.
func Errorf(format string, args ...interface{}) {
	logger.Error(format, args...)
}

// Fatalf logs a fatal message.
func Fatalf(format string, args ...interface{}) {
	logger.Fatal(format, args...)
}
