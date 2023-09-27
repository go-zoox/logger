package logger

// IsDebugLevel returns true if the logger level is debug.
func (l *Logger) IsDebugLevel() bool {
	return l.GetLevel() == LevelDebug
}

// IsInfoLevel returns true if the logger level is info.
func (l *Logger) IsInfoLevel() bool {
	return l.GetLevel() == LevelInfo
}

// IsWarnLevel returns true if the logger level is warn.
func (l *Logger) IsWarnLevel() bool {
	return l.GetLevel() == LevelWarn
}

// IsErrorLevel returns true if the logger level is error.
func (l *Logger) IsErrorLevel() bool {
	return l.GetLevel() == LevelError
}

// IsFatalLevel returns true if the logger level is fatal.
func (l *Logger) IsFatalLevel() bool {
	return l.GetLevel() == LevelFatal
}
