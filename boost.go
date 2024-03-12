package boost

import (
	baseContext "context"
	"github.com/google/uuid"
	"github.com/lowl11/boost/data/enums/colors"
	"github.com/lowl11/boost/data/enums/modes"
	"github.com/lowl11/boost/data/interfaces"
	"github.com/lowl11/boost/errors"
	"github.com/lowl11/boost/internal/di_container"
	"github.com/lowl11/boost/internal/fast_handler"
	"github.com/lowl11/boost/internal/greeting"
	"github.com/lowl11/boost/internal/healthcheck"
	"github.com/lowl11/boost/log"
	"github.com/lowl11/boost/pkg/system/cancel"
	"github.com/lowl11/boost/pkg/system/cron"
	"github.com/lowl11/boost/pkg/system/di"
	"github.com/lowl11/boost/pkg/system/types"
	"github.com/lowl11/boost/pkg/system/validator"
	"github.com/lowl11/boost/pkg/web/middlewares"
	"github.com/lowl11/boost/pkg/web/queue/msgbus"
	"github.com/lowl11/boost/pkg/web/rpc"
	"github.com/lowl11/boost/pkg/web/socket"

	"net/http"
	"os"
	"os/signal"
	"reflect"
	"syscall"
	"time"
)

const (
	methodAny  = "ANY"
	emptyGroup = ""
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

	// Cors config params
	CorsEnabled bool
	CorsOrigin  string
	CorsHeaders []string
	CorsMethods []string
	CorsVary    []string

	// Custom handler of panics
	PanicHandler types.PanicHandler
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
	cron        *cron.Cron
	healthcheck *healthcheck.Healthcheck
	listener    Listener

	ctx          baseContext.Context
	appCtxCancel func()
}

// New method creates new instance of Boost App
func New(configs ...Config) *App {
	// run initializer
	runInitializer()

	// register "Controller" interface type, it will be used by method "MapControllers()"
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
		log.Fatal("Create validator error:", err)
	}

	// turn off model validations
	if config.ValidationOff {
		validate.TurnOff()
	}

	appContext, appCtxCancel := baseContext.WithCancel(baseContext.Background())

	// create Boost App instance
	app := &App{
		config:      config,
		handler:     fast_handler.New(validate),
		healthcheck: healthcheck.New(),

		ctx:          appContext,
		appCtxCancel: appCtxCancel,
	}

	di_container.Get().RegisterImplementation(app.healthcheck)

	// default middlewares
	app.Use(
		middlewares.Secure(),
	)

	// if timeout was set in config
	if config.Timeout != 0 {
		app.Use(middlewares.Timeout(config.Timeout))
	}

	// set panic handler
	if config.PanicHandler != nil {
		app.handler.SetPanicHandler(config.PanicHandler)
	}

	// need to register "Controllers"
	di_container.Get().RegisterImplementation(app)
	di_container.Get().SetAppType(reflect.TypeOf(app))

	return app
}

func (app *App) Context() baseContext.Context {
	return app.ctx
}

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
		log.Fatal("Dependency Injection error:", err)
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
	app.handler.RunAsync(app.ctx, port)

	// shutdown app
	app.shutdown()
}

func (app *App) RunFlag() {
	// run app by using port from command argument
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
		log.Fatal("Dependency Injection error:", err)
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
		log.Fatal("Dependency Injection error:", err)
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
		log.Fatal("Dependency Injection error:", err)
	}

	// count listener binds (for greeting print)
	app.handler.GetCounter().ListenerBind(app.listener.EventsCount())

	// print greeting text
	greeting.New(app.handler.GetCounter(), greeting.Context{
		Mode: modes.Listener,
	}).
		MainColor(colors.Gray).
		SpecificColor(colors.Green).
		Print()

	// run server
	if err := app.listener.Run(amqpConnectionURL); err != nil {
		log.Fatal(err)
	}
}

// Listener returns message bus listener.
// Method return only single instance of listener i.e. singleton
func (app *App) Listener() Listener {
	if app.listener == nil {
		app.listener = msgbus.NewListener(app.ctx)
	}

	return app.listener
}

// Healthcheck add new application service to healthcheck
func (app *App) Healthcheck(name, url string) {
	app.healthcheck.Register(name, url)
}

func (app *App) UseStat(endpoint string) {
	app.GET(endpoint, staticEndpointStat(di.Get[healthcheck.Healthcheck]()))
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

func (app *App) Websocket(path string, handler socket.HandlerFunc) {
	websocketHandler(app, path, socket.NewHandler(handler))
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
		app.cron = cron.New(app.config.CronConfig, app.handler.GetCounter())
	}

	return app.cron
}

// RpcApp returns boost RPC application
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

	appMiddlewares := make([]types.MiddlewareFunc, 0, len(middlewareFunc))

	for _, mFunc := range middlewareFunc {
		if mFunc == nil {
			continue
		}

		appMiddlewares = append(appMiddlewares, mFunc)
	}

	app.handler.RegisterGlobalMiddlewares(appMiddlewares...)
}

func (app *App) useGroup(groupID uuid.UUID, middlewareFunc ...MiddlewareFunc) {
	if len(middlewareFunc) == 0 {
		return
	}

	groupMiddlewares := make([]types.MiddlewareFunc, 0, len(middlewareFunc))

	for _, mFunc := range middlewareFunc {
		if mFunc == nil {
			continue
		}

		groupMiddlewares = append(groupMiddlewares, mFunc)
	}

	app.handler.RegisterGroupMiddlewares(groupID, groupMiddlewares...)
}

// shutdown handle signal for shutting down App
func (app *App) shutdown() {
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, os.Kill, syscall.SIGTERM)
	<-signalChannel

	log.Info("Gracefully shutdown...")
	app.appCtxCancel()
	cancel.Get().Wait()
	log.Info("App shutdown")

	// exist from the app
	os.Exit(0)
}

func (app *App) NeedBoost() {
	app.handler.NeedBoost()
}

func websocketHandler(router routing, path string, handler *socket.Handler) {
	router.GET(path, socket.New(func(conn *socket.Conn) {
		for {
			messageType, message, err := conn.ReadMessage()
			if err != nil {
				log.Error("Websocket read message error:", err)
				continue
			}

			if err = handler.Run(conn, messageType, message); err != nil {
				log.Error("Run websocket handler error:", err)
				continue
			}
		}
	})).Use(func(ctx interfaces.Context) error {
		if ctx.Header("Connection") == "Upgrade" && ctx.Header("Upgrade") == "websocket" {
			ctx.FastHttpContext().SetUserValue("allowed", true)
			return ctx.Next()
		}

		return errors.
			New("Websocket required upgrade").
			SetType("Socket_RequiredUpgrade").
			SetHttpCode(http.StatusUpgradeRequired)
	})
}
