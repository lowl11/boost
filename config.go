package boost

import (
	"github.com/lowl11/boost/config"
	"github.com/lowl11/boost/internal/helpers/type_helper"
	"github.com/lowl11/boost/pkg/system/configuration"
)

func initConfig(cfg Config) {
	configData := configuration.Config{}

	configData.EnvironmentVariableName = type_helper.GetString(cfg.EnvironmentVariableName)
	configData.EnvironmentFileName = type_helper.GetString(cfg.EnvironmentFileName)
	configData.Environment = type_helper.GetString(cfg.Environment)
	configData.BaseFolder = type_helper.GetString(cfg.ConfigBaseFolder)

	if configData.Environment == "" {
		configData.Environment = config.Env()
	}

	configuration.Init(configData)
}
