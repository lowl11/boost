package greeting

import "github.com/lowl11/boost/pkg/enums/colors"

func (greeting *Greeting) Print() {
	greeting.appendHeader()

	greeting.appendLogo()
	greeting.appendDescription()
	greeting.printStatistic()

	greeting.appendFooter()

	greeting.printMessage()
}

func (greeting *Greeting) MainColor(color string) {
	if color == "" {
		greeting.mainColor = colors.White
	}

	greeting.mainColor = color
}
