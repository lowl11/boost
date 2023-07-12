package path_searcher

import "strings"

func FirstPart(path string) string {
	if path == "" {
		return ""
	}

	container := strings.Split(path, "/")
	if len(container) < 2 {
		return ""
	}

	return container[1]
}

func (searcher *Searcher) Get() string {
	if searcher.position >= len(searcher.container) {
		return searcher.getLast()
	}

	if searcher.position < 0 {
		return searcher.getFirst()
	}

	return searcher.container[searcher.position]
}

func (searcher *Searcher) Next() *Searcher {
	searcher.position++
	return searcher
}

func (searcher *Searcher) IsVariable() bool {
	return searcher.Get()[0] == ':'
}

func (searcher *Searcher) End() bool {
	return searcher.position >= len(searcher.container)
}
