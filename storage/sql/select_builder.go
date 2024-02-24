package sql

import (
	"context"
	"github.com/lowl11/boost/internal/storages"
	"strings"
)

type selectBuilder struct {
	columns         []string
	tableName       string
	aliasName       string
	joins           []Join
	where           Where
	orderByColumns  []string
	isDescending    bool
	havingAggregate Aggregate
	groupAggregate  Aggregate
	groupByColumns  []string
	offset          int
	limit           int
}

func newSelectBuilder(columns ...string) *selectBuilder {
	builder := &selectBuilder{
		columns:        columns,
		joins:          make([]Join, 0, 2),
		where:          newWhere(),
		groupByColumns: make([]string, 0),

		offset: -1,
		limit:  -1,
	}
	builder.refreshColumns()
	return builder
}

func (builder *selectBuilder) Single(ctx context.Context, args ...any) Scanner {
	return newScanner(builder.String(), ctx, args...).Single()
}

func (builder *selectBuilder) List(ctx context.Context, args ...any) Scanner {
	return newScanner(builder.String(), ctx, args...)
}

func (builder *selectBuilder) Scan(ctx context.Context, result any, args ...any) error {
	return newScanner(builder.String(), ctx, args...).Scan(result)
}

func (builder *selectBuilder) ScanSingle(ctx context.Context, result any, args ...any) error {
	return newScanner(builder.String(), ctx, args...).Single().Scan(result)
}

func (builder *selectBuilder) String() string {
	// builder
	query := strings.Builder{}
	query.Grow(300)

	// select
	query.WriteString("SELECT\n\t")

	// append table
	if len(builder.columns) == 0 {
		query.WriteString("*")
	} else {
		query.WriteString(strings.Join(builder.columns, ", \n\t"))
	}

	query.WriteString("\nFROM ")
	query.WriteString(builder.tableName)
	if len(builder.aliasName) > 0 {
		query.WriteString(" AS ")
		query.WriteString(builder.aliasName)
	}

	// append join
	for _, join := range builder.joins {
		query.WriteString("\n\t")
		query.WriteString(join.String())
	}

	// append where
	whereClause := builder.where.(Query).String()
	if len(whereClause) != 0 {
		query.WriteString("\nWHERE \n\t")
		query.WriteString(whereClause)
		query.WriteString("\n")
	}

	// append order by
	if len(builder.orderByColumns) > 0 {
		query.WriteString("ORDER BY ")
		query.WriteString(strings.Join(builder.orderByColumns, ", "))
		if !builder.isDescending {
			query.WriteString(" ASC")
		} else {
			query.WriteString(" DESC")
		}
		query.WriteString("\n")
	}

	// append having
	if builder.havingAggregate != nil {
		query.WriteString("HAVING ")
		query.WriteString(builder.havingAggregate.(Query).String())
		query.WriteString("\n")
	}

	// append group by
	if builder.groupAggregate != nil {
		query.WriteString("GROUP BY ")
		if len(builder.groupByColumns) > 0 {
			for index, column := range builder.groupByColumns {
				query.WriteString(column)

				if index < len(builder.groupByColumns)-1 {
					query.WriteString(", ")
				}
			}
		} else {
			query.WriteString(builder.groupAggregate.(Query).String())
		}
		query.WriteString("\n")
	}

	// append offset
	if builder.offset > -1 {
		query.WriteString("\nOFFSET " + storages.ToString(builder.offset) + "\n")
	}

	// append limit
	if builder.limit > -1 {
		query.WriteString("LIMIT " + storages.ToString(builder.limit) + "\n")
	}

	return query.String()
}

func (builder *selectBuilder) Select(columns ...string) SelectBuilder {
	return builder.setColumns(columns...)
}

func (builder *selectBuilder) From(tableName string) SelectBuilder {
	return builder.setTable(tableName)
}

func (builder *selectBuilder) As(aliasName string) SelectBuilder {
	return builder.setAlias(aliasName)
}

func (builder *selectBuilder) Join(tableName, aliasName, joinColumn, mainColumn string) SelectBuilder {
	return builder.addJoin("INNER JOIN ", tableName, aliasName, joinColumn, mainColumn)
}

func (builder *selectBuilder) LeftJoin(tableName, aliasName, joinColumn, mainColumn string) SelectBuilder {
	return builder.addJoin("LEFT JOIN ", tableName, aliasName, joinColumn, mainColumn)
}

func (builder *selectBuilder) RightJoin(tableName, aliasName, joinColumn, mainColumn string) SelectBuilder {
	return builder.addJoin("RIGHT JOIN ", tableName, aliasName, joinColumn, mainColumn)
}

func (builder *selectBuilder) Where(whereFunc func(builder Where)) SelectBuilder {
	return builder.applyWhere(whereFunc)
}

func (builder *selectBuilder) OrderBy(columns ...string) SelectBuilder {
	return builder.setOrderByColumns(columns...)
}

func (builder *selectBuilder) OrderByDescending(columns ...string) SelectBuilder {
	return builder.
		setOrderByColumns(columns...).
		orderByDescending()
}

