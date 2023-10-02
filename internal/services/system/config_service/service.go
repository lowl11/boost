package config_service

type Service struct {
	variables map[string]string

	baseFolder              string
	environment             string // dev, test, production or any other
	environmentBase         string // config.yml - base file
	environmentVariableName string // env, but can be environment for example
	environmentFileName     string // .env, but maybe it will be another file
}

func New() *Service {
	return &Service{
		variables: make(map[string]string),

		baseFolder:      defaultBaseFolder,
		environment:     defaultBaseFolder + defaultEnvironment,
		environmentBase: defaultBaseFolder + baseConfigName,

		environmentVariableName: defaultEnvironmentVariableName,
		environmentFileName:     defaultEnvironmentFileName,
	}
}
