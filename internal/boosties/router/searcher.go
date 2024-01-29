package router

import (
	"github.com/lowl11/boost/internal/helpers/path_helper"
	"regexp"
	"strings"
)

type searcherService struct {
	searchPath   string
	iteratorPath string

	isVariable bool
	isAny      bool

	params map[string]string
}

func newSearcher(searchPath, iteratorPath string) *searcherService {
	return &searcherService{
		searchPath:   searchPath,
		iteratorPath: iteratorPath,

		isVariable: strings.Contains(iteratorPath, ":"),
		isAny:      strings.Contains(iteratorPath, "*"),
	}
}

func (searcher *searcherService) Find() bool {
	if searcher.isVariable {
		return searcher.findVariable()
	}

	return searcher.findAny()
}

func (searcher *searcherService) Params() map[string]string {
	return searcher.params
}

func (searcher *searcherService) findVariable() bool {
	// if paths are equal - found
	variables, equals := path_helper.Equals(searcher.searchPath, searcher.iteratorPath)
	searcher.params = variables

	return equals
}

func (searcher *searcherService) findAny() bool {
	if !searcher.isAny {
		return false
	}

	iteratorPath := strings.ReplaceAll(searcher.iteratorPath, "*", "")

	reg := regexp.MustCompile("(" + iteratorPath + ").*?")
	match := reg.FindAllString(searcher.searchPath, -1)
	if len(match) == 0 {
		return false
	}

	return match[0] == iteratorPath
}
