package path_searcher

import "strings"

type Searcher struct {
	container []string
	position  int
}

func New(path string) *Searcher {
	return &Searcher{
		container: strings.Split(path, "/")[1:],
	}
}
