package paths

import (
	"github.com/lowl11/boost/pkg/io/list"
	"strings"
)

func Build(args ...string) string {
	builder := strings.Builder{}
	list.Of(args).Each(func(index int, item string) {
		builder.WriteString(item)

		if index < len(args)-1 {
			builder.WriteString(dash)
		}
	})
	return builder.String()
}

func GetFolderName(path string) (string, string) {
	pathArray := strings.Split(path, "/")
	if len(pathArray) == 1 {
		return path, path
	}

	return strings.Join(pathArray[:len(pathArray)-1], "/"), pathArray[len(pathArray)-1]
}
