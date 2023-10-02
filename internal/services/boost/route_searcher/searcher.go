package route_searcher

import "strings"

type Searcher struct {
	searchPath   string
	iteratorPath string

	isVariable bool
	isAny      bool

	params map[string]string
}

func New(searchPath, iteratorPath string) *Searcher {
	return &Searcher{
		searchPath:   searchPath,
		iteratorPath: iteratorPath,

		isVariable: strings.Contains(iteratorPath, ":"),
		isAny:      strings.Contains(iteratorPath, "*"),
	}
}
