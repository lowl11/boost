package logapi

type ILogger interface {
	Debug(args ...any)
	Info(args ...any)
	Warn(args ...any)
	Error(err error, args ...any)
	Fatal(err error, args ...any)
}
