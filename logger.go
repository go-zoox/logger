package logger

import (
	"fmt"

	csc "github.com/go-zoox/logger/components/constants"
	csm "github.com/go-zoox/logger/components/message"
	cst "github.com/go-zoox/logger/components/transport"
	"github.com/go-zoox/logger/transport/console"
)

// Logger is the main logger object.
type Logger struct {
	level      int
	messageCh  chan *csm.Message
	transports map[string]cst.Transport
}

// Options is the options for the logger.
type Options struct {
	Level      int
	transports map[string]cst.Transport
}

// New creates a new logger object.
func New(options ...Options) *Logger {
	level := csc.LevelDebug
	transports := map[string]cst.Transport{
		"console": console.New(),
	}

	if len(options) > 0 {
		if options[0].Level > 0 {
			level = options[0].Level
		}
		if options[0].transports != nil {
			transports = options[0].transports
		}
	}

	return &Logger{
		messageCh:  make(chan *csm.Message, csc.LogOutputBuffer),
		level:      level,
		transports: transports,
	}
}

// SetLevel sets the level of the logger.
func (l *Logger) SetLevel(level int) {
	l.level = level
}

func (l *Logger) write(message string, level int) {
	if l.level > level {
		return
	}

	m := &csm.Message{
		Level:   level,
		Message: message,
	}
	for _, transport := range l.transports {
		transport.Write(m)
	}
}

func (l *Logger) getFormat(name string, format string) string {
	return fmt.Sprintf("[%s] ", name) + format
}

// Debug logs a debug message.
func (l *Logger) Debug(format string, v ...interface{}) {
	l.write(fmt.Sprintf(l.getFormat("DEBUG", format), v...), csc.LevelDebug)
}

// Info logs an info message.
func (l *Logger) Info(format string, v ...interface{}) {
	l.write(fmt.Sprintf(l.getFormat("INFO", format), v...), csc.LevelInfo)
}

// Warn logs a warning message.
func (l *Logger) Warn(format string, v ...interface{}) {
	l.write(fmt.Sprintf(l.getFormat("WARN", format), v...), csc.LevelWarn)
}

// Error logs an error message.
func (l *Logger) Error(format string, v ...interface{}) {
	l.write(fmt.Sprintf(l.getFormat("ERROR", format), v...), csc.LevelError)
}

// Fatal logs a fatal message.
func (l *Logger) Fatal(format string, v ...interface{}) {
	l.write(fmt.Sprintf(l.getFormat("FATAL", format), v...), csc.LevelFatal)
}
