package searcher

func MatchAll() map[string]any {
	return map[string]any{
		"query": map[string]any{
			"match_all": map[string]any{},
		},
	}
}

func MultiMatch(query string, fields []string) map[string]any {
	return map[string]any{
		"query": map[string]any{
			"multi_match": map[string]any{
				"query":  query,
				"fields": fields,
			},
		},
	}
}
