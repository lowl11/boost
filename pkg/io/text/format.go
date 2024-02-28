package text

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"regexp"
	"strings"
	"unicode"
)

var (
	_manySpacesReg = regexp.MustCompile("\\s{2,}")
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

// Username makes given text to username standard.
// Example:
//
// Input: John Smith
// Output: john_smith
func Username(username string) string {
	username = strings.TrimSpace(username)
	username = strings.ToLower(username)
	username = _manySpacesReg.ReplaceAllString(username, " ")
	username = strings.ReplaceAll(username, " ", "_")
	username = OnlyLetter(username)
	return username
}

// OnlyLetter keeps only latin letters & underscore ("_")
// Example:
//
// Input: john_smith_jr.
// Output: john_smith_jr
func OnlyLetter(text string) string {
	var output []rune
	for _, char := range text {
		if unicode.IsLetter(char) || char == '_' {
			output = append(output, char)
		}
	}
	return string(output)
}
