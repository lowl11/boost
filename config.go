package boost

import (
	"github.com/lowl11/boost/config"
	"github.com/lowl11/boost/pkg/system/configuration"
)

func initConfig(cfg Config) {
	configData := configuration.Config{}

	configData.EnvironmentVariableName = cfg.EnvironmentVariableName
	configData.EnvironmentFileName = cfg.EnvironmentFileName
	configData.Environment = cfg.Environment
	configData.BaseFolder = cfg.ConfigBaseFolder

	if configData.Environment == "" {
		configData.Environment = config.Env()
	}

	configuration.Init(configData)
}
