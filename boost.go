package boost

import (
	"github.com/lowl11/boost/internal/fast_handler"
	"github.com/lowl11/boost/pkg/destroyer"
	"github.com/lowl11/boost/pkg/middlewares"
	"github.com/lowl11/lazylog/logapi"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Config is model of App configuration
type Config struct {
	// Timeout for every request in this App
	Timeout time.Duration

	// CustomLoggers is container with custom loggers inherited by interface logapi.ILogger
	CustomLoggers []logapi.ILogger

	// LogJSON turns on JSON mode. Logs will be printed in JSON format
	LogJSON bool
	// LogLevel controls which logs should be printed
	LogLevel int
	// LogFolderName change default logs containing folder name. Default folder is /logs
	LogFolderName string
	// LogFilePattern change default logs file names pattern. Default pattern is info
	LogFilePattern string

	// EnvironmentVariableName sets environment (dev/test/production) from OS variable name
	EnvironmentVariableName string
	// EnvironmentFileName sets environment file name. Default is .env
	EnvironmentFileName string
	// Environment sets environment value (dev/test/production)
	Environment string
	// ConfigBaseFolder sets base folder for profiles. Default is /profiles
	ConfigBaseFolder string
}

func defaultConfig() Config {
	return Config{
		Timeout: 0,
	}
}

// App is Boost application to run application
type App struct {
	config    Config
	handler   *fast_handler.Handler
	destroyer *destroyer.Destroyer
}

// New method creates new instance of Boost App
func New(configs ...Config) *App {
	// init config
	var config Config
	if len(configs) > 0 {
		config = configs[0]
	} else {
		config = defaultConfig()
	}

	// init
	initLogger(config)
	initConfig(config)

	// create Boost App instance
	app := &App{
		config:    config,
		handler:   fast_handler.New(),
		destroyer: destroyer.New(),
	}

	// catch shutdown signal
	go app.shutdown()

	// default middlewares
	app.Use(
		middlewares.CORS(),
		middlewares.Secure(),
	)

	// if timeout was set in config
	if config.Timeout != 0 {
		app.Use(middlewares.Timeout(config.Timeout))
	}

	return app
}

// shutdown handle signal for shutting down App
func (app *App) shutdown() {
	// create a channel to receive signals
	signalChannel := make(chan os.Signal, 1)

	// notify the signal channel when a SIGINT or SIGTERM signal is received
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)

	<-signalChannel

	// run destroy actions
	app.destroyer.Destroy()

	// call shutdown
	os.Exit(0)
}
