package boost

import (
	"github.com/lowl11/boost/config"
	"github.com/lowl11/boost/pkg/system/configuration"
	"github.com/lowl11/boost/pkg/system/types"
)

func initConfig(cfg Config) {
	configData := configuration.Config{}

	configData.EnvironmentVariableName = types.GetString(cfg.EnvironmentVariableName)
	configData.EnvironmentFileName = types.GetString(cfg.EnvironmentFileName)
	configData.Environment = types.GetString(cfg.Environment)
	configData.BaseFolder = types.GetString(cfg.ConfigBaseFolder)

	if configData.Environment == "" {
		configData.Environment = config.Env()
	}

	configuration.Init(configData)
}
