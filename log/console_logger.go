package log

import (
	"encoding/json"
	"fmt"
	"github.com/lowl11/boost/pkg/system/logging"
	"github.com/lowl11/boost/pkg/system/types"
	"log"
	"os"
	"strings"
	"time"
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

	if !logging.GetConfig().JsonMode {
		logger.writer.SetPrefix(buildPrefix(debugLevel))
		message = buildMessage(args...)
	} else {
		message = buildJSON(jsonDebugLevel, args...)
	}

	logger.writer.Println(consoleDebug(message))
}

func (logger *consoleLogger) Debugf(format string, args ...any) {
	var message string

	if !logging.GetConfig().JsonMode {
		logger.writer.SetPrefix(buildPrefix(debugLevel))
		message = buildFormatMessage(format, args...)
	} else {
		message = buildFormatJSON(jsonDebugLevel, format, args...)
	}

	logger.writer.Println(consoleDebug(message))
}

func (logger *consoleLogger) Info(args ...any) {
	var message string

	if !logging.GetConfig().JsonMode {
		logger.writer.SetPrefix(buildPrefix(infoLevel))
		message = buildMessage(args...)
	} else {
		message = buildJSON(jsonInfoLevel, args...)
	}

	logger.writer.Println(consoleInfo(message))
}

func (logger *consoleLogger) Infof(format string, args ...any) {
	var message string

	if !logging.GetConfig().JsonMode {
		logger.writer.SetPrefix(buildPrefix(infoLevel))
		message = buildFormatMessage(format, args...)
	} else {
		message = buildFormatJSON(jsonInfoLevel, format, args...)
	}

	logger.writer.Println(consoleInfo(message))
}

func (logger *consoleLogger) Warn(args ...any) {
	var message string

	if !logging.GetConfig().JsonMode {
		logger.writer.SetPrefix(buildPrefix(warnLevel))
		message = buildMessage(args...)
	} else {
		message = buildJSON(jsonWarnLevel, args...)
	}

	logger.writer.Println(consoleWarn(message))
}

func (logger *consoleLogger) Warnf(format string, args ...any) {
	var message string

	if !logging.GetConfig().JsonMode {
		logger.writer.SetPrefix(buildPrefix(warnLevel))
		message = buildFormatMessage(format, args...)
	} else {
		message = buildFormatJSON(jsonWarnLevel, format, args...)
	}

	logger.writer.Println(consoleWarn(message))
}

func (logger *consoleLogger) Error(args ...any) {
	var message string

	if !logging.GetConfig().JsonMode {
		logger.writer.SetPrefix(buildPrefix(errorLevel))
		message = buildMessage(args...)
	} else {
		message = buildJSON(jsonErrorLevel, args...)
	}

	logger.writer.Println(consoleError(message))
}

func (logger *consoleLogger) Errorf(format string, args ...any) {
	var message string

	if !logging.GetConfig().JsonMode {
		logger.writer.SetPrefix(buildPrefix(errorLevel))
		message = buildFormatMessage(format, args...)
	} else {
		message = buildFormatJSON(jsonErrorLevel, format, args...)
	}

	logger.writer.Println(consoleError(message))
}

func (logger *consoleLogger) Fatal(args ...any) {
	var message string

	if !logging.GetConfig().JsonMode {
		logger.writer.SetPrefix(buildPrefix(fatalLevel))
		message = buildMessage(args...)
	} else {
		message = buildJSON(jsonFatalLevel, args...)
	}

	logger.writer.Println(consoleFatal(message))
}

func (logger *consoleLogger) Fatalf(format string, args ...any) {
	var message string

	if !logging.GetConfig().JsonMode {
		logger.writer.SetPrefix(buildPrefix(fatalLevel))
		message = buildFormatMessage(format, args...)
	} else {
		message = buildFormatJSON(jsonFatalLevel, format, args...)
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

func buildMessage(args ...any) string {
	if len(args) == 0 {
		return ""
	}

	stringArgs := strings.Builder{}
	for _, arg := range args {
		stringArgs.WriteString(types.ToString(arg))
		stringArgs.WriteString(" ")
	}
	return stringArgs.String()[:stringArgs.Len()-1]
}

func buildFormatMessage(format string, args ...any) string {
	message := strings.Builder{}
	_, _ = fmt.Fprintf(&message, format, args...)
	return message.String()
}

func buildPrefix(level string) string {
	prefix := strings.Builder{}
	if logging.GetConfig().NoPrefix && !logging.GetConfig().NoTime {
		prefix.WriteString(getTime())
		prefix.WriteString(" ")
		return prefix.String()
	}

	if logging.GetConfig().NoPrefix {
		return ""
	}

	if logging.GetConfig().NoTime {
		return level
	}

	prefix.WriteString(getTime())
	prefix.WriteString(" ")
	prefix.WriteString(level)
	return prefix.String()
}

func buildJSON(level string, args ...any) string {
	logMessage := &consoleLogMessage{
		Message: buildMessage(args...),
		Level:   level,
		Time:    getTime(),
	}

	logMessageInBytes, err := json.Marshal(logMessage)
	if err != nil {
		return "|ERROR IN BUILDING MESSAGE|"
	}

	return string(logMessageInBytes)
}

func buildFormatJSON(level, format string, args ...any) string {
	logMessage := &consoleLogMessage{
		Message: buildFormatMessage(format, args...),
		Level:   level,
		Time:    getTime(),
	}

	logMessageInBytes, err := json.Marshal(logMessage)
	if err != nil {
		return "|ERROR IN BUILDING MESSAGE|"
	}

	return string(logMessageInBytes)
}

func getTime() string {
	if logging.GetConfig().NoTime {
		return ""
	}

	return time.Now().Format("02-01-2006 15:04:05")
}

type consoleLogMessage struct {
	Time    string `json:"time,omitempty"`
	Level   string `json:"level"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}
