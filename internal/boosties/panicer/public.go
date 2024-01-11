package panicer

import (
	"fmt"
	"github.com/lowl11/boost/data/errors"
	"runtime"
	"strings"
)

func Handle(err any) error {
	if err == nil {
		return nil
	}

	parsedError := fromAny(err)
	return errors.
		New("PANIC RECOVER: "+parsedError).
		SetType("PanicError").
		AddContext("trace", getStackTrace())
}

func StackTrace() []string {
	return getStackTrace()
}

func getStackTrace() []string {
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
