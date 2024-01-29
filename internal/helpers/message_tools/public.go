package message_tools

import (
	"encoding/json"
	"strings"
)

var (
	JsonMode     bool
	NoTimeMode   bool
	NoPrefixMode bool
)

func Build(args ...any) string {
	if len(args) == 0 {
		return ""
	}

	stringArgs := strings.Builder{}
	for _, arg := range args {
		stringArgs.WriteString(toString(arg))
		stringArgs.WriteString(" ")
	}
	return stringArgs.String()[:stringArgs.Len()-1]
}

func BuildPrefix(level string) string {
	prefix := strings.Builder{}
	if NoPrefixMode && !NoTimeMode {
		prefix.WriteString(getTime())
		prefix.WriteString(" ")
		return prefix.String()
	}

	if NoPrefixMode {
		return ""
	}

	if NoTimeMode {
		return level
	}

	prefix.WriteString(getTime())
	prefix.WriteString(" ")
	prefix.WriteString(level)
	return prefix.String()
}

func Json(level string, args ...any) string {
	logMessage := &LogMessage{
		Message: Build(args...),
		Level:   level,
		Time:    getTime(),
	}

	logMessageInBytes, err := json.Marshal(logMessage)
	if err != nil {
		return "|ERROR IN BUILDING MESSAGE|"
	}

	return string(logMessageInBytes)
}
