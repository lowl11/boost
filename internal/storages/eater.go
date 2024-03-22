package storages

import (
	"github.com/lowl11/boost/pkg/io/flex"
	"strings"
)

func Eat(entity any) (string, string, []string) {
	var tableName string
	var aliasName string

	fs, _ := flex.Struct(entity)
	for _, f := range fs.Fields() {
		tags := f.Tag("ef")
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

	ent, _ := flex.Struct(entity)
	fields := ent.FieldsRow()
	columns := make([]string, 0, len(fields))
	for _, field := range fields {
		columns = append(columns, field.Tag("db")[0])
	}

	return tableName, aliasName, columns
}
