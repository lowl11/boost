package folder

import (
	"fmt"
	"github.com/lowl11/boost/pkg/io/object"
	"github.com/lowl11/boost/pkg/io/paths"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// Create folder in given path.
// If folder already exist does nothing
func Create(path, name string) error {
	newFolderPath := paths.Build(path, name)
	if Exist(newFolderPath) {
		return nil
	}

	return os.Mkdir(newFolderPath, os.ModePerm)
}

// Delete folder.
// Given flag withContent delete all files in folder.
// If folder does not exist does nothing
func Delete(path string, withContent bool) error {
	if NotExist(path) {
		return nil
	}

	if withContent {
		if err := os.RemoveAll(path); err != nil {
			return err
		}

		return nil
	}

	return os.Remove(path)
}

// Rename folder name
func Rename(oldPath, newName string) error {
	newPath := strings.ReplaceAll(oldPath, filepath.Base(oldPath), newName)
	return os.Rename(oldPath, newPath)
}

// Exist folder
func Exist(folderPath string) bool {
	_, err := os.Stat(folderPath)
	return !os.IsNotExist(err)
}

// NotExist folder
func NotExist(folderPath string) bool {
	stat, err := os.Stat(folderPath)
	if err != nil {
		return false
	}

	if stat == nil {
		return os.IsNotExist(err)
	}

	return stat == nil
}

// Objects return list of files & folders in custom model.
// Also returned list of objects sorted by alphabet and "isDirectory" flag
func Objects(path string) ([]object.Object, error) {
	objectList := make([]object.Object, 0)
	folderObjects, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, objectItem := range folderObjects {
		objectName := objectItem.Name()
		isFolder := objectItem.IsDir()
		objectPath := buildObjectPath(path, objectName)

		objectList = append(objectList, object.Object{
			Name:        objectName,
			Path:        objectPath,
			IsFolder:    isFolder,
			ObjectCount: Count(objectPath),
		})
	}

	// sort by folders & files
	sort.Slice(objectList, func(i, j int) bool {
		return objectList[i].IsFolder
	})

	// sort folder by alphabet
	sort.Slice(objectList, func(i, j int) bool {
		return (objectList[i].Name < objectList[j].Name) && (objectList[i].IsFolder && objectList[j].IsFolder)
	})

	// sort files by alphabet
	sort.Slice(objectList, func(i, j int) bool {
		return (objectList[i].Name < objectList[j].Name) && (!objectList[i].IsFolder && !objectList[j].IsFolder)
	})

	return objectList, nil
}

// ObjectsWithDepth return list of files & folders in custom model with all children.
// Also returned list of objects sorted by alphabet and "isDirectory" flag
func ObjectsWithDepth(path, memoryPath string) ([]object.Object, error) {
	objectList := make([]object.Object, 0, 100)
	folderObjects, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	relativeRootPath := filepath.Dir(path)
	if memoryPath == "" {
		memoryPath = relativeRootPath
	}

	for _, objectItem := range folderObjects {
		// main meta info
		objectName := objectItem.Name()
		isFolder := objectItem.IsDir()
		objectPath := buildObjectPath(path, objectName)
		objectCount := Count(objectPath)

		// getting children
		children := make([]object.Object, 0, objectCount)
		children, err = ObjectsWithDepth(objectPath, memoryPath)
		if err != nil {
			children = make([]object.Object, 0, objectCount)
		}

		// memory path
		//var nextPath string
		//if isFolder {
		//	nextPath = objectName
		//}
		//memoryPath = fmt.Sprintf("%s/%s", memoryPath, nextPath)
		//objectMemoryPath := buildMemoryObjectPath(memoryPath, objectName)
		objectMemoryPath := memoryPath

		objectList = append(objectList, object.Object{
			Name:         objectName,
			Path:         objectPath,
			RelativePath: objectMemoryPath,
			IsFolder:     isFolder,
			ObjectCount:  objectCount,
			Children:     children,
		})
	}

	// sort by folders & files
	sort.Slice(objectList, func(i, j int) bool {
		return objectList[i].IsFolder
	})

	// sort folder by alphabet
	sort.Slice(objectList, func(i, j int) bool {
		return (objectList[i].Name < objectList[j].Name) && (objectList[i].IsFolder && objectList[j].IsFolder)
	})

	// sort files by alphabet
	sort.Slice(objectList, func(i, j int) bool {
		return (objectList[i].Name < objectList[j].Name) && (!objectList[i].IsFolder && !objectList[j].IsFolder)
	})

	return objectList, nil
}

// Count return count of folder objects
func Count(path string) int {
	files, err := os.ReadDir(path)
	if err != nil {
		return 0
	}

	return len(files)
}

// Empty is folder empty
func Empty(path string) bool {
	return Count(path) == 0
}

func buildObjectPath(path, name string) string {
	builder := strings.Builder{}
	builder.Grow(len(path) + len(name) + 1)

	_, _ = fmt.Fprintf(&builder, "%s/%s", path, name)
	return builder.String()
}
