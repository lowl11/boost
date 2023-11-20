package boost

import (
	"github.com/google/uuid"
	"github.com/lowl11/boost/data/enums/colors"
	"github.com/lowl11/boost/data/enums/modes"
	"github.com/lowl11/boost/internal/boosties/di_container"
	"github.com/lowl11/boost/internal/fast_handler"
	"github.com/lowl11/boost/internal/services/boost/greeting"
	"github.com/lowl11/boost/log"
	"github.com/lowl11/boost/pkg/system/cron"
	"github.com/lowl11/boost/pkg/system/types"
	"github.com/lowl11/boost/pkg/web/queue/msgbus"
	"github.com/lowl11/boost/pkg/web/rpc"
	"net/http"
)

const (
	methodAny  = "ANY"
	emptyGroup = ""
)

// Run starts listening TCP with given port
func (app *App) Run(port string) {
	if len(port) == 0 {
		panic("Port is empty")
	}

	// set cors config
	app.handler.SetCorsConfig(fast_handler.CorsConfig{
		Enabled: app.config.CorsEnabled,
		Origin:  app.config.CorsOrigin,
		Headers: app.config.CorsHeaders,
		Methods: app.config.CorsMethods,
		Vary:    app.config.CorsVary,
	})

	// check DI registers
	if err := di_container.Get().Check(); err != nil {
		log.Fatal(err, "Dependency Injection error")
	}

	// register static endpoints
	registerStaticEndpoints(app, app.healthcheck)

	// print greeting text
	greeting.New(app.handler.GetCounter(), greeting.Context{
		Mode: modes.Http,
		Port: port,
	}).
		MainColor(colors.Gray).
		SpecificColor(colors.Green).
		Print()

	// run server app
	log.Fatal(app.handler.Run(port))
}

func (app *App) RunFlag() {
	flagPort := Flag("port")
	if flagPort == "" {
		panic("Not given flag 'port=<PORT>'")
	}

	app.Run(flagPort)
}

// RunRPC starts listening TCP with given port
func (app *App) RunRPC(port string) {
	if app.rpcServer == nil {
		panic("RPC server is NULL. Add one handler at least to initialize")
	}

	// check DI registers
	if err := di_container.Get().Check(); err != nil {
		log.Fatal(err, "Dependency Injection error")
	}

	// print greeting text
	greeting.New(app.handler.GetCounter(), greeting.Context{
		Mode: modes.RPC,
		Port: port,
	}).
		MainColor(colors.Gray).
		SpecificColor(colors.Green).
		Print()

	// run server app
	log.Fatal(app.rpcServer.Run(port))
}

// RunCron starts listening TCP with given port
func (app *App) RunCron() {
	if app.cron == nil {
		panic("Cron App is NULL. Add at least one action to initialize")
	}

	// check DI registers
	if err := di_container.Get().Check(); err != nil {
		log.Fatal(err, "Dependency Injection error")
	}

	// print greeting text
	greeting.New(app.handler.GetCounter(), greeting.Context{
		Mode: modes.Cron,
	}).
		MainColor(colors.Gray).
		SpecificColor(colors.Green).
		Print()

	// run server app
	app.cron.Run()
}

func (app *App) RunListener(amqpConnectionURL string) {
	if app.listener == nil {
		panic("Message bus Listener is NULL. Add at least one binding to initialize")
	}

	// check DI registers
	if err := di_container.Get().Check(); err != nil {
		log.Fatal(err, "Dependency Injection error")
	}

	app.handler.GetCounter().ListenerBind(app.listener.EventsCount())

	// print greeting text
	greeting.New(app.handler.GetCounter(), greeting.Context{
		Mode: modes.Listener,
	}).
		MainColor(colors.Gray).
		SpecificColor(colors.Green).
		Print()

	// run server
	log.Fatal(app.listener.Run(amqpConnectionURL))
}

