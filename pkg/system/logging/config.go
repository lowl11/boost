package logging

import "github.com/lowl11/boost/data/interfaces"

const (
	defaultFileName   = "info"
	defaultFolderName = "logs"
)

var _config *Config

type Config struct {
	FileName      string
	FolderName    string
	NoFile        bool
	NoTime        bool
	NoPrefix      bool
	JsonMode      bool
	LogLevel      uint
	CustomLoggers []interfaces.ILogger
}

func Init(config Config) {
	if _config != nil {
		return
	}

	_config = &config
}

func GetConfig() *Config {
	if _config == nil {
		_config = &Config{
			FileName:   defaultFileName,
			FolderName: defaultFolderName,
		}
	}
	return _config
}

func GetFileAndFolder(config *Config) (string, string) {
	fileName := config.FileName
	folderName := config.FolderName
	if fileName == "" {
		fileName = defaultFileName
	}
	if folderName == "" {
		folderName = defaultFolderName
	}
	return fileName, folderName
}
