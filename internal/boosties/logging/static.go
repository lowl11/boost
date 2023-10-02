package logging

import (
	"github.com/lowl11/boost/data/interfaces"
)

type LogConfig struct {
	FileName      string
	FolderName    string
	NoFile        bool
	NoTime        bool
	NoPrefix      bool
	JsonMode      bool
	LogLevel      uint
	CustomLoggers []interfaces.ILogger
}

func SetConfig(config LogConfig) {
	_config = &config
}

func Init(config LogConfig) {
	if _logger != nil {
		return
	}
	_config = &config

	_init()
}

func Initialized() bool {
	return _logger != nil
}

func Debug(args ...any) {
	if _logger == nil {
		_init()
	}
	_logger.Debug(args...)
}

func Info(args ...any) {
	if _logger == nil {
		_init()
	}
	_logger.Info(args...)
}

func Warn(args ...any) {
	if _logger == nil {
		_init()
	}
	_logger.Warn(args...)
}

func Error(err error, args ...any) {
	if _logger == nil {
		_init()
	}
	_logger.Error(err, args...)
}

func Fatal(err error, args ...any) {
	if _logger == nil {
		_init()
	}
	_logger.Fatal(err, args...)
}
