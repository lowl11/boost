package rpc

import (
	"github.com/lowl11/boost/internal/boosties/grpc_server"
	"google.golang.org/grpc"
)

type Config struct {
	Options []grpc.ServerOption
}

func defaultConfig() Config {
	return Config{
		Options: []grpc.ServerOption{},
	}
}

type App struct {
	server *grpc_server.Server
}

func New(config ...Config) *App {
	var cfg Config
	if len(config) > 0 {
		cfg = config[0]
	} else {
		cfg = defaultConfig()
	}

	return &App{
		server: grpc_server.New(cfg.Options...),
	}
}
