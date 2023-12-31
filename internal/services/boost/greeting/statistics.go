package greeting

import (
	"github.com/lowl11/boost/internal/boosties/static_controller"
	"github.com/lowl11/boost/internal/helpers/type_helper"
	"github.com/lowl11/boost/internal/services/boost/greeting/printer"
	"os"
	"strings"
)

func (greeting *Greeting) getHttpStatistic() string {
	const (
		startLine = " │ "
		endLine   = " │\n"

		spaces = 17
	)

	builder := strings.Builder{}

	// first line
	routes := type_helper.ToString(greeting.counter.GetRoutes()-static_controller.RouteCount, false)
	groups := type_helper.ToString(greeting.counter.GetGroups(), false)

	builder.WriteString(printer.Color(startLine, greeting.getMainColor()))
	builder.WriteString("Routes: ........")

	builder.WriteString(printer.Color(routes, greeting.getSpecificColor()))

	builder.WriteString(printer.Spaces(spaces - len(routes) - len(groups)))
	builder.WriteString("Groups: ........")
	builder.WriteString(printer.Color(groups, greeting.getSpecificColor()))
	builder.WriteString(printer.Color(endLine, greeting.getMainColor()))

	// second line
	port := greeting.ctx.Port
	pid := type_helper.ToString(os.Getpid(), false)

	builder.WriteString(printer.Color(startLine, greeting.getMainColor()))
	builder.WriteString("Port: ..........")

	builder.WriteString(printer.Color(port, greeting.getSpecificColor()))

	builder.WriteString(printer.Spaces(spaces - len(port) - len(pid)))
	builder.WriteString("PID: ...........")
	builder.WriteString(printer.Color(pid, greeting.getSpecificColor()))
	builder.WriteString(printer.Color(endLine, greeting.getMainColor()))

	return builder.String()
}

func (greeting *Greeting) getRPCStatistic() string {
	const (
		startLine = " │ "
		endLine   = " │\n"

		spaces = 17
	)

	builder := strings.Builder{}

	// first line
	port := greeting.ctx.Port
	pid := type_helper.ToString(os.Getpid(), false)

	builder.WriteString(printer.Color(startLine, greeting.getMainColor()))
	builder.WriteString("Port: ..........")

	builder.WriteString(printer.Color(port, greeting.getSpecificColor()))

	builder.WriteString(printer.Spaces(spaces - len(port) - len(pid)))
	builder.WriteString("PID: ...........")
	builder.WriteString(printer.Color(pid, greeting.getSpecificColor()))
	builder.WriteString(printer.Color(endLine, greeting.getMainColor()))

	return builder.String()
}

func (greeting *Greeting) getCronStatistic() string {
	const (
		startLine = " │ "
		endLine   = " │\n"

		spaces = 17
	)

	builder := strings.Builder{}

	// first line
	actions := type_helper.ToString(greeting.counter.GetCronActions(), false)
	pid := type_helper.ToString(os.Getpid(), false)

	builder.WriteString(printer.Color(startLine, greeting.getMainColor()))
	builder.WriteString("Actions: ........")

	builder.WriteString(printer.Color(actions, greeting.getSpecificColor()))

	builder.WriteString(printer.Spaces(spaces - len(actions) - len(pid) - 1))
	builder.WriteString("PID: ...........")
	builder.WriteString(printer.Color(pid, greeting.getSpecificColor()))
	builder.WriteString(printer.Color(endLine, greeting.getMainColor()))

	return builder.String()
}

func (greeting *Greeting) getListenerStatistic() string {
	const (
		startLine = " │ "
		endLine   = " │\n"

		spaces = 19
	)

	builder := strings.Builder{}

	// first line
	binds := type_helper.ToString(greeting.counter.GetListenerBind(), false)
	pid := type_helper.ToString(os.Getpid(), false)

	builder.WriteString(printer.Color(startLine, greeting.getMainColor()))
	builder.WriteString("Binds: ........")

	builder.WriteString(printer.Color(binds, greeting.getSpecificColor()))

	builder.WriteString(printer.Spaces(spaces - len(binds) - len(pid) - 1))
	builder.WriteString("PID: ...........")
	builder.WriteString(printer.Color(pid, greeting.getSpecificColor()))
	builder.WriteString(printer.Color(endLine, greeting.getMainColor()))

	return builder.String()
}
