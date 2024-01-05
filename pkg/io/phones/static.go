package phones

import (
	"errors"
	"regexp"
	"unicode/utf8"
)

var (
	removeOddSymbols = regexp.MustCompile("[^+\\d]")
)

func CheckPhone(phone string) error {
	// check length (min: 10, max: 20)
	length := utf8.RuneCountInString(phone)
	if length < 10 || length > 20 {
		return errors.New("does not match phone pattern")
	}

	return nil
}

func MaskPhone(phone string) string {
	phone = removeOddSymbols.ReplaceAllString(phone, "")

	length := utf8.RuneCountInString(phone)
	switch length {
	case 10:
		phone = "+7" + phone
	case 11:
		if phone[0] == '8' {
			phone = "+7" + trimFirstRune(phone)
		}
	}

	return phone
}
