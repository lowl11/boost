package configuration

var _config *Config

type Config struct {
	EnvironmentVariableName string
	EnvironmentFileName     string
	Environment             string
	BaseFolder              string
}

func Init(config Config) {
	_config = &config
}

func Get() *Config {
	if _config == nil {
		_config = &Config{
			EnvironmentVariableName: "",
			EnvironmentFileName:     "",
			Environment:             "",
			BaseFolder:              "",
		}
	}

	return _config
}
