package boost

import (
	"os"
	"strings"
)

// Flag returns value of the given flag name.
// Function search combination like "key=value"
func Flag(value string) string {
	for _, arg := range os.Args {
		before, after, found := strings.Cut(arg, "=")
		if !found {
			continue
		}

		if before == value {
			return after
		}
	}

	return ""
}

// FlagExist checks exist flag or not
func FlagExist(value string) bool {
	if Flag(value) != "" {
		return true
	}

	for _, arg := range os.Args {
		if arg == value {
			return true
		}
	}

	return false
}
