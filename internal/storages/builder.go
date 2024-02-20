package storages

import (
	"fmt"
	"strings"
)

func MakeName(name string) string {
	query := strings.Builder{}
	query.Grow(len(name) + 2)
	if strings.Contains(name, ".") {
		before, after, _ := strings.Cut(name, ".")
		_, _ = fmt.Fprintf(&query, "\"%s\".\"%s\"", before, after)
	} else {
		_, _ = fmt.Fprintf(&query, "\"%s\"", name)
	}
	return query.String()
}
