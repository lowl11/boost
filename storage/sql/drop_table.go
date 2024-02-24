package sql

import "strings"

type DropTableBuilder interface {
	Query

	Table(tableName string) DropTableBuilder
}

type dropTableBuilder struct {
	tableName string
}

func newDropTableBuilder(tableName ...string) DropTableBuilder {
	builder := &dropTableBuilder{}

	if len(tableName) > 0 {
		builder.tableName = tableName[0]
	}

	return builder
}

func (builder *dropTableBuilder) String() string {
	query := strings.Builder{}
	query.WriteString("DROP TABLE ")
	query.WriteString(builder.tableName)
	return query.String()
}

func (builder *dropTableBuilder) Table(tableName string) DropTableBuilder {
	builder.tableName = tableName
	return builder
}
