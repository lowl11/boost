package sql

import (
	"github.com/lowl11/boost/internal/storages"
	"strings"
)

type whereBuilder struct {
	alias      string
	conditions []string
	or         bool
}

func newWhere(alias ...string) *whereBuilder {
	where := &whereBuilder{
		conditions: make([]string, 0, 5),
	}

	if len(alias) > 0 && len(alias[0]) > 0 {
		where.alias = alias[0]
	}

	return where
}

func newWhereOr(alias ...string) *whereBuilder {
	where := &whereBuilder{
		conditions: make([]string, 0, 5),
		or:         true,
	}

	if len(alias) > 0 && len(alias[0]) > 0 {
		where.alias = alias[0]
	}

	return where
}

func (where *whereBuilder) String() string {
	separator := " AND " + "\n\t"
	if where.or {
		separator = " OR "
	}

	result := strings.Join(where.conditions, separator)
	if where.or {
		return storages.WhereBrackets(result)
	}

	return result
}

func (where *whereBuilder) SetAlias(alias string) Where {
	where.alias = alias
	return where
}

func (where *whereBuilder) Not(condition func(Where) Where) Where {
	where.add(storages.BuildWhereNot(condition(newWhere(where.alias)).(Query).String()))
	return where
}

func (where *whereBuilder) Or(condition func(Where) Where) Where {
	where.add(condition(newWhereOr(where.alias)).(Query).String())
	return where
}

func (where *whereBuilder) Bool(field string, result bool) Where {
	if !result {
		return where.add(storages.BuildWhereNot(field))
	}

	return where.add(field)
}

func (where *whereBuilder) Equal(field string, value any) Where {
	return where.add(storages.BuildWhere(where.alias, field, "=", value))
}

func (where *whereBuilder) NotEqual(field string, value any) Where {
	return where.add(storages.BuildWhere(where.alias, field, "!=", value))
}

func (where *whereBuilder) In(field string, values []any) Where {
	queryArray := strings.Builder{}

	queryArray.WriteString("(")
	for index, value := range values {
		queryArray.WriteString(storages.ToString(value))

		if index < len(values)-1 {
			queryArray.WriteString(", ")
		}
	}
	queryArray.WriteString(")")

	return where.add(storages.BuildWhereArray(field, "IN", queryArray.String()))
}

func (where *whereBuilder) Is(field string, value any) Where {
	return where.add(storages.BuildWhere(where.alias, field, "IS", value))
}

func (where *whereBuilder) IsNull(field string) Where {
	return where.add(storages.BuildWhere(where.alias, field, "IS", "$NULL"))
}

func (where *whereBuilder) IsNotNull(field string) Where {
	return where.add(storages.BuildWhere(where.alias, field, "IS NOT", "$NULL"))
}

func (where *whereBuilder) Like(field, value string) Where {
	return where.add(storages.BuildWhere(where.alias, field, "LIKE", value))
}

func (where *whereBuilder) ILike(field, value string) Where {
	return where.add(storages.BuildWhere(where.alias, field, "ILIKE", value))
}

func (where *whereBuilder) Between(field string, left, right any) Where {
	return where.add(storages.BuildWhereBetween(field, left, right))
}

func (where *whereBuilder) Gte(field string, value any) Where {
	return where.add(storages.BuildWhere(where.alias, field, ">=", value))
}

func (where *whereBuilder) Gt(field string, value any) Where {
	return where.add(storages.BuildWhere(where.alias, field, ">", value))
}

func (where *whereBuilder) Lte(field string, value any) Where {
	return where.add(storages.BuildWhere(where.alias, field, "<=", value))
}

func (where *whereBuilder) Lt(field string, value any) Where {
	return where.add(storages.BuildWhere(where.alias, field, "<", value))
}

func (where *whereBuilder) add(condition string) *whereBuilder {
	where.conditions = append(where.conditions, condition)
	return where
}
