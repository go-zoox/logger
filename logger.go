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

func Trace(message string) {
	logrus.Trace(message)
}

func Debug(message string) {
	logrus.Debug(message)
}

func Info(message string) {
	logrus.Info(message)
}

func Warn(message string) {
	logrus.Warn(message)
}

func Error(message string) {
	logrus.Error(message)
}

func Fatal(message string) {
	logrus.Fatal(message)
}

func New() *logrus.Logger {
	return logrus.New()
}
