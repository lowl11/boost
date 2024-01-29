package log

func Debug(args ...any) {
	getLogger().Debug(args...)
}

func Info(args ...any) {
	getLogger().Info(args...)
}

func Warn(args ...any) {
	getLogger().Warn(args...)
}

func Error(args ...any) {
	getLogger().Error(args...)
}

func Fatal(args ...any) {
	getLogger().Fatal(args...)
}
