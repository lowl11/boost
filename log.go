package boost

import "github.com/lowl11/lazylog/log/log_internal"

func initLogger(config Config) {
	logConfig := log_internal.LogConfig{
		JsonMode: config.LogJSON,
		LogLevel: uint(config.LogLevel),
	}

	if config.LogFolderName != "" {
		logConfig.FolderName = config.LogFolderName
	}

	if config.LogFilePattern != "" {
		logConfig.FileName = config.LogFilePattern
	}

	if config.CustomLoggers != nil && len(config.CustomLoggers) > 0 {
		logConfig.CustomLoggers = config.CustomLoggers
	}

	log_internal.Init(logConfig)
}
