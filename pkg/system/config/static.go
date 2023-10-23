package config

import (
	"github.com/lowl11/boost/internal/boosties/configuration"
	"github.com/lowl11/boost/pkg/io/file"
	"os"
	"regexp"
	"strings"
)

func Get(key string) string {
	value := configuration.Get(key)
	if value == "" {
		value = os.Getenv(key)
	}

	return value
}

func Env() string {
	envValue := strings.ToLower(Get("env"))
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

func Load() {
	configuration.Load()
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
