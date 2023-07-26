package greeting

func (greeting *Greeting) Print() {
	greeting.appendHeader()

	greeting.appendLogo()
	greeting.appendDescription()
	greeting.appendStatistic()

	greeting.appendFooter()

	greeting.printMessage()
}

func (greeting *Greeting) MainColor(color string) *Greeting {
	if color == "" {
		return greeting
	}

	greeting.mainColor = color
	return greeting
}

func (greeting *Greeting) SpecificColor(color string) *Greeting {
	if color == "" {
		return greeting
	}

	greeting.specificColor = color
	return greeting
}
