package configuration

import (
	"github.com/lowl11/boost/internal/services/system/config_service"
	"github.com/lowl11/boost/log"
	"sync"
)

var (
	// _configServicePool contains *config_service.Service
	_configServicePool sync.Pool
	_initialized       bool
)

type Config struct {
	EnvironmentVariableName string
	EnvironmentFileName     string
	Environment             string
	BaseFolder              string
}

func Init(config Config) {
	configService := config_service.
		New()

	if config.EnvironmentVariableName != "" {
		configService.EnvironmentVariableName(config.EnvironmentVariableName)
	}

	if config.EnvironmentFileName != "" {
		configService.EnvironmentFileName(config.EnvironmentFileName)
	}

	if config.Environment != "" {
		configService.Environment(config.Environment)
	}

	if config.BaseFolder != "" {
		configService.BaseFolder(config.BaseFolder)
	}

	_configServicePool = sync.Pool{
		New: func() any {
			return configService
		},
	}

	_initialized = true
	if err := configService.Read(); err != nil {
		log.Fatal(err)
	}
}

func Get(key string) string {
	if !_initialized {
		return ""
	}

	configPtr := _configServicePool.Get().(*config_service.Service)
	if configPtr == nil {
		panic("configService is NULL")
	}

	return configPtr.Get(key)
}
