package sql

import "strings"

type Column interface {
	Query

	Name(name string) Column
	GetName() string
	DataType(dataType DataType) Column
	GetDataType() DataType
}

type columnBuilder struct {
	name     string
	dataType DataType
}

func newColumnBuilder(name ...string) Column {
	builder := &columnBuilder{}

	if len(name) > 0 {
		builder.name = name[0]
	}

	return builder
}

func (builder columnBuilder) String() string {
	query := strings.Builder{}
	query.WriteString(builder.name)
	query.WriteString(" ")
	builder.dataType.Write(getSQL(), &query)
	return strings.TrimSpace(strings.ReplaceAll(query.String(), "% FIELD_NAME %", builder.name))
}

func (builder columnBuilder) Name(name string) Column {
	builder.name = name
	return builder
}

func (builder columnBuilder) GetName() string {
	return builder.name
}

func (builder columnBuilder) DataType(dataType DataType) Column {
	builder.dataType = dataType
	return builder
}

func (builder columnBuilder) GetDataType() DataType {
	return builder.dataType
}
