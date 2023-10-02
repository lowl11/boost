package log_levels

const (
	DEBUG = iota
	INFO
	WARN
	ERROR
	FATAL
)

func Check(level, checkLevel uint) bool {
	return level > checkLevel
}
