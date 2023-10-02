package rpc

import (
	"github.com/lowl11/boost/log"
	"google.golang.org/grpc"
)

func (app *App) Run(port string) error {
	return app.server.Start(port)
}

func (app *App) Close() {
	log.Error(app.server.Close(), "Shutdown gRPC server error")
}

func (app *App) RegisterService(desc *grpc.ServiceDesc, impl interface{}) {
	app.server.RegisterService(desc, impl)
}
