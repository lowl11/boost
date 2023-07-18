package boost

import (
	"github.com/lowl11/boost/internal/fast_handler"
	"github.com/lowl11/boost/pkg/destroyer"
	"github.com/lowl11/boost/pkg/middlewares"
	"github.com/lowl11/lazyconfig/config/config_internal"
	"github.com/lowl11/lazylog/log/log_internal"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Config is model of App configuration
type Config struct {
	Timeout time.Duration
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
	log_internal.Init(log_internal.LogConfig{})
	config_internal.Init(config_internal.Config{})

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
