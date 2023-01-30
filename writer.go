package logger

// Write realizes io.Writer
func (l *Logger) Write(p []byte) (n int, err error) {
	for _, transport := range l.transports {
		if n, err := transport.Write(p); err != nil {
			return n, err
		}
	}

	return len(p), nil
}
