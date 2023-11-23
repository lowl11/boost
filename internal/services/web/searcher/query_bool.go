package searcher

import (
	"github.com/lowl11/boost/data/enums/query_types"
)

type QueryPair struct {
	Mode  string // Match, Term, Filter
	Field string
	Value any
}

type BoolQuery struct {
	must   []QueryPair
	should []QueryPair
}

func QueryBool() *BoolQuery {
	return &BoolQuery{
		must:   make([]QueryPair, 0),
		should: make([]QueryPair, 0),
	}
}

func (q *BoolQuery) Get() map[string]any {
	query := map[string]any{
		"query": map[string]any{
			"bool": map[string]any{},
		},
	}

	if len(q.must) > 0 {
		expressions := make([]map[string]any, 0, len(q.must))
		for _, expression := range q.must {
			expressions = append(expressions, map[string]any{
				expression.Mode: map[string]any{
					expression.Field: expression.Value,
				},
			})
		}
		query["query"].(map[string]any)["bool"].(map[string]any)["must"] = expressions
	}

	if len(q.should) > 0 {
		expressions := make([]map[string]any, 0, len(q.must))
		for _, expression := range q.should {
			expressions = append(expressions, map[string]any{
				expression.Mode: map[string]any{
					expression.Field: expression.Value,
				},
			})
		}
		query["query"].(map[string]any)["bool"].(map[string]any)["should"] = expressions
	}

	return query
}

func (q *BoolQuery) Must(mode, field string, value any) *BoolQuery {
	q.must = append(q.must, QueryPair{
		Mode:  mode,
		Field: field,
		Value: value,
	})
	return q
}

func (q *BoolQuery) MustMatch(field string, value any) *BoolQuery {
	return q.Must(query_types.Match, field, value)
}

func (q *BoolQuery) MustTerm(field string, value any) *BoolQuery {
	return q.Must(query_types.Term, field, value)
}

func (q *BoolQuery) MustFilter(field string, value any) *BoolQuery {
	return q.Must(query_types.Filter, field, value)
}

func (q *BoolQuery) Should(mode, field string, value any) *BoolQuery {
	q.should = append(q.should, QueryPair{
		Mode:  mode,
		Field: field,
		Value: value,
	})
	return q
}

func (q *BoolQuery) ShouldMatch(field string, value any) *BoolQuery {
	return q.Should(query_types.Match, field, value)
}

func (q *BoolQuery) ShouldTerm(field string, value any) *BoolQuery {
	return q.Should(query_types.Term, field, value)
}

func (q *BoolQuery) ShouldFilter(field string, value any) *BoolQuery {
	return q.Should(query_types.Filter, field, value)
}
