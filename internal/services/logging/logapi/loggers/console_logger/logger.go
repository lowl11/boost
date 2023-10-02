package console_logger

import (
	"log"
	"os"
)

type Logger struct {
	writer *log.Logger
}

func Create() *Logger {
	writer := log.New(os.Stdout, "", 0)

	return &Logger{
		writer: writer,
	}
}
