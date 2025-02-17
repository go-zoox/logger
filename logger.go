package logger

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/go-zoox/chalk"
	"github.com/go-zoox/datetime"
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

// Option is the options for the logger.
type Option struct {
	Level      string
	Transports map[string]cst.Transport
	TimeFormat string
}

// New creates a new logger object.
func New(option ...func(opt *Option)) *Logger {
	opt := &Option{
		Level: csc.LevelInfo,
		Transports: map[string]cst.Transport{
			"console": console.New(),
		},
		TimeFormat: "YYYY/MM/DD HH:mm:ss",
	}
	for _, o := range option {
		o(opt)
	}

	l := &Logger{
		messageCh:  make(chan *csm.Message, csc.LogOutputBuffer),
		timeFormat: opt.TimeFormat,
		transports: opt.Transports,
	}

	l.SetLevel(opt.Level)

	return l
}

// SetLevel sets the level of the logger.
func (l *Logger) SetLevel(level string) error {
	l.level = strings.ToUpper(level)
	return nil
}

// GetLevel gets the level of the logger.
func (l *Logger) GetLevel() string {
	return l.level
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

// SetStdout sets the stdout of the logger.
func (l *Logger) SetStdout(stdout io.Writer) {
	l.transports = map[string]cst.Transport{
		"console": console.New(func(opt *console.Option) {
			opt.Stdout = stdout
		}),
	}
}

func (l *Logger) write(level string, format string, args ...interface{}) {
	// fmt.Printf("[logger.write] 系统日志等级（%s），用户调用日志方法：%s => %v\n", l.level, level, csc.LevelMap[l.level] > csc.LevelMap[level])
	if level == LevelInfo {
		// @TODO History for author, loves use Infof for user, so we need to keep it.
	} else if csc.LevelMap[l.level] > csc.LevelMap[level] {
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
	for _, transport := range l.transports {
		transport.WriteWithLevel([]byte(m.Message), m.Level)
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
