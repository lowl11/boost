package sql

import (
	"context"
	"github.com/lowl11/boost/internal/storages"
	"github.com/lowl11/boost/pkg/io/flex"
	"strings"
)

type InsertBuilder interface {
	Query

	Exec(ctx context.Context) error

	GetParamStatus() (string, bool)
	To(tableName string) InsertBuilder
	OnConflict(query string) InsertBuilder
	Values(pairs ...Pair) InsertBuilder
	Entity(entity any) InsertBuilder
	EntityList(list []any) InsertBuilder
}

type insertBuilder struct {
	tableName     string
	pairs         []Pair
	conflict      string
	multiplePairs [][]Pair
	entity        any
}

func newInsertBuilder(pairs ...Pair) *insertBuilder {
	return &insertBuilder{
		pairs:         pairs,
		multiplePairs: make([][]Pair, 0),
	}
}

func (builder *insertBuilder) String() string {
	query := strings.Builder{}

	query.WriteString("INSERT INTO ")
	query.WriteString(builder.tableName)
	query.WriteString(" (")
	for index, pair := range builder.pairs {
		query.WriteString(pair.Column)

		if index < len(builder.pairs)-1 {
			query.WriteString(", ")
		}
	}
	query.WriteString(")\n")

	isMultiple := len(builder.multiplePairs) > 0

	var isNamedValues bool
	if len(builder.pairs) > 0 {
		if builder.pairs[0].Value == nil {
			isNamedValues = true
		}
	}

	if !isMultiple {
		query.WriteString("VALUES (")
	} else {
		query.WriteString("VALUES\n")
	}

	if !isMultiple {
		for index, pair := range builder.pairs {
			if isNamedValues {
				query.WriteString(":")
				query.WriteString(pair.Column)
			} else {
				query.WriteString(storages.ToString(pair.Value))
			}

			if index < len(builder.pairs)-1 {
				query.WriteString(", ")
			}
		}
	} else {
		for mIndex, mPairs := range builder.multiplePairs {
			query.WriteString("\t(")
			for index, pair := range mPairs {
				query.WriteString(storages.ToString(pair.Value))

				if index < len(mPairs)-1 {
					query.WriteString(", ")
				}
			}

			query.WriteString(")")
			if mIndex < len(builder.multiplePairs)-1 {
				query.WriteString(",")
			}
			query.WriteString("\n")
		}
	}

	if !isMultiple {
		query.WriteString(")\n")
	}

	if len(builder.conflict) > 0 {
		query.WriteString("ON CONFLICT ")
		query.WriteString(builder.conflict)
		query.WriteString("\n")
	}

	return query.String()
}

func (builder *insertBuilder) GetParamStatus() (string, bool) {
	var isParam bool
	if len(builder.pairs) > 0 {
		isParam = builder.pairs[0].Value == nil
	}
	return builder.String(), isParam
}

func (builder *insertBuilder) Pairs(pairs ...Pair) InsertBuilder {
	if len(pairs) == 0 {
		return builder
	}

	builder.pairs = pairs
	return builder
}

func (builder *insertBuilder) To(tableName string) InsertBuilder {
	builder.tableName = tableName
	return builder
}

func (builder *insertBuilder) OnConflict(query string) InsertBuilder {
	builder.conflict = query
	return builder
}

func (builder *insertBuilder) Values(pairs ...Pair) InsertBuilder {
	builder.multiplePairs = append(builder.multiplePairs, pairs)
	return builder
}

func (builder *insertBuilder) Entity(entity any) InsertBuilder {
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
		Pairs(pairs...).
		To(table)
}

func (builder *insertBuilder) EntityList(list []any) InsertBuilder {
	if len(list) == 0 {
		return builder
	}

	table, _, columns := storages.Eat(list[0])

	getPairs := func(columns []string, entity any) []Pair {
		pairs := make([]Pair, 0, len(columns))
		fStr, _ := flex.Struct(entity)
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
				Value:  fStr.FieldValueByTag("db", column),
			})
		}
		return pairs
	}

	for _, entity := range list {
		builder.Values(getPairs(columns, entity)...)
	}

	if builder.pairs == nil {
		builder.pairs = make([]Pair, 0, len(columns))
	}
	for _, col := range columns {
		builder.pairs = append(builder.pairs, Pair{
			Column: col,
		})
	}
	builder.tableName = table
	return builder
}

func (builder *insertBuilder) Exec(ctx context.Context) error {
	return newExecutor(ctx, builder.String(), builder.entity).Exec()
}
