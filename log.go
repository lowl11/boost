package boost

import (
	"github.com/lowl11/boost/pkg/system/logging"
	"github.com/lowl11/boost/pkg/system/types"
)

func initLogger(config Config) {
	logConfig := logging.Config{
		JsonMode: config.LogJSON,
		LogLevel: uint(config.LogLevel),
	}

	logConfig.FolderName = types.GetString(config.LogFolderName)
	logConfig.FileName = types.GetString(config.LogFilePattern)

	if config.CustomLoggers != nil && len(config.CustomLoggers) > 0 {
		logConfig.CustomLoggers = config.CustomLoggers
	}
}
