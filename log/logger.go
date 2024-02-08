package log

import (
	"github.com/lowl11/boost/data/interfaces"
	"github.com/lowl11/boost/pkg/system/logging"
	"os"
	"sync"
	"time"
)

const (
	defaultExitDuration = 250
)

type logger struct {
	loggers       []interfaces.ILogger
	customLoggers []interfaces.ILogger
	mutex         sync.Mutex
	line          *line

	exitDuration     time.Duration
	customDuration   time.Duration
	isCustomDuration bool

	level uint
}

func getLogger() *logger {
	if _logger != nil {
		return _logger
	}

	_logger = &logger{
		loggers: []interfaces.ILogger{
			newConsole(),
		},
		mutex:        sync.Mutex{},
		exitDuration: time.Millisecond * defaultExitDuration,
	}

	config := logging.GetConfig()

	// custom loggers
	for _, customLogger := range config.CustomLoggers {
		_logger.Custom(customLogger)
	}

	// file logger
	if !config.NoFile {
		_logger.File(logging.GetFileAndFolder(config))
	}

	// no time flag
	if config.NoTime {
		_logger.NoTime()
	}

	// no prefix flag
	if config.NoPrefix {
		_logger.NoPrefix()
	}

	// json mode
	if config.JsonMode {
		_logger.JSON()
	}

	if config.LogLevel > 0 {
		_logger.Level(config.LogLevel)
	}

	return _logger
}

func (logger *logger) Level(level uint) *logger {
	if level > _FATAL {
		return logger
	}

	logger.level = level
	return logger
}

func (logger *logger) File(fileBase string, filePath ...string) *logger {
	var singleFilePath string
	if len(filePath) > 0 {
		singleFilePath = filePath[0]
	}

	logger.loggers = append(logger.loggers, newFile(fileBase, singleFilePath))
	return logger
}

func (logger *logger) Custom(customLogger interfaces.ILogger) *logger {
	logger.customLoggers = append(logger.customLoggers, customLogger)
	if !logger.isCustomDuration {
		logger.exitDuration = logger.exitDuration + time.Millisecond*defaultExitDuration
	}

	if logger.line == nil && len(logger.customLoggers) > 0 {
		logger.line = newLine()
	}

	return logger
}

func (logger *logger) CustomExitDuration(duration time.Duration) *logger {
	if duration < defaultExitDuration {
		return logger
	}

	logger.isCustomDuration = true
	logger.customDuration = duration

	return logger
}

func (logger *logger) JSON() *logger {
	logging.GetConfig().JsonMode = true
	return logger
}

func (logger *logger) NoTime() *logger {
	logging.GetConfig().NoTime = true
	return logger
}

func (logger *logger) NoPrefix() *logger {
	logging.GetConfig().NoPrefix = true
	return logger
}

func (logger *logger) Debug(args ...any) {
	if len(args) == 0 {
		return
	}

	logger.mutex.Lock()
	defer logger.mutex.Unlock()

	// skip log by level
	if checkLevel(logger.level, _DEBUG) {
		return
	}

	for _, loggerItem := range logger.loggers {
		loggerItem.Debug(args...)
	}

	for _, customLogger := range logger.customLoggers {
		logger.line.AddInfo(customLogger.Debug, args...)
	}
}

func (logger *logger) Debugf(format string, args ...any) {
	logger.mutex.Lock()
	defer logger.mutex.Unlock()

	// skip log by level
	if checkLevel(logger.level, _DEBUG) {
		return
	}

	for _, loggerItem := range logger.loggers {
		loggerItem.Debugf(format, args...)
	}

	for _, customLogger := range logger.customLoggers {
		logger.line.AddFormatInfo(customLogger.Debugf, format, args...)
	}
}

func (logger *logger) Info(args ...any) {
	if len(args) == 0 {
		return
	}

	logger.mutex.Lock()
	defer logger.mutex.Unlock()

	// skip log by level
	if checkLevel(logger.level, _INFO) {
		return
	}

	for _, loggerItem := range logger.loggers {
		loggerItem.Info(args...)
	}

	for _, customLogger := range logger.customLoggers {
		logger.line.AddInfo(customLogger.Info, args...)
	}
}

func (logger *logger) Infof(format string, args ...any) {
	logger.mutex.Lock()
	defer logger.mutex.Unlock()

	// skip log by level
	if checkLevel(logger.level, _INFO) {
		return
	}

	for _, loggerItem := range logger.loggers {
		loggerItem.Infof(format, args...)
	}

	for _, customLogger := range logger.customLoggers {
		logger.line.AddFormatInfo(customLogger.Infof, format, args...)
	}
}

func (logger *logger) Warn(args ...any) {
	if len(args) == 0 {
		return
	}

	logger.mutex.Lock()
	defer logger.mutex.Unlock()

	// skip log by level
	if checkLevel(logger.level, _WARN) {
		return
	}

	for _, loggerItem := range logger.loggers {
		loggerItem.Warn(args...)
	}

	for _, customLogger := range logger.customLoggers {
		logger.line.AddInfoCustom(customLogger.Warn, args...)
	}
}

func (logger *logger) Warnf(format string, args ...any) {
	logger.mutex.Lock()
	defer logger.mutex.Unlock()

	// skip log by level
	if checkLevel(logger.level, _WARN) {
		return
	}

	for _, loggerItem := range logger.loggers {
		loggerItem.Warnf(format, args...)
	}

	for _, customLogger := range logger.customLoggers {
		logger.line.AddInfoFormatCustom(customLogger.Warnf, format, args...)
	}
}

func (logger *logger) Error(args ...any) {
	logger.mutex.Lock()
	defer logger.mutex.Unlock()

	// skip log by level
	if checkLevel(logger.level, _ERROR) {
		return
	}

	for _, loggerItem := range logger.loggers {
		loggerItem.Error(args...)
	}

	for _, customLogger := range logger.customLoggers {
		logger.line.AddErrorCustom(customLogger.Error, args...)
	}
}

func (logger *logger) Errorf(format string, args ...any) {
	logger.mutex.Lock()
	defer logger.mutex.Unlock()

	// skip log by level
	if checkLevel(logger.level, _ERROR) {
		return
	}

	for _, loggerItem := range logger.loggers {
		loggerItem.Errorf(format, args...)
	}

	for _, customLogger := range logger.customLoggers {
		logger.line.AddErrorFormatCustom(customLogger.Errorf, format, args...)
	}
}

func (logger *logger) Fatal(args ...any) {
	logger.mutex.Lock()
	defer logger.mutex.Unlock()

	for _, loggerItem := range logger.loggers {
		loggerItem.Fatal(args...)
	}

	for _, customLogger := range logger.customLoggers {
		logger.line.AddErrorCustom(customLogger.Fatal, args...)
	}

	if logger.isCustomDuration {
		time.Sleep(logger.customDuration)
	} else {
		time.Sleep(logger.exitDuration)
	}

	// todo: implement graceful shutdown for .Fatal methods
	os.Exit(1)
}

func (logger *logger) Fatalf(format string, args ...any) {
	logger.mutex.Lock()
	defer logger.mutex.Unlock()

	for _, loggerItem := range logger.loggers {
		loggerItem.Fatalf(format, args...)
	}

	for _, customLogger := range logger.customLoggers {
		logger.line.AddErrorFormatCustom(customLogger.Fatalf, format, args...)
	}

	if logger.isCustomDuration {
		time.Sleep(logger.customDuration)
	} else {
		time.Sleep(logger.exitDuration)
	}

	os.Exit(1)
}
