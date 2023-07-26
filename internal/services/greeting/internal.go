package greeting

import (
	"github.com/lowl11/boost/internal/boosties/static_controller"
	"github.com/lowl11/boost/internal/services/greeting/printer"
	"github.com/lowl11/boost/pkg/enums/colors"
)

const (
	header = " ┌───────────────────────────────────────────────────┐"
	footer = " └───────────────────────────────────────────────────┘\n"

	logo = `
 │            | |__   ___   ___  ___| |_             │
 │            | '_ \ / _ \ / _ \/ __| __|            │
 │            | |_) | (_) | (_) \__ \ |_             │
 │            |_.__/ \___/ \___/|___/\__|            │
`
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

func (greeting *Greeting) appendHeader() {
	greeting.message += printer.Color(header, greeting.getMainColor())
}

func (greeting *Greeting) appendFooter() {
	greeting.message += printer.Color(footer, greeting.getMainColor())
}

func (greeting *Greeting) appendLogo() {
	greeting.message += printer.Color(logo, greeting.getMainColor())
}

func (greeting *Greeting) appendDescription() {
	greeting.message += printer.Color(description, greeting.getMainColor())
}

func (greeting *Greeting) printStatistic() {
	greeting.message += printer.Build(
		statistics,
		printer.Color(greeting.counter.GetRoutes()-static_controller.RouteCount, colors.Cyan),
		printer.Color(greeting.counter.GetGroups(), colors.Cyan),
	)
}

func (greeting *Greeting) getMainColor() string {
	if greeting.mainColor == "" {
		greeting.mainColor = colors.White
	}

	return greeting.mainColor
}
