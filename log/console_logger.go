package log

import (
	"github.com/lowl11/boost/internal/helpers/message_tools"
	"log"
	"os"
	"strings"
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

const (
	reset  = "\033[0m"
	red    = "\033[31m"
	green  = "\033[32m"
	yellow = "\033[33m"
	blue   = "\033[34m"
	purple = "\033[35m"
	cyan   = "\033[36m"
	gray   = "\033[37m"
	white  = "\033[97m"
)

func (logger *consoleLogger) Debug(args ...any) {
	var message string

	if !message_tools.JsonMode {
		logger.writer.SetPrefix(message_tools.BuildPrefix(debugLevel))
		message = message_tools.Build(args...)
	} else {
		message = message_tools.Json(jsonDebugLevel, args...)
	}

	logger.writer.Println(consoleDebug(message))
}

func (logger *consoleLogger) Info(args ...any) {
	var message string

	if !message_tools.JsonMode {
		logger.writer.SetPrefix(message_tools.BuildPrefix(infoLevel))
		message = message_tools.Build(args...)
	} else {
		message = message_tools.Json(jsonInfoLevel, args...)
	}

	logger.writer.Println(consoleInfo(message))
}

func (logger *consoleLogger) Warn(args ...any) {
	var message string

	if !message_tools.JsonMode {
		logger.writer.SetPrefix(message_tools.BuildPrefix(warnLevel))
		message = message_tools.Build(args...)
	} else {
		message = message_tools.Json(jsonWarnLevel, args...)
	}

	logger.writer.Println(consoleWarn(message))
}

func (logger *consoleLogger) Error(args ...any) {
	var message string

	if !message_tools.JsonMode {
		logger.writer.SetPrefix(message_tools.BuildPrefix(errorLevel))
		message = message_tools.Build(args...)
	} else {
		message = message_tools.Json(jsonErrorLevel, args...)
	}

	logger.writer.Println(consoleError(message))
}

func (logger *consoleLogger) Fatal(args ...any) {
	var message string

	if !message_tools.JsonMode {
		logger.writer.SetPrefix(message_tools.BuildPrefix(fatalLevel))
		message = message_tools.Build(args...)
	} else {
		message = message_tools.Json(jsonFatalLevel, args...)
	}

	logger.writer.Println(consoleFatal(message))
}

func consoleDebug(text string) string {
	return color(purple, text)
}

func consoleInfo(text string) string {
	return color(green, text)
}

func consoleWarn(text string) string {
	return color(yellow, text)
}

func consoleError(text string) string {
	return color(red, text)
}

func consoleFatal(text string) string {
	return color(gray, text)
}

func color(color, text string) string {
	coloredText := strings.Builder{}
	coloredText.WriteString(color)
	coloredText.WriteString(text)
	coloredText.WriteString(reset)
	return coloredText.String()
}
