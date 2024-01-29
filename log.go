package boost

import (
	"github.com/lowl11/boost/internal/helpers/type_helper"
	"github.com/lowl11/boost/pkg/system/logging"
)

func initLogger(config Config) {
	logConfig := logging.Config{
		JsonMode: config.LogJSON,
		LogLevel: uint(config.LogLevel),
	}

	logConfig.FolderName = type_helper.GetString(config.LogFolderName)
	logConfig.FileName = type_helper.GetString(config.LogFilePattern)

	if config.CustomLoggers != nil && len(config.CustomLoggers) > 0 {
		logConfig.CustomLoggers = config.CustomLoggers
	}
}
