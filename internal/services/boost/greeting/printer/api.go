package printer

import (
	"fmt"
	"github.com/lowl11/boost/data/enums/colors"
	"github.com/lowl11/boost/internal/helpers/type_helper"
	"os"
	"strings"
)

func Print(text string, args ...any) {
	_, _ = fmt.Fprintf(os.Stdout, text, args...)
}

func Build(text string, args ...any) string {
	builder := strings.Builder{}
	_, _ = fmt.Fprintf(&builder, text, args...)
	return builder.String()
}

func Color(text any, color string) string {
	textInString := type_helper.ToString(text, false)

	coloredText := strings.Builder{}
	coloredText.WriteString(color)
	coloredText.WriteString(textInString)
	coloredText.WriteString(colors.Reset)
	return coloredText.String()
}

func Spaces(count int) string {
	const space = " "

	var spaces string
	for i := 0; i < count; i++ {
		spaces += space
	}

	return spaces
}
