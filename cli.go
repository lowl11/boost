package boost

import (
	"os"
	"strings"
)

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
