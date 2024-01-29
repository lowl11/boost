package exception

import (
	"fmt"
	"runtime"
	"strings"
)

func StackTrace() []string {
	stackTrace := make([]string, 0, 10)

	var file string
	var line int
	var callerCatch bool
	callerIterator := 4
	for {
		_, file, line, callerCatch = runtime.Caller(callerIterator)
		if !callerCatch || strings.Contains(file, "github.com/valyala/fasthttp") {
			break
		}

		builder := strings.Builder{}
		builder.Grow(len(file) + 11)
		_, _ = fmt.Fprintf(&builder, "%s: line %d", file, line)
		stackTrace = append(stackTrace, builder.String())
		callerIterator++
	}

	return stackTrace
}
