package boost

import (
	"github.com/lowl11/boost/internal/fast_handler"
	"github.com/lowl11/boost/pkg/destroyer"
	"github.com/lowl11/lazyconfig/config/config_internal"
	"github.com/lowl11/lazylog/log/log_internal"
	"os"
	"os/signal"
	"syscall"
)

type App struct {
	handler   *fast_handler.Handler
	destroyer *destroyer.Destroyer
}

func New() *App {
	log_internal.Init(log_internal.LogConfig{})
	config_internal.Init(config_internal.Config{})

	app := &App{
		handler:   fast_handler.New(),
		destroyer: destroyer.New(),
	}
	go app.shutdown()

	return app
}

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
