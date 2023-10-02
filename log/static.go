package log

import (
	"github.com/lowl11/boost/internal/boosties/logging"
)

func Debug(args ...any) {
	logging.Debug(args...)
}

func Info(args ...any) {
	logging.Info(args...)
}

func Warn(args ...any) {
	logging.Warn(args...)
}

func Error(err error, args ...any) {
	logging.Error(err, args...)
}

func Fatal(err error, args ...any) {
	logging.Fatal(err, args...)
}
