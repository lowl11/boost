package log

func Debug(args ...any) {
	getLogger().Debug(args...)
}

func Debugf(format string, args ...any) {
	getLogger().Debugf(format, args...)
}

func Info(args ...any) {
	getLogger().Info(args...)
}

func Infof(format string, args ...any) {
	getLogger().Infof(format, args...)
}

func Warn(args ...any) {
	getLogger().Warn(args...)
}

func Warnf(format string, args ...any) {
	getLogger().Warnf(format, args...)
}

func Error(args ...any) {
	getLogger().Error(args...)
}

func Errorf(format string, args ...any) {
	getLogger().Errorf(format, args...)
}

func Fatal(args ...any) {
	getLogger().Fatal(args...)
}

func Fatalf(format string, args ...any) {
	getLogger().Fatalf(format, args...)
}
