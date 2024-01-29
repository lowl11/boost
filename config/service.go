package config

import (
	"github.com/lowl11/boost/internal/helpers/env_helper"
	"github.com/lowl11/boost/pkg/io/file"
	"github.com/lowl11/boost/pkg/io/folder"
	"gopkg.in/yaml.v3"
	"os"
	"strings"
)

const (
	defaultBaseFolder              = "profiles/"
	defaultEnvironment             = "dev"
	defaultEnvironmentVariableName = "env"
	defaultEnvironmentFileName     = ".env"

	baseConfigName = "config.yml"
)

type configService struct {
	variables    map[string]string
	envVariables map[string]string

	baseFolder              string
	environment             string // dev, test, production or any other
	environmentBase         string // config.yml - base file
	environmentVariableName string // env, but can be environment for example
	environmentFileName     string // .env, but maybe it will be another file
}

func newService() *configService {
	return &configService{
		variables: make(map[string]string),

		baseFolder:      defaultBaseFolder,
		environment:     defaultBaseFolder + defaultEnvironment,
		environmentBase: defaultBaseFolder + baseConfigName,

		environmentVariableName: defaultEnvironmentVariableName,
		environmentFileName:     defaultEnvironmentFileName,
	}
}

func (service *configService) Load() {
	if service.envVariables == nil {
		return
	}

	for key, value := range service.envVariables {
		_ = os.Setenv(key, value)
	}
}

func (service *configService) Read() error {
	// read .env file
	envFileContent, err := env_helper.Read(service.environmentFileName)
	if err != nil {
		return err
	}

	for key, value := range envFileContent {
		_ = os.Setenv(key, value)
	}

	if !folder.Exist(service.baseFolder) {
		//return errors.New("base folder does not exist: " + service.baseFolder)
		return nil
	}

	service.envVariables = envFileContent

	baseVariables := make(map[string]string)

	if file.Exist(service.environmentBase) {
		// read base config.yml file
		envBaseContent, err := file.Read(service.environmentBase)
		if err != nil {
			return err
		}

		envBaseContent, err = env_helper.ReplaceVariables(envBaseContent, envFileContent)
		if err != nil {
			return err
		}

		if err = yaml.Unmarshal(envBaseContent, &baseVariables); err != nil {
			return err
		}
	}

	// build <environment>.yml file name
	envFileName := service.environment + ".yml"
	if !strings.Contains(envFileName, defaultBaseFolder) {
		envFileName = defaultBaseFolder + envFileName
	}

	// read <environment>.yml file
	envContent, err := file.Read(envFileName)
	if err != nil {
		return err
	}

	// replace variables from file
	config, err := env_helper.ReplaceVariables(envContent, envFileContent)
	if err != nil {
		return err
	}

	// set data to values map
	if err = yaml.Unmarshal(config, &service.variables); err != nil {
		return err
	}

	// check if there is no such variable
	// even if variable with such key exist, need to check if in current it is empty
	for key, baseValue := range baseVariables {
		if currentValue, ok := service.variables[key]; (!ok || currentValue == "") && baseValue != "" {
			// if basic config value is "variable"
			if varValue, isVariable := env_helper.IsVariable(baseValue); isVariable {
				// try getting value of variable from .env file
				fileValue, fileOk := envFileContent[varValue]
				if fileOk {
					service.variables[key] = fileValue
				} else {
					// if no .env value, try search it in environment
					osValue := os.Getenv(varValue)
					if osValue != "" {
						service.variables[key] = os.Getenv(varValue)
					}
				}

				// if there is no value in .env file & environment
				if service.variables[key] != "" {
					continue
				}
			}

			// set value
			service.variables[key] = baseValue
		}
	}

	return nil
}

func (service *configService) Get(key string) string {
	return service.variables[key]
}

func (service *configService) BaseFolder(baseFolder string) *configService {
	if baseFolder == "" {
		return service
	}

	if baseFolder[len(baseFolder)-1] != '/' {
		baseFolder += "/"
	}

	// update all paths
	service.baseFolder = baseFolder
	service.environment = service.baseFolder + service.environment
	service.environment = service.baseFolder + service.environmentBase
	return service
}

func (service *configService) Environment(environment string) *configService {
	if environment == "" {
		return service
	}

	service.environment = environment
	return service
}

func (service *configService) EnvironmentVariableName(variableName string) *configService {
	if variableName == "" {
		return service
	}

	service.environmentVariableName = variableName
	return service
}

func (service *configService) EnvironmentFileName(fileName string) *configService {
	if fileName == "" {
		return service
	}

	service.environmentFileName = fileName
	return service
}
