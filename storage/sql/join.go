package sql

import "strings"

type joinBuilder struct {
	joinType   string
	table      string
	alias      string
	joinColumn string
	mainColumn string
}

func newJoin(joinType string) *joinBuilder {
	return &joinBuilder{
		joinType: joinType,
	}
}

func (join *joinBuilder) String() string {
	query := strings.Builder{}

	query.WriteString(join.joinType)
	query.WriteString(join.table)
	query.WriteString(" AS ")
	query.WriteString(join.alias)
	query.WriteString(" ON ")
	query.WriteString("(")
	query.WriteString(join.joinColumn)
	query.WriteString(" = ")
	query.WriteString(join.mainColumn)
	query.WriteString(")")

	return query.String()
}

func (join *joinBuilder) Table(tableName string) Join {
	join.table = tableName
	return join
}

func (join *joinBuilder) Alias(aliasName string) Join {
	join.alias = aliasName
	return join
}

func (join *joinBuilder) JoinColumn(column string) Join {
	before, after, found := strings.Cut(column, ".")
	if found {
		join.joinColumn = "\"" + before + "\"." + after
	} else {
		join.joinColumn = column
	}
	return join
}

func (join *joinBuilder) MainColumn(column string) Join {
	before, after, found := strings.Cut(column, ".")
	if found {
		join.mainColumn = "\"" + before + "\"." + after
	} else {
		join.mainColumn = column
	}
	return join
}
