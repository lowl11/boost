package sql

import "strings"

type AlterTableBuilder interface {
	Query

	Table(tableName string) AlterTableBuilder
	SQL(sql string) AlterTableBuilder
	AddColumn(column string) AlterTableBuilder
	DropColumn(column string) AlterTableBuilder
	RenameColumn(column, newName string) AlterTableBuilder
	AlterColumn(column string) AlterTableBuilder
	Set(string) AlterTableBuilder
	Type(DataType) AlterTableBuilder
	Reset() AlterTableBuilder
	Restart() AlterTableBuilder
	Add() AlterTableBuilder
	Drop(string) AlterTableBuilder
}

type alterTableBuilder struct {
	tableName string
	mode      string
	column    string
	newName   string
	sql       string

	dataType       DataType
	setAttributes  string
	dropAttributes string

	isSet     bool
	isType    bool
	isReset   bool
	isRestart bool
	isAdd     bool
	isDrop    bool
}

func newAlterTableBuilder(tableName ...string) AlterTableBuilder {
	builder := &alterTableBuilder{
		sql: "Postgres",
	}

	if len(tableName) > 0 {
		builder.tableName = tableName[0]
	}

	return builder
}

func (builder *alterTableBuilder) String() string {
	query := strings.Builder{}
	query.WriteString("ALTER TABLE ")
	query.WriteString(builder.tableName)
	query.WriteString("\n")
	query.WriteString(builder.mode)
	query.WriteString(" ")
	query.WriteString(builder.column)

	switch builder.mode {
	case "ADD":
		query.WriteString(" ")
		builder.dataType.Write(builder.sql, &query)
	case "ALTER COLUMN":
		query.WriteString(" ")

		if builder.isSet {
			query.WriteString("SET ")
			query.WriteString(builder.setAttributes)
		} else if builder.isType {
			query.WriteString("TYPE ")
			query.WriteString(builder.dataType.Name())
		} else if builder.isDrop {
			query.WriteString("DROP ")
			query.WriteString(builder.dropAttributes)
		}
	case "RENAME COLUMN":
		query.WriteString(" TO ")
		query.WriteString(builder.newName)
	}

	return query.String()
}

func (builder *alterTableBuilder) Table(tableName string) AlterTableBuilder {
	builder.tableName = tableName
	return builder
}

func (builder *alterTableBuilder) AddColumn(column string) AlterTableBuilder {
	builder.mode = "ADD"
	builder.column = column
	return builder
}

func (builder *alterTableBuilder) DropColumn(column string) AlterTableBuilder {
	builder.mode = "DROP COLUMN"
	builder.column = column
	return builder
}

func (builder *alterTableBuilder) RenameColumn(column, newName string) AlterTableBuilder {
	builder.mode = "RENAME COLUMN"
	builder.column = column
	builder.newName = newName
	return builder
}

func (builder *alterTableBuilder) AlterColumn(column string) AlterTableBuilder {
	if builder.sql == "MySQL" {
		builder.mode = "MODIFY COLUMN"
	} else {
		builder.mode = "ALTER COLUMN"
	}
	builder.column = column
	return builder
}

func (builder *alterTableBuilder) Set(attributes string) AlterTableBuilder {
	builder.isSet = true
	builder.setAttributes = attributes
	return builder
}

func (builder *alterTableBuilder) Type(dt DataType) AlterTableBuilder {
	builder.isType = true
	builder.dataType = dt
	return builder
}

func (builder *alterTableBuilder) Add() AlterTableBuilder {
	builder.isAdd = true
	return builder
}

func (builder *alterTableBuilder) Drop(attributes string) AlterTableBuilder {
	builder.isDrop = true
	builder.dropAttributes = attributes
	return builder
}

func (builder *alterTableBuilder) Reset() AlterTableBuilder {
	builder.isReset = true
	return builder
}

func (builder *alterTableBuilder) Restart() AlterTableBuilder {
	builder.isRestart = true
	return builder
}

func (builder *alterTableBuilder) SQL(sql string) AlterTableBuilder {
	builder.sql = sql
	return builder
}
