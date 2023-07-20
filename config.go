package boost

import (
	"github.com/lowl11/boost/internal/helpers/type_helper"
	"github.com/lowl11/lazyconfig/config/config_internal"
)

func initConfig(config Config) {
	configData := config_internal.Config{}

	configData.EnvironmentVariableName = type_helper.GetString(config.EnvironmentVariableName)
	configData.EnvironmentFileName = type_helper.GetString(config.EnvironmentFileName)
	configData.Environment = type_helper.GetString(config.Environment)
	configData.BaseFolder = type_helper.GetString(config.ConfigBaseFolder)

	config_internal.Init(configData)
}
