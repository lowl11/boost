package sql

import "strings"

type DropIndexBuilder interface {
	Query

	SQL(sql string) DropIndexBuilder
	Name(name string) DropIndexBuilder
	Table(tableName string) DropIndexBuilder
}

type dropIndexBuilder struct {
	sql   string
	name  string
	table string
}

func newDropIndexBuilder(name ...string) DropIndexBuilder {
	builder := &dropIndexBuilder{
		sql: "Postgres",
	}

	if len(name) > 0 {
		builder.name = name[0]
	}

	return builder
}

func (builder *dropIndexBuilder) String() string {
	query := strings.Builder{}

	if builder.sql == "MySQL" {
		query.WriteString("ALTER TABLE ")
		query.WriteString(builder.table)
		query.WriteString("\n")
	}

	query.WriteString("DROP INDEX ")

	switch builder.sql {
	case "Postgres":
		query.WriteString(builder.table)
		query.WriteString(".")
		query.WriteString("\"")
		query.WriteString(builder.name)
		query.WriteString("\"")
	case "MSSQL":
		query.WriteString(builder.name)
		query.WriteString(" ON ")
		query.WriteString(builder.table)
	case "MySQL":
		query.WriteString(builder.name)
	}

	return query.String()
}

func (builder *dropIndexBuilder) SQL(sql string) DropIndexBuilder {
	builder.sql = sql
	return builder
}

func (builder *dropIndexBuilder) Name(name string) DropIndexBuilder {
	builder.name = name
	return builder
}

func (builder *dropIndexBuilder) Table(tableName string) DropIndexBuilder {
	builder.table = tableName
	return builder
}
