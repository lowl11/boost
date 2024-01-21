package logging

import (
	"github.com/lowl11/boost/data/interfaces"
	"github.com/lowl11/boost/internal/services/logging/logapi"
)

const (
	defaultFileName   = "info"
	defaultFolderName = "logs"
)

var (
	_logger interfaces.ILogger
	_config *LogConfig
)

func _init() {
	if _config == nil {
		_config = &LogConfig{}
	}

	loggerInstance := logapi.New()

	// custom loggers
	for _, customLogger := range _config.CustomLoggers {
		loggerInstance.Custom(customLogger)
	}

	// file logger
	if !_config.NoFile {
		loggerInstance.File(getFileAndFolder(_config))
	}

	// no time flag
	if _config.NoTime {
		loggerInstance.NoTime()
	}

	// no prefix flag
	if _config.NoPrefix {
		loggerInstance.NoPrefix()
	}

	// json mode
	if _config.JsonMode {
		loggerInstance.JSON()
	}

	if _config.LogLevel > 0 {
		loggerInstance.Level(_config.LogLevel)
	}

	_logger = loggerInstance
}

func getFileAndFolder(config *LogConfig) (string, string) {
	fileName := config.FileName
	folderName := config.FolderName
	if fileName == "" {
		fileName = defaultFileName
	}
	if folderName == "" {
		folderName = defaultFolderName
	}
	return fileName, folderName
}
