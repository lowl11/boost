package file_logger

import (
	"log"
)

type Logger struct {
	writer *log.Logger

	fileName string
	fileBase string
	filePath string
}

func Create(fileBase, filePath string) *Logger {
	logger := &Logger{
		fileBase: fileBase,
		filePath: filePath,
	}
	logger.writer = log.New(logger.createFile(), "", 0)

	return logger
}
