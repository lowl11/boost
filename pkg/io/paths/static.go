package paths

import "strings"

func Build(args ...string) string {
	if len(args) == 0 {
		return ""
	}

	builder := strings.Builder{}
	for index, item := range args {
		builder.WriteString(item)

		if index < len(args)-1 {
			builder.WriteString(dash)
		}
	}

	return builder.String()
}

func GetFolderName(path string) (string, string) {
	pathArray := strings.Split(path, "/")
	if len(pathArray) == 1 {
		return path, path
	}

	return strings.Join(pathArray[:len(pathArray)-1], "/"), pathArray[len(pathArray)-1]
}
