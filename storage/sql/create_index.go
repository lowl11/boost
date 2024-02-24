package sql

import "strings"

type CreateIndexBuilder interface {
	Query

	IfNotExist() CreateIndexBuilder
	Unique() CreateIndexBuilder
	Name(name string) CreateIndexBuilder
	TableColumns(tableName string, columns ...string) CreateIndexBuilder
}

type createIndexBuilder struct {
	name       string
	table      string
	unique     bool
	columns    []string
	ifNotExist bool
}

func newCreateIndexBuilder(name ...string) CreateIndexBuilder {
	builder := &createIndexBuilder{
		columns: []string{},
	}

	if len(name) > 0 {
		builder.name = name[0]
	}

	return builder
}

func (builder *createIndexBuilder) String() string {
	query := strings.Builder{}

	if builder.unique {
		query.WriteString("CREATE UNIQUE INDEX ")
	} else {
		query.WriteString("CREATE INDEX ")
	}

	if builder.ifNotExist {
		query.WriteString("IF NOT EXISTS ")
	}

	query.WriteString(builder.name)
	query.WriteString("\n")
	query.WriteString("ON ")
	query.WriteString(builder.table)
	query.WriteString(" (")
	for index, column := range builder.columns {
		query.WriteString(column)

		if index < len(builder.columns)-1 {
			query.WriteString(", ")
		}
	}
	query.WriteString(")")

	return query.String()
}

func (builder *createIndexBuilder) IfNotExist() CreateIndexBuilder {
	builder.ifNotExist = true
	return builder
}

func (builder *createIndexBuilder) Unique() CreateIndexBuilder {
	builder.unique = true
	return builder
}

func (builder *createIndexBuilder) Name(name string) CreateIndexBuilder {
	builder.name = name
	return builder
}

func (builder *createIndexBuilder) TableColumns(tableName string, columns ...string) CreateIndexBuilder {
	builder.table = tableName
	builder.columns = columns
	return builder
}
