package boost

import (
	"github.com/lowl11/boost/pkg/system/logging"
)

func initLogger(config Config) {
	logConfig := logging.Config{
		JsonMode: config.LogJSON,
		LogLevel: uint(config.LogLevel),
	}

	logConfig.FolderName = config.LogFolderName
	logConfig.FileName = config.LogFilePattern

	if config.CustomLoggers != nil && len(config.CustomLoggers) > 0 {
		logConfig.CustomLoggers = config.CustomLoggers
	}
}
