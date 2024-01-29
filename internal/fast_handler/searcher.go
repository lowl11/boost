package fast_handler

import (
	"regexp"
	"strings"
	"unicode/utf8"
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
	variables, equals := pathEquals(searcher.searchPath, searcher.iteratorPath)
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

func pathEquals(searchPath, routePath string) (map[string]string, bool) {
	searchArray := strings.Split(searchPath, "/")
	routeArray := strings.Split(routePath, "/")

	if len(searchArray) != len(routeArray) {
		return nil, false
	}

	variables := make(map[string]string)
	for index, item := range searchArray {
		if index == 0 && item == "" {
			continue
		}

		if index >= len(routeArray) {
			return nil, false
		}

		if isVar(routeArray[index]) {
			variables[routeArray[index][1:]] = searchArray[index]
			continue
		}

		if item != routeArray[index] {
			return nil, false
		}
	}

	return variables, true
}

func isVar(value string) bool {
	if value == "" {
		return false
	}

	return value[0] == ':'
}

func isLastSlash(path string) bool {
	if path == "" {
		return false
	}

	length := utf8.RuneCountInString(path)
	return path[length-1] == '/'
}

func removeLast(path string) string {
	if path == "" {
		return ""
	}

	length := utf8.RuneCountInString(path)
	return path[:length-1]
}
