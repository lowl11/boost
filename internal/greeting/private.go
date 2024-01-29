package greeting

import (
	"github.com/lowl11/boost/data/enums/modes"
	"strings"
)

const (
	header = " ┌───────────────────────────────────────────────────┐"
	footer = " └───────────────────────────────────────────────────┘\n"
)

func (greeting *Greeting) printMessage() {
	if greeting.printed {
		return
	}

	_print(greeting.message)
	greeting.printed = true
}

func (greeting *Greeting) getMainColor() string {
	return greeting.mainColor
}

func (greeting *Greeting) getSpecificColor() string {
	return greeting.specificColor
}

func (greeting *Greeting) appendHeader() {
	greeting.message += color(header, greeting.getMainColor())
}

func (greeting *Greeting) appendFooter() {
	greeting.message += color(footer, greeting.getMainColor())
}

func (greeting *Greeting) appendLogo() {
	greeting.message += greeting.getLogo()
}

func (greeting *Greeting) appendMode() {
	const (
		startLine = " │ "
		endLine   = " │\n"

		beforeSpaces = 18
		afterSpaces  = 24
	)

	modeLength := len(greeting.ctx.Mode)

	builder := strings.Builder{}
	builder.WriteString(color(startLine, greeting.getMainColor()))
	builder.WriteString(spaces(beforeSpaces))
	builder.WriteString("Mode: ")
	builder.WriteString(color(greeting.ctx.Mode, greeting.getSpecificColor()))
	builder.WriteString(spaces(afterSpaces - modeLength + 1))
	builder.WriteString(color(endLine, greeting.getMainColor()))

	greeting.message += builder.String()
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
	builder.WriteString(color(startLine, greeting.getMainColor()))
	builder.WriteString(color(firstLine, greeting.getSpecificColor()))
	builder.WriteString(color(" "+endLine, greeting.getMainColor()))

	// second line
	builder.WriteString(color(startLine, greeting.getMainColor()))
	builder.WriteString(color(secondLine, greeting.getSpecificColor()))
	builder.WriteString(color(endLine, greeting.getMainColor()))

	// third line
	builder.WriteString(color(startLine, greeting.getMainColor()))
	builder.WriteString(color(thirdLine, greeting.getSpecificColor()))
	builder.WriteString(color(" "+endLine, greeting.getMainColor()))

	// fourth line
	builder.WriteString(color(startLine, greeting.getMainColor()))
	builder.WriteString(color(fourthLine, greeting.getSpecificColor()))
	builder.WriteString(color(endLine, greeting.getMainColor()))

	return builder.String()
}

func (greeting *Greeting) getStatistic() string {
	switch greeting.ctx.Mode {
	case modes.Http:
		return greeting.getHttpStatistic()
	case modes.RPC:
		return greeting.getRPCStatistic()
	case modes.Cron:
		return greeting.getCronStatistic()
	case modes.Listener:
		return greeting.getListenerStatistic()
	}

	return ""
}
