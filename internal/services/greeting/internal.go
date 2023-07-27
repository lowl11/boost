package greeting

import (
	"github.com/lowl11/boost/internal/boosties/static_controller"
	"github.com/lowl11/boost/internal/helpers/type_helper"
	"github.com/lowl11/boost/internal/services/greeting/printer"
	"strings"
)

const (
	header = " ┌───────────────────────────────────────────────────┐"
	footer = " └───────────────────────────────────────────────────┘\n"

	description = ` │     Minimalist Go framework based on FastHTTP     │
 │          https://github.com/lowl11/boost          │
`

	statistics = ` │ Routes: ........%s                                 │
 │ Groups: ........%s                                 │
`
)

func (greeting *Greeting) printMessage() {
	printer.Print(greeting.message)
}

func (greeting *Greeting) getMainColor() string {
	return greeting.mainColor
}

func (greeting *Greeting) getSpecificColor() string {
	return greeting.specificColor
}

func (greeting *Greeting) appendHeader() {
	greeting.message += printer.Color(header, greeting.getMainColor())
}

func (greeting *Greeting) appendFooter() {
	greeting.message += printer.Color(footer, greeting.getMainColor())
}

func (greeting *Greeting) appendLogo() {
	greeting.message += greeting.getLogo()
}

func (greeting *Greeting) appendDescription() {
	greeting.message += printer.Color(description, greeting.getMainColor())
}

func (greeting *Greeting) appendStatistic() {
	greeting.message += greeting.getStatistic()
}

func (greeting *Greeting) getLogo() string {
	const (
		startLine = " │            "
		endLine   = "            │\n"

		firstLine  = "| |__   ___   ___  ___| |_"
		secondLine = "| '_ \\ / _ \\ / _ \\/ __| __|"
		thirdLine  = "| |_) | (_) | (_) \\__ \\ |_"
		fourthLine = "|_.__/ \\___/ \\___/|___/\\__|"
	)

	builder := strings.Builder{}
	builder.WriteString("\n")

	// first line
	builder.WriteString(printer.Color(startLine, greeting.getMainColor()))
	builder.WriteString(printer.Color(firstLine, greeting.getSpecificColor()))
	builder.WriteString(printer.Color(" "+endLine, greeting.getMainColor()))

	// second line
	builder.WriteString(printer.Color(startLine, greeting.getMainColor()))
	builder.WriteString(printer.Color(secondLine, greeting.getSpecificColor()))
	builder.WriteString(printer.Color(endLine, greeting.getMainColor()))

	// third line
	builder.WriteString(printer.Color(startLine, greeting.getMainColor()))
	builder.WriteString(printer.Color(thirdLine, greeting.getSpecificColor()))
	builder.WriteString(printer.Color(" "+endLine, greeting.getMainColor()))

	// fourth line
	builder.WriteString(printer.Color(startLine, greeting.getMainColor()))
	builder.WriteString(printer.Color(fourthLine, greeting.getSpecificColor()))
	builder.WriteString(printer.Color(endLine, greeting.getMainColor()))

	return builder.String()
}

func (greeting *Greeting) getStatistic() string {
	const (
		startLine = " │ "
		endLine   = " │\n"

		betweenSpacesLen = 17
	)

	builder := strings.Builder{}

	// first line
	routes := type_helper.ToString(greeting.counter.GetRoutes() - static_controller.RouteCount)
	groups := type_helper.ToString(greeting.counter.GetGroups())

	builder.WriteString(printer.Color(startLine, greeting.getMainColor()))
	builder.WriteString("Routes: ........")

	builder.WriteString(printer.Color(routes, greeting.getSpecificColor()))

	builder.WriteString(printer.Spaces(betweenSpacesLen - len(routes) - len(groups)))
	builder.WriteString("Groups: ........")
	builder.WriteString(printer.Color(groups, greeting.getSpecificColor()))
	builder.WriteString(printer.Color(endLine, greeting.getMainColor()))

	return builder.String()
}
