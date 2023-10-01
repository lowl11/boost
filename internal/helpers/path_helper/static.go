package path_helper

import (
	"strings"
	"unicode/utf8"
)

func Equals(searchPath, routePath string) (map[string]string, bool) {
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

		if IsVariable(routeArray[index]) {
			variables[routeArray[index][1:]] = searchArray[index]
			continue
		}

		if item != routeArray[index] {
			return nil, false
		}
	}

	return variables, true
}

func IsVariable(value string) bool {
	if value == "" {
		return false
	}

	return value[0] == ':'
}

func IsLastSlash(path string) bool {
	if path == "" {
		return false
	}

	length := utf8.RuneCountInString(path)
	return path[length-1] == '/'
}

func RemoveLast(path string) string {
	if path == "" {
		return ""
	}

	length := utf8.RuneCountInString(path)
	return path[:length-1]
}
