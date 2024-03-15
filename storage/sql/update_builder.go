package sql

import (
	"context"
	"github.com/lowl11/boost/internal/storages"
	"github.com/lowl11/boost/pkg/io/types"
	"strings"
)

type UpdateBuilder interface {
	Query

	Exec(ctx context.Context, args ...any) error

	GetParam() (string, bool)
	From(tableName string) UpdateBuilder
	Set(pairs ...Pair) UpdateBuilder
	Where(func(Where)) UpdateBuilder
	Entity(entity any) UpdateBuilder
}

type updateBuilder struct {
	tableName string
	where     Where
	setPairs  []Pair
	entity    any
}

func newUpdateBuilder(tableName ...string) *updateBuilder {
	builder := &updateBuilder{
		where:    newWhere(),
		setPairs: make([]Pair, 0),
	}

	if len(tableName) > 0 {
		builder.tableName = tableName[0]
	}

	return builder
}

func (builder *updateBuilder) String() string {
	query := strings.Builder{}

	query.WriteString("UPDATE ")
	query.WriteString(builder.tableName)
	query.WriteString("\n")

	if len(builder.setPairs) > 0 {
		isParam := types.String(builder.setPairs[0].Value) == ""

		query.WriteString("SET\n")
		for index, pair := range builder.setPairs {
			query.WriteString("\t")
			query.WriteString(pair.Column)
			query.WriteString(" = ")

			if isParam {
				query.WriteString(":" + pair.Column)
			} else {
				query.WriteString(storages.ToString(pair.Value))
			}
			if index < len(builder.setPairs)-1 {
				query.WriteString(",\n")
			}
		}
		query.WriteString("\n")
	}

	whereClause := builder.where.(Query).String()
	if len(whereClause) != 0 {
		query.WriteString("WHERE \n\t")
		query.WriteString(whereClause)
		query.WriteString("\n")
	}

	return query.String()
}

func (builder *updateBuilder) GetParam() (string, bool) {
	var isParam bool
	if len(builder.setPairs) > 0 {
		isParam = types.String(builder.setPairs[0].Value) == ""
	}
	return builder.String(), isParam
}

func (builder *updateBuilder) From(tableName string) UpdateBuilder {
	builder.tableName = tableName
	return builder
}

func (builder *updateBuilder) Set(pairs ...Pair) UpdateBuilder {
	builder.setPairs = pairs
	return builder
}

func (builder *updateBuilder) Where(whereFunc func(builder Where)) UpdateBuilder {
	return builder.applyWhere(whereFunc)
}

func (builder *updateBuilder) applyWhere(whereFunc func(builder Where)) UpdateBuilder {
	whereFunc(builder.where)
	return builder
}

func (builder *updateBuilder) Entity(entity any) UpdateBuilder {
	table, _, columns := storages.Eat(entity)
	pairs := func(entity any) []Pair {
		pairs := make([]Pair, 0, len(columns))
		for _, column := range columns {
			if strings.Contains(column, ".") {
				_, after, found := strings.Cut(column, ".")
				if found {
					column = after
				}
			}
			column = strings.ReplaceAll(column, "\"", "")
			pairs = append(pairs, Pair{
				Column: column,
				Value:  ":" + column,
			})
		}
		return pairs
	}(entity)

	builder.entity = entity
	return builder.
		From(table).
		Set(pairs...)
}

func (builder *updateBuilder) Exec(ctx context.Context, args ...any) error {
	return newExecutor(ctx, builder.String(), builder.entity).Exec(args...)
}
