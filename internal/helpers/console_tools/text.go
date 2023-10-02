package console_tools

import "strings"

func Debug(text string) string {
	return color(purple, text)
}

func Info(text string) string {
	return color(green, text)
}

func Warn(text string) string {
	return color(yellow, text)
}

func Error(text string) string {
	return color(red, text)
}

func Fatal(text string) string {
	return color(gray, text)
}

func color(color, text string) string {
	coloredText := strings.Builder{}
	coloredText.WriteString(color)
	coloredText.WriteString(text)
	coloredText.WriteString(reset)
	return coloredText.String()
}
