package file_logger

import (
	"fmt"
	"github.com/lowl11/boost/pkg/io/folder"
	"log"
	"os"
	"time"
)

const (
	fileNamePattern = "%s_%s.log"
)

func (logger *Logger) createFile() *os.File {
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

func (logger *Logger) updateFile() {
	fileName := fmt.Sprintf(fileNamePattern, logger.fileBase, time.Now().Format("2006-01-02"))

	if logger.fileName != fileName {
		logger.writer = log.New(logger.createFile(), "", 0)
	}
}
