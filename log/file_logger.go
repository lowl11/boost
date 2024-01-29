package log

import (
	"fmt"
	"github.com/lowl11/boost/pkg/io/folder"
	"github.com/lowl11/boost/pkg/system/logging"
	"log"
	"os"
	"time"
)

type fileLogger struct {
	writer *log.Logger

	fileName string
	fileBase string
	filePath string
}

func newFile(fileBase, filePath string) *fileLogger {
	fileLoggerInstance := &fileLogger{
		fileBase: fileBase,
		filePath: filePath,
	}

	fileLoggerInstance.writer = log.New(fileLoggerInstance.createFile(), "", 0)
	return fileLoggerInstance
}

func (logger *fileLogger) Debug(args ...any) {
	logger.updateFile()
	var message string

	if !logging.GetConfig().JsonMode {
		logger.writer.SetPrefix(buildPrefix(debugLevel))
		message = buildMessage(args...)
	} else {
		message = buildJSON(jsonDebugLevel, args...)
	}

	logger.writer.Println(message)
}

func (logger *fileLogger) Info(args ...any) {
	logger.updateFile()
	var message string

	if !logging.GetConfig().JsonMode {
		logger.writer.SetPrefix(buildPrefix(infoLevel))
		message = buildMessage(args...)
	} else {
		message = buildJSON(jsonInfoLevel, args...)
	}

	logger.writer.Println(message)
}

func (logger *fileLogger) Warn(args ...any) {
	logger.updateFile()
	var message string

	if !logging.GetConfig().JsonMode {
		logger.writer.SetPrefix(buildPrefix(warnLevel))
		message = buildMessage(args...)
	} else {
		message = buildJSON(jsonWarnLevel, args...)
	}

	logger.writer.Println(message)
}

func (logger *fileLogger) Error(args ...any) {
	logger.updateFile()
	var message string

	if !logging.GetConfig().JsonMode {
		logger.writer.SetPrefix(buildPrefix(errorLevel))
		message = buildMessage(args...)
	} else {
		message = buildJSON(jsonErrorLevel, args...)
	}

	logger.writer.Println(message)
}

func (logger *fileLogger) Fatal(args ...any) {
	logger.updateFile()
	var message string

	if !logging.GetConfig().JsonMode {
		logger.writer.SetPrefix(buildPrefix(fatalLevel))
		message = buildMessage(args...)
	} else {
		message = buildJSON(jsonFatalLevel, args...)
	}

	logger.writer.Println(message)
}

const (
	fileNamePattern = "%s_%s.log"
)

func (logger *fileLogger) createFile() *os.File {
	// build log file name
	fileName := fmt.Sprintf(fileNamePattern, logger.fileBase, time.Now().Format("2006-01-02"))

	// destination folder
	if logger.filePath != "" {
		fileName = logger.filePath + "/" + fileName
		if !folder.Exist(logger.filePath) {
			if err := os.Mkdir(logger.filePath, os.ModePerm); err != nil {
				log.Println("Creating logs folder error")
			}
		}
	}

	// create file
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil
	}

	logger.fileName = fileName
	return file
}

func (logger *fileLogger) updateFile() {
	fileName := fmt.Sprintf(fileNamePattern, logger.fileBase, time.Now().Format("2006-01-02"))

	if logger.fileName != fileName {
		logger.writer = log.New(logger.createFile(), "", 0)
	}
}
