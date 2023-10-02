package config_service

import (
	"errors"
	"github.com/lowl11/boost/internal/helpers/env_helper"
	"github.com/lowl11/boost/pkg/io/file"
	"github.com/lowl11/boost/pkg/io/folder"
	"gopkg.in/yaml.v3"
	"os"
	"strings"
)

func (service *Service) Read() error {
	if !folder.Exist(service.baseFolder) {
		return errors.New("base folder does not exist: " + service.baseFolder)
	}

	// read .env file
	envFileContent, err := env_helper.Read(service.environmentFileName)
	if err != nil {
		return err
	}

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

func (service *Service) Get(key string) string {
	return service.variables[key]
}

func (service *Service) BaseFolder(baseFolder string) *Service {
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

func (service *Service) Environment(environment string) *Service {
	if environment == "" {
		return service
	}

	service.environment = environment
	return service
}

func (service *Service) EnvironmentVariableName(variableName string) *Service {
	if variableName == "" {
		return service
	}

	service.environmentVariableName = variableName
	return service
}

func (service *Service) EnvironmentFileName(fileName string) *Service {
	if fileName == "" {
		return service
	}

	service.environmentFileName = fileName
	return service
}
