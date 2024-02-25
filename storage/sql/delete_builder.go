package sql

import (
	"context"
	"strings"
)

type DeleteBuilder interface {
	Query

	Exec(ctx context.Context, args ...any) error

	From(tableName string) DeleteBuilder
	Where(func(Where)) DeleteBuilder
}

type deleteBuilder struct {
	tableName string
	where     Where
}

func newDeleteBuilder(tableName ...string) *deleteBuilder {
	builder := &deleteBuilder{
		where: newWhere(),
	}

	if len(tableName) > 0 {
		builder.tableName = tableName[0]
	}

	return builder
}

func (builder *deleteBuilder) String() string {
	// builder
	query := strings.Builder{}
	query.Grow(300)

	// delete
	query.WriteString("DELETE FROM ")
	query.WriteString(builder.tableName)
	query.WriteString("\n")

	// where
	whereClause := builder.where.(Query).String()
	if len(whereClause) != 0 {
		query.WriteString("WHERE \n\t")
		query.WriteString(whereClause)
		query.WriteString("\n")
	}

	return query.String()
}

func (builder *deleteBuilder) From(tableName string) DeleteBuilder {
	builder.tableName = tableName
	return builder
}

func (builder *deleteBuilder) Where(whereFunc func(builder Where)) DeleteBuilder {
	return builder.applyWhere(whereFunc)
}

func (builder *deleteBuilder) applyWhere(whereFunc func(builder Where)) DeleteBuilder {
	whereFunc(builder.where)
	return builder
}

func (builder *deleteBuilder) Exec(ctx context.Context, args ...any) error {
	return newParamExecutor(ctx, builder.String()).Exec(args...)
}
