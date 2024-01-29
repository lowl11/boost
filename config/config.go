package config

import (
	"github.com/lowl11/boost/data/interfaces"
	"github.com/lowl11/boost/internal/boosties/context"
	"github.com/lowl11/boost/log"
	"github.com/lowl11/boost/pkg/io/file"
	"os"
	"regexp"
	"strings"
)

var _configService *configService

func init() {
	load()
}

func Get(key string) interfaces.Param {
	value := _configService.Get(key)
	if value == "" {
		value = os.Getenv(key)
	}

	return context.NewParam(value)
}

func Env() string {
	envValue := strings.ToLower(Get("env").String())
	if envValue == "" {
		// check .env file exist
		if !file.Exist(".env") {
			return ""
		}

		envFileContent, err := file.Read(".env")
		if err != nil {
			return ""
		}

		reg, _ := regexp.Compile("((\\bENV\\b)|(\\benv\\b))=(.*)")
		match := reg.FindAllString(string(envFileContent), -1)
		if len(match) > 0 {
			splitValue := strings.Split(match[0], "=")
			if len(splitValue) < 2 {
				return ""
			}

			_ = os.Setenv("env", splitValue[1])
			return splitValue[1]
		}
	}

	return envValue
}

func IsProduction() bool {
	return Env() == "production"
}

func IsTest() bool {
	return Env() == "test"
}

func IsDev() bool {
	return Env() == "dev"
}

func load() {
	_configService = newService()
	if err := _configService.Read(); err != nil {
		log.Fatal("Load configuration error:", err)
	}
}
