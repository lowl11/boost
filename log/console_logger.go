package log

import (
	"github.com/lowl11/boost/internal/helpers/console_tools"
	"github.com/lowl11/boost/internal/helpers/message_tools"
	"log"
	"os"
)

type consoleLogger struct {
	writer *log.Logger
}

func newConsole() *consoleLogger {
	writer := log.New(os.Stdout, "", 0)

	return &consoleLogger{
		writer: writer,
	}
}

const (
	debugLevel = "[DEBUG] "
	infoLevel  = "[INFO] "
	warnLevel  = "[WARN] "
	errorLevel = "[ERROR] "
	fatalLevel = "[FATAL] "

	jsonDebugLevel = "DEBUG"
	jsonInfoLevel  = "INFO"
	jsonWarnLevel  = "WARN"
	jsonErrorLevel = "ERROR"
	jsonFatalLevel = "FATAL"
)

func (logger *consoleLogger) Debug(args ...any) {
	var message string

	if !message_tools.JsonMode {
		logger.writer.SetPrefix(message_tools.BuildPrefix(debugLevel))
		message = message_tools.Build(args...)
	} else {
		message = message_tools.Json(jsonDebugLevel, args...)
	}

	logger.writer.Println(console_tools.Debug(message))
}

func (logger *consoleLogger) Info(args ...any) {
	var message string

	if !message_tools.JsonMode {
		logger.writer.SetPrefix(message_tools.BuildPrefix(infoLevel))
		message = message_tools.Build(args...)
	} else {
		message = message_tools.Json(jsonInfoLevel, args...)
	}

	logger.writer.Println(console_tools.Info(message))
}

func (logger *consoleLogger) Warn(args ...any) {
	var message string

	if !message_tools.JsonMode {
		logger.writer.SetPrefix(message_tools.BuildPrefix(warnLevel))
		message = message_tools.Build(args...)
	} else {
		message = message_tools.Json(jsonWarnLevel, args...)
	}

	logger.writer.Println(console_tools.Warn(message))
}

func (logger *consoleLogger) Error(args ...any) {
	var message string

	if !message_tools.JsonMode {
		logger.writer.SetPrefix(message_tools.BuildPrefix(errorLevel))
		message = message_tools.Build(args...)
	} else {
		message = message_tools.Json(jsonErrorLevel, args...)
	}

	logger.writer.Println(console_tools.Error(message))
}

func (logger *consoleLogger) Fatal(args ...any) {
	var message string

	if !message_tools.JsonMode {
		logger.writer.SetPrefix(message_tools.BuildPrefix(fatalLevel))
		message = message_tools.Build(args...)
	} else {
		message = message_tools.Json(jsonFatalLevel, args...)
	}

	logger.writer.Println(console_tools.Fatal(message))
}
