package logger

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

// type Logger struct {
// 	logrus.Logger
// }

func SetLevel(level string) (err error) {
	switch level {
	case "trace":
		logrus.SetLevel(logrus.TraceLevel)
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "warn", "warning":
		logrus.SetLevel(logrus.WarnLevel)
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	case "fatal":
		logrus.SetLevel(logrus.FatalLevel)
	default:
		err = fmt.Errorf("not a valid logger Level: %s, available: %s", level, "trace,debug,info,warn,error,fatal")
	}

	return err
}

func Trace(args ...interface{}) {
	logrus.Trace(args...)
}

func Debug(args ...interface{}) {
	logrus.Debug(args...)
}

func Info(args ...interface{}) {
	logrus.Info(args...)
}

func Warn(args ...interface{}) {
	logrus.Warn(args...)
}

func Error(args ...interface{}) {
	logrus.Error(args...)
}

func Fatal(args ...interface{}) {
	logrus.Fatal(args...)
}

func New() *logrus.Logger {
	return logrus.New()
}
