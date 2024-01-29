package greeting

import (
	"fmt"
	"github.com/lowl11/boost/data/enums/colors"
	"github.com/lowl11/boost/pkg/system/types"
	"os"
	"strings"
)

func _print(text string, args ...any) {
	_, _ = fmt.Fprintf(os.Stdout, text, args...)
}

func build(text string, args ...any) string {
	builder := strings.Builder{}
	_, _ = fmt.Fprintf(&builder, text, args...)
	return builder.String()
}

func color(text any, color string) string {
	textInString := types.ToString(text)

	coloredText := strings.Builder{}
	coloredText.WriteString(color)
	coloredText.WriteString(textInString)
	coloredText.WriteString(colors.Reset)
	return coloredText.String()
}

func spaces(count int) string {
	const space = " "

	var spaces string
	for i := 0; i < count; i++ {
		spaces += space
	}

	return spaces
}
