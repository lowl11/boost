package phones

import (
	"errors"
	"regexp"
	"unicode/utf8"
)

var (
	removeOddSymbols = regexp.MustCompile("[^+\\d]")
)

func Check(phone string) error {
	// check length (min: 10, max: 20)
	length := utf8.RuneCountInString(phone)
	if length < 10 || length > 20 {
		return errors.New("does not match phone pattern")
	}

	// todo: check for code

	return nil
}

func Is(phone string) bool {
	return Check(Mask(phone)) == nil
}

func Mask(phone string) string {
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
