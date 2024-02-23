package storages

import (
	"fmt"
	"strings"
)

func BuildWhere(alias, field, sign string, value any) string {
	valueString := ToString(value)

	builder := strings.Builder{}
	if len(alias) > 0 {
		if !strings.Contains(field, ".") {
			_, _ = fmt.Fprintf(&builder, "%s.%s %s %s", alias, field, sign, valueString)
		} else {
			before, after, _ := strings.Cut(field, ".")
			if strings.Contains(alias, before) {
				_, _ = fmt.Fprintf(&builder, "%s.%s %s %s", alias, after, sign, valueString)
			} else {
				_, _ = fmt.Fprintf(&builder, "\"%s\".%s %s %s", before, after, sign, valueString)
			}
		}
	} else {
		_, _ = fmt.Fprintf(&builder, "%s %s %s", field, sign, valueString)
	}

	return builder.String()
}

func BuildWhereBetween(field string, left, right any) string {
	leftString := ToString(left)
	rightString := ToString(right)

	builder := strings.Builder{}
	_, _ = fmt.Fprintf(&builder, "%s BETWEEN %s AND %s", field, leftString, rightString)
	return builder.String()
}

func BuildWhereArray(field, sign, value string) string {
	builder := strings.Builder{}
	_, _ = fmt.Fprintf(&builder, "%s %s %s", field, sign, value)
	return builder.String()
}

func BuildWhereNot(condition string) string {
	builder := strings.Builder{}
	builder.Grow(len(condition) + 5)
	_, _ = fmt.Fprintf(&builder, "NOT(%s)", condition)
	return builder.String()
}

func WhereBrackets(condition string) string {
	builder := strings.Builder{}
	builder.Grow(len(condition) + 5)
	_, _ = fmt.Fprintf(&builder, "(%s)", condition)
	return builder.String()
}