func (app *App) Listener() Listener {
	if app.listener == nil {
		app.listener = msgbus.NewListener()
	}

	return app.listener
}

// Destroy adds function which will be called in shutdown
func (app *App) Destroy(destroyFunc types.DestroyFunc) {
	if destroyFunc == nil {
		return
	}

	app.destroyer.AddFunction(destroyFunc)
}

// Healthcheck add new application service to healthcheck
func (app *App) Healthcheck(name, url string) {
	app.healthcheck.Register(name, url)
}

// ANY add new route to App with method ANY.
// Note: ANY will receive any type of HTTP method
func (app *App) ANY(path string, action HandlerFunc) Route {
	return app.handler.RegisterRoute(methodAny, path, action, emptyGroup)
}

// GET add new route to App with method GET
func (app *App) GET(path string, action HandlerFunc) Route {
	return app.handler.RegisterRoute(http.MethodGet, path, action, emptyGroup)
}

// POST add new route to App with method POST
func (app *App) POST(path string, action HandlerFunc) Route {
	return app.handler.RegisterRoute(http.MethodPost, path, action, emptyGroup)
}

// PUT add new route to App with method PUT
func (app *App) PUT(path string, action HandlerFunc) Route {
	return app.handler.RegisterRoute(http.MethodPut, path, action, emptyGroup)
}

// DELETE add new route to App with method DELETE
func (app *App) DELETE(path string, action HandlerFunc) Route {
	return app.handler.RegisterRoute(http.MethodDelete, path, action, emptyGroup)
}

func (app *App) groupANY(path string, action HandlerFunc, groupID string) Route {
	return app.handler.RegisterRoute(methodAny, path, action, groupID)
}

func (app *App) groupGET(path string, action HandlerFunc, groupID string) Route {
	return app.handler.RegisterRoute(http.MethodGet, path, action, groupID)
}

func (app *App) groupPOST(path string, action HandlerFunc, groupID string) Route {
	return app.handler.RegisterRoute(http.MethodPost, path, action, groupID)
}

func (app *App) groupPUT(path string, action HandlerFunc, groupID string) Route {
	return app.handler.RegisterRoute(http.MethodPut, path, action, groupID)
}

func (app *App) groupDELETE(path string, action HandlerFunc, groupID string) Route {
	return app.handler.RegisterRoute(http.MethodDelete, path, action, groupID)
}

// Group creates new group for endpoints with base url/endpoint
func (app *App) Group(base string) Group {
	app.handler.RegisterGroup()
	return newGroup(app, base)
}

// CronApp returns boost cron application
func (app *App) CronApp() cron.CronRouter {
	if app.cron == nil {
		app.cron = cron.New(app.config.CronConfig)
	}

	return app.cron
}

func (app *App) RpcApp() *rpc.App {
	if app.rpcServer == nil {
		app.rpcServer = rpc.New(app.config.RpcConfig)
	}

	return app.rpcServer
}

// Use add usable middleware to global App.
// Note: this method adds middleware function to every endpoint
func (app *App) Use(middlewareFunc ...MiddlewareFunc) {
	if len(middlewareFunc) == 0 {
		return
	}

	middlewares := make([]types.MiddlewareFunc, 0, len(middlewareFunc))

	for _, mFunc := range middlewareFunc {
		if mFunc == nil {
			continue
		}

		middlewares = append(middlewares, mFunc)
	}

	app.handler.RegisterGlobalMiddlewares(middlewares...)
}

func (app *App) useGroup(groupID uuid.UUID, middlewareFunc ...MiddlewareFunc) {
	if len(middlewareFunc) == 0 {
		return
	}

	middlewares := make([]types.MiddlewareFunc, 0, len(middlewareFunc))

	for _, mFunc := range middlewareFunc {
		if mFunc == nil {
			continue
		}

		middlewares = append(middlewares, mFunc)
	}

	app.handler.RegisterGroupMiddlewares(groupID, middlewares...)
}
