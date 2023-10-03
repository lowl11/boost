package boost

import (
	"github.com/lowl11/boost/internal/boosties/di_container"
	"github.com/lowl11/boost/internal/fast_handler"
	"github.com/lowl11/boost/internal/services/boost/healthcheck"
	"github.com/lowl11/boost/internal/services/system/validator"
	"github.com/lowl11/boost/log"
	"github.com/lowl11/boost/pkg/system/cron"
	"github.com/lowl11/boost/pkg/web/destroyer"
	middlewares2 "github.com/lowl11/boost/pkg/web/middlewares"
	"github.com/lowl11/boost/pkg/web/rpc"
	"os"
	"os/signal"
	"reflect"
	"syscall"
	"time"
)

// Config is model of App configuration
type Config struct {
	// Timeout for every request in this App
	Timeout time.Duration

	// ValidationOff turns of validation
	ValidationOff bool

	// CustomLoggers is container with custom loggers inherited by interface logapi.ILogger
	CustomLoggers []ILogger

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

	// CronConfig config of Cron Application
	CronConfig cron.Config

	// RpcConfig config of gRPC Application
	RpcConfig rpc.Config
}

func defaultConfig() Config {
	return Config{
		Timeout: 0,
	}
}

// App is Boost application to run application
type App struct {
	config      Config
	handler     *fast_handler.Handler
	rpcServer   *rpc.App
	destroyer   *destroyer.Destroyer
	cron        *cron.Cron
	healthcheck *healthcheck.Healthcheck
	listener    Listener
}

// New method creates new instance of Boost App
func New(configs ...Config) *App {
	di_container.Get().SetControllerInterface(reflect.TypeOf(new(Controller)))

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

	// create validator
	validate, err := validator.New()
	if err != nil {
		log.Fatal(err, "Create validator error")
	}

	// turn off model validations
	if config.ValidationOff {
		validate.TurnOff()
	}

	// create Boost App instance
	app := &App{
		config:      config,
		handler:     fast_handler.New(validate),
		destroyer:   destroyer.New(),
		healthcheck: healthcheck.New(),
	}

	// catch shutdown signal
	go app.shutdown()

	// default middlewares
	app.Use(
		middlewares2.CORS(),
		middlewares2.Secure(),
	)

	// if timeout was set in config
	if config.Timeout != 0 {
		app.Use(middlewares2.Timeout(config.Timeout))
	}

	di_container.Get().RegisterImplementation(app)
	di_container.Get().SetAppType(reflect.TypeOf(app))
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
