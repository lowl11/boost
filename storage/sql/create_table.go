package sql

import "strings"

type CreateTableBuilder interface {
	Query

	Table(tableName string) CreateTableBuilder
	IfNotExist() CreateTableBuilder
	Column(columns ...Column) CreateTableBuilder
	PartitionBy(columnNames ...string) CreateTableBuilder

	Sql(sql string) CreateTableBuilder
}

type createTableBuilder struct {
	tableName        string
	columns          []Column
	ifNotExist       bool
	partitionColumns []string

	sql string
}

func newCreateTable(tableName ...string) CreateTableBuilder {
	builder := &createTableBuilder{
		columns:          make([]Column, 0, 10),
		partitionColumns: make([]string, 0),
	}

	if len(tableName) > 0 {
		builder.tableName = tableName[0]
	}

	return builder
}

func (builder *createTableBuilder) String() string {
	query := strings.Builder{}
	query.Grow(500)
	query.WriteString("CREATE TABLE ")
	if builder.ifNotExist {
		query.WriteString("IF NOT EXISTS ")
	}
	query.WriteString(builder.tableName)
	query.WriteString(" (\n")

	for index, column := range builder.columns {
		query.WriteString("\t")
		query.WriteString(column.String())
		if index < len(builder.columns)-1 {
			query.WriteString(",\n")
		}
	}

	query.WriteString("\n)")
	if len(builder.partitionColumns) > 0 {
		query.WriteString("\npartition by LIST (")
		for index, col := range builder.partitionColumns {
			query.WriteString(col)
			if index < len(builder.partitionColumns)-1 {
				query.WriteString(", ")
			}
		}
		query.WriteString(")")
	}
	return query.String()
}

func (builder *createTableBuilder) Table(tableName string) CreateTableBuilder {
	builder.tableName = tableName
	return builder
}

func (builder *createTableBuilder) IfNotExist() CreateTableBuilder {
	builder.ifNotExist = true
	return builder
}

func (builder *createTableBuilder) Column(columns ...Column) CreateTableBuilder {
	builder.columns = append(builder.columns, columns...)
	return builder
}

func (builder *createTableBuilder) PartitionBy(columnNames ...string) CreateTableBuilder {
	builder.partitionColumns = columnNames
	return builder
}

func (builder *createTableBuilder) Sql(sql string) CreateTableBuilder {
	builder.sql = sql
	return builder
}
