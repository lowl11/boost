package route_searcher

func (searcher *Searcher) Find() bool {
	if searcher.isVariable {
		return searcher.findVariable()
	}

	return searcher.findAny()
}

func (searcher *Searcher) Params() map[string]string {
	return searcher.params
}
