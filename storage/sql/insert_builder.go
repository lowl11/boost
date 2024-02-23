package sql

import (
	"github.com/lowl11/boost/internal/storages"
	"strings"
)

type insertBuilder struct {
	tableName     string
	pairs         []Pair
	conflict      string
	multiplePairs [][]Pair
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
