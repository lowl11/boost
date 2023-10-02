package boost

import (
	"github.com/lowl11/boost/internal/boosties/logging"
	"github.com/lowl11/boost/internal/helpers/type_helper"
)

func initLogger(config Config) {
	logConfig := logging.LogConfig{
		JsonMode: config.LogJSON,
		LogLevel: uint(config.LogLevel),
	}

	logConfig.FolderName = type_helper.GetString(config.LogFolderName)
	logConfig.FileName = type_helper.GetString(config.LogFilePattern)

	if config.CustomLoggers != nil && len(config.CustomLoggers) > 0 {
		logConfig.CustomLoggers = config.CustomLoggers
	}

	logging.Init(logConfig)
}
