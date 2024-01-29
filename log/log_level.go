package log

const (
	_DEBUG = iota
	_INFO
	_WARN
	_ERROR
	_FATAL
)

func checkLevel(level, checkLevel uint) bool {
	return level > checkLevel
}
