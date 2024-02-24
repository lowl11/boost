package sql

import "strings"

type TruncateTableBuilder interface {
	Query

	Table(tableName string) TruncateTableBuilder
}

type truncateTableBuilder struct {
	tableName string
}

func newTruncateTableBuilder(tableName ...string) TruncateTableBuilder {
	builder := &truncateTableBuilder{}

	if len(tableName) > 0 {
		builder.tableName = tableName[0]
	}

	return builder
}

func (builder *truncateTableBuilder) String() string {
	query := strings.Builder{}
	query.WriteString("TRUNCATE TABLE ")
	query.WriteString(builder.tableName)
	return query.String()
}

func (builder *truncateTableBuilder) Table(tableName string) TruncateTableBuilder {
	builder.tableName = tableName
	return builder
}