func (builder *selectBuilder) Having(aggregateFunc func(aggregate Aggregate)) SelectBuilder {
	return builder.setHaving(aggregateFunc)
}

func (builder *selectBuilder) GroupBy(columns ...string) SelectBuilder {
	return builder.setGroupBy(columns...)
}

func (builder *selectBuilder) GroupByAggregate(aggregateFunc func(aggregate Aggregate)) SelectBuilder {
	return builder.setGroupByAggregate(aggregateFunc)
}

func (builder *selectBuilder) Offset(offset int) SelectBuilder {
	return builder.setOffset(offset)
}

func (builder *selectBuilder) Limit(limit int) SelectBuilder {
	return builder.setLimit(limit)
}

func (builder *selectBuilder) Page(pageSize, pageNumber int) SelectBuilder {
	if pageSize < 0 {
		return builder
	}

	from := pageSize * (pageNumber - 1)
	to := from + pageSize
	return builder.setOffset(from).setLimit(to)
}

func (builder *selectBuilder) orderByDescending() *selectBuilder {
	builder.isDescending = true
	return builder
}

func (builder *selectBuilder) setOrderByColumns(columns ...string) *selectBuilder {
	builder.orderByColumns = columns
	return builder
}

func (builder *selectBuilder) setColumns(columns ...string) *selectBuilder {
	if len(columns) == 0 {
		return builder
	}

	builder.columns = columns
	return builder.refreshColumns()
}

func (builder *selectBuilder) setTable(tableName string) *selectBuilder {
	if len(tableName) == 0 {
		return builder
	}

	builder.tableName = storages.MakeName(tableName)
	return builder
}

func (builder *selectBuilder) setAlias(aliasName string) *selectBuilder {
	if len(aliasName) == 0 {
		return builder
	}

	builder.aliasName = storages.MakeName(aliasName)
	builder.where.SetAlias(builder.aliasName)
	return builder.refreshColumns()
}

func (builder *selectBuilder) applyWhere(whereFunc func(builder Where)) *selectBuilder {
	whereFunc(builder.where)
	return builder
}

func (builder *selectBuilder) setHaving(aggregateFunc func(aggregate Aggregate)) *selectBuilder {
	aggregateFunc(builder.havingAggregate)
	return builder
}

func (builder *selectBuilder) setGroupBy(columns ...string) *selectBuilder {
	builder.groupByColumns = columns
	return builder
}

func (builder *selectBuilder) setGroupByAggregate(aggregateFunc func(aggregate Aggregate)) *selectBuilder {
	aggregateFunc(builder.groupAggregate)
	return builder
}

func (builder *selectBuilder) setOffset(offset int) *selectBuilder {
	if offset < 0 {
		return builder
	}

	builder.offset = offset
	return builder
}

func (builder *selectBuilder) setLimit(limit int) *selectBuilder {
	if limit < 0 {
		return builder
	}

	builder.limit = limit
	return builder
}

func (builder *selectBuilder) addJoin(joinType, tableName, aliasName, joinColumn, mainColumn string) *selectBuilder {
	builder.joins = append(builder.joins, newJoin(joinType).
		Table(tableName).
		Alias("\""+aliasName+"\"").
		JoinColumn(joinColumn).
		MainColumn(mainColumn))
	return builder
}

func (builder *selectBuilder) refreshColumns() *selectBuilder {
	if len(builder.columns) == 0 {
		return builder
	}

	isNamed := func(name string) bool {
		return strings.Contains(name, "\"")
	}

	isCustom := func(name string) bool {
		return strings.Contains(name, " ")
	}

	isPointer := func(name string) bool {
		return strings.Contains(name, ".*")
	}

	isAliased := func(name string) bool {
		return strings.Contains(name, ".")
	}

	for i := 0; i < len(builder.columns); i++ {
		// aggregate
		if strings.Contains(builder.columns[i], "COUNT(") {
			continue
		}

		if isCustom(builder.columns[i]) {
			before, after, _ := strings.Cut(builder.columns[i], " ")
			table, column, _ := strings.Cut(before, ".")
			if strings.Contains(table, "\"") {
				builder.columns[i] = table + "." + column + " " + after
			} else {
				builder.columns[i] = "\"" + table + "\"" + ".\"" + column + "\" " + after
			}

			continue
		}

		// already field name with dot. Example: product.title -> "product"."title"
		if before, after, found := strings.Cut(builder.columns[i], "."); found {
			if isNamed(builder.columns[i]) || isPointer(builder.columns[i]) {
				continue
			}

			builder.columns[i] = storages.MakeName(before) + "." + storages.MakeName(after)
			continue
		}

		if len(builder.aliasName) != 0 {
			if isNamed(builder.columns[i]) && isAliased(builder.columns[i]) {
				continue
			}

			if isNamed(builder.columns[i]) {
				builder.columns[i] = builder.aliasName + "." + builder.columns[i]
			} else {
				builder.columns[i] = builder.aliasName + "." + storages.MakeName(builder.columns[i])
			}
		} else {
			if !isNamed(builder.columns[i]) {
				builder.columns[i] = storages.MakeName(builder.columns[i])
			}
		}
	}

	return builder
}
