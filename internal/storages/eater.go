package storages

import (
	"github.com/lowl11/flex"
	"strings"
)

func Eat[T any](entity T) (string, string, []string) {
	var tableName string
	var aliasName string

	fs := flex.Struct(entity).Fields()
	for _, f := range fs {
		tags := flex.Field(f).Tag("ef")
		if len(tags) == 0 {
			continue
		}

		for _, tag := range tags {
			before, after, found := strings.Cut(tag, ":")
			if !found {
				continue
			}

			switch before {
			case "table":
				tableName = after
			case "alias":
				aliasName = after
			}
		}

		if len(tableName) > 0 && len(aliasName) > 0 {
			break
		}
	}

	fields := flex.Struct(entity).FieldsRow()
	columns := make([]string, 0, len(fields))
	for _, field := range fields {
		columns = append(columns, flex.Field(field).Tag("db")[0])
	}

	return tableName, aliasName, columns
}
