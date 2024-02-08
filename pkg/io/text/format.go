package text

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strings"
	"unicode"
)

// Title makes text first word starts with upper case, other words with lower case.
// Example:
//
//	Input: HELLO WORLD
//	Output: Hello world
func Title(text string) string {
	if text == "" {
		return ""
	}

	text = strings.ToLower(text)
	runes := []rune(text)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

// EveryTitle makes every word start with uppercase
// Example:
//
//	Input: HELLO WORLD
//	Output: Hello World
func EveryTitle(text string) string {
	return cases.Title(language.Und).String(strings.ToLower(text))
}
