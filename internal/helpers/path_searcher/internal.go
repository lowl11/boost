package path_searcher

func (searcher *Searcher) getFirst() string {
	if searcher.isEmpty() {
		return ""
	}

	return searcher.container[0]
}

func (searcher *Searcher) getLast() string {
	if searcher.isEmpty() {
		return ""
	}

	return searcher.container[len(searcher.container)-1]
}

func (searcher *Searcher) isEmpty() bool {
	return len(searcher.container) == 0
}
