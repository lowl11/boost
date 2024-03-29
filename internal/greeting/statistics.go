package greeting

import (
	"github.com/lowl11/boost/pkg/io/types"
	"os"
	"strings"
)

func (greeting *Greeting) getHttpStatistic() string {
	const (
		startLine = " │ "
		endLine   = " │\n"

		_spaces                     = 17
		staticControllerRoutesCount = 2
	)

	builder := strings.Builder{}

	// first line
	routes := types.String(greeting.counter.GetRoutes() - staticControllerRoutesCount)
	groups := types.String(greeting.counter.GetGroups())

	builder.WriteString(color(startLine, greeting.getMainColor()))
	builder.WriteString("Routes: ........")

	builder.WriteString(color(routes, greeting.getSpecificColor()))

	builder.WriteString(spaces(_spaces - len(routes) - len(groups)))
	builder.WriteString("Groups: ........")
	builder.WriteString(color(groups, greeting.getSpecificColor()))
	builder.WriteString(color(endLine, greeting.getMainColor()))

	// second line
	port := greeting.ctx.Port
	pid := types.String(os.Getpid())

	builder.WriteString(color(startLine, greeting.getMainColor()))
	builder.WriteString("Port: ..........")

	builder.WriteString(color(port, greeting.getSpecificColor()))

	builder.WriteString(spaces(_spaces - len(port) - len(pid)))
	builder.WriteString("PID: ...........")
	builder.WriteString(color(pid, greeting.getSpecificColor()))
	builder.WriteString(color(endLine, greeting.getMainColor()))

	return builder.String()
}

func (greeting *Greeting) getRPCStatistic() string {
	const (
		startLine = " │ "
		endLine   = " │\n"

		_spaces = 17
	)

	builder := strings.Builder{}

	// first line
	port := greeting.ctx.Port
	pid := types.String(os.Getpid())

	builder.WriteString(color(startLine, greeting.getMainColor()))
	builder.WriteString("Port: ..........")

	builder.WriteString(color(port, greeting.getSpecificColor()))

	builder.WriteString(spaces(_spaces - len(port) - len(pid)))
	builder.WriteString("PID: ...........")
	builder.WriteString(color(pid, greeting.getSpecificColor()))
	builder.WriteString(color(endLine, greeting.getMainColor()))

	return builder.String()
}

func (greeting *Greeting) getCronStatistic() string {
	const (
		startLine = " │ "
		endLine   = " │\n"

		_spaces = 17
	)

	builder := strings.Builder{}

	// first line
	actions := types.String(greeting.counter.GetCronActions())
	pid := types.String(os.Getpid())

	builder.WriteString(color(startLine, greeting.getMainColor()))
	builder.WriteString("Actions: ........")

	builder.WriteString(color(actions, greeting.getSpecificColor()))

	builder.WriteString(spaces(_spaces - len(actions) - len(pid) - 1))
	builder.WriteString("PID: ...........")
	builder.WriteString(color(pid, greeting.getSpecificColor()))
	builder.WriteString(color(endLine, greeting.getMainColor()))

	return builder.String()
}

func (greeting *Greeting) getListenerStatistic() string {
	const (
		startLine = " │ "
		endLine   = " │\n"

		_spaces = 19
	)

	builder := strings.Builder{}

	// first line
	binds := types.String(greeting.counter.GetListenerBind())
	pid := types.String(os.Getpid())

	builder.WriteString(color(startLine, greeting.getMainColor()))
	builder.WriteString("Binds: ........")

	builder.WriteString(color(binds, greeting.getSpecificColor()))

	builder.WriteString(spaces(_spaces - len(binds) - len(pid) - 1))
	builder.WriteString("PID: ...........")
	builder.WriteString(color(pid, greeting.getSpecificColor()))
	builder.WriteString(color(endLine, greeting.getMainColor()))

	return builder.String()
}
