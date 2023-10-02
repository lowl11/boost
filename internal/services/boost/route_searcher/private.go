package route_searcher

import (
	"github.com/lowl11/boost/internal/helpers/path_helper"
	"regexp"
	"strings"
)

func (searcher *Searcher) findVariable() bool {
	// if paths are equal - found
	variables, equals := path_helper.Equals(searcher.searchPath, searcher.iteratorPath)
	searcher.params = variables

	return equals
}

func (searcher *Searcher) findAny() bool {
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
