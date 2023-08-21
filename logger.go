package logger

import (
	"fmt"
	"os"
	"strings"

	"github.com/go-zoox/chalk"
	"github.com/go-zoox/datetime"
	"github.com/go-zoox/logger/components/constants"
	csc "github.com/go-zoox/logger/components/constants"
	csm "github.com/go-zoox/logger/components/message"
	cst "github.com/go-zoox/logger/components/transport"
	"github.com/go-zoox/logger/transport/console"
)

// Logger is the main logger object.
type Logger struct {
	level      string
	timeFormat string
	messageCh  chan *csm.Message
	transports map[string]cst.Transport
}

// Options is the options for the logger.
type Options struct {
	Level      string
	Transports map[string]cst.Transport
	TimeFormat string
}

// New creates a new logger object.
func New(options ...*Options) *Logger {
	level := DefaultLevel
	transports := map[string]cst.Transport{
		"console": console.New(),
	}
	timeFormat := "YYYY/MM/DD HH:mm:ss"

	if len(options) > 0 {
		if options[0].Level != "0" {
			level = options[0].Level
		}
		if options[0].Transports != nil {
			transports = options[0].Transports
		}
		if options[0].TimeFormat != "" {
			timeFormat = options[0].TimeFormat
		}
	}

	return &Logger{
		messageCh:  make(chan *csm.Message, csc.LogOutputBuffer),
		level:      level,
		timeFormat: timeFormat,
		transports: transports,
	}
}

// SetLevel sets the level of the logger.
func (l *Logger) SetLevel(level string) error {
	l.level = strings.ToUpper(level)
	return nil
}

// SetTransports sets the transports of the logger.
//
//	it will overrides system transport, if you only want append transport, just use AppendTransports.
func (l *Logger) SetTransports(transports map[string]cst.Transport) error {
	l.transports = transports
	return nil
}

// AppendTransports appends the transports of the logger.
func (l *Logger) AppendTransports(transports map[string]cst.Transport) error {
	for key, transport := range transports {
		if _, ok := l.transports[key]; ok {
			return fmt.Errorf("transport(%s) already used, please add another", key)
		}

		l.transports[key] = transport
	}

	return nil
}

// SetTimeFormat sets the time format.
func (l *Logger) SetTimeFormat(format string) {
	l.timeFormat = format
}

func (l *Logger) write(level string, format string, args ...interface{}) {
	if csc.LevelMap[l.level] > csc.LevelMap[level] {
		return
	}

	message := fmt.Sprintf(format, args...)

	time := datetime.Now().Format(l.timeFormat)

	levelX := chalk.Blue("INFO")
	switch level {
	case csc.LevelDebug:
		levelX = chalk.Gray("DEBUG")
	case csc.LevelInfo:
		levelX = chalk.Blue("INFO")
	case csc.LevelWarn:
		levelX = chalk.Yellow("WARN")
	case csc.LevelError:
		levelX = chalk.Red("ERROR")
	case csc.LevelFatal:
		levelX = chalk.Red("FATAL")
	}

	m := &csm.Message{
		Level:   level,
		Message: fmt.Sprintf("%s %s %s", time, levelX, message),
	}
	if constants.LevelMap[l.level] <= constants.LevelMap[m.Level] {
		for _, transport := range l.transports {
			transport.Write([]byte(m.Message))
		}
	}

	// fatal show exit after write log
	if level == csc.LevelFatal {
		os.Exit(1)
	}
}

// Debug logs a debug message.
func (l *Logger) Debug(format string, v ...interface{}) {
	l.write(csc.LevelDebug, format, v...)
}

// Info logs an info message.
func (l *Logger) Info(format string, v ...interface{}) {
	l.write(csc.LevelInfo, format, v...)
}

// Warn logs a warning message.
func (l *Logger) Warn(format string, v ...interface{}) {
	l.write(csc.LevelWarn, format, v...)
}

// Error logs an error message.
func (l *Logger) Error(format string, v ...interface{}) {
	l.write(csc.LevelError, format, v...)
}

// Fatal logs a fatal message.
func (l *Logger) Fatal(format string, v ...interface{}) {
	l.write(csc.LevelFatal, format, v...)
}

// Debugf logs a debug message.
func (l *Logger) Debugf(format string, v ...interface{}) {
	l.write(csc.LevelDebug, format, v...)
}

// Infof logs an info message.
func (l *Logger) Infof(format string, v ...interface{}) {
	l.write(csc.LevelInfo, format, v...)
}

// Warnf logs a warning message.
func (l *Logger) Warnf(format string, v ...interface{}) {
	l.write(csc.LevelWarn, format, v...)
}

// Errorf logs an error message.
func (l *Logger) Errorf(format string, v ...interface{}) {
	l.write(csc.LevelError, format, v...)
}

// Fatalf logs a fatal message.
func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.write(csc.LevelFatal, format, v...)
}

// Printf logs a message.
func (l *Logger) Printf(format string, v ...interface{}) {
	l.Infof(format, v...)
}
