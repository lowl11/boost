package explorer

import (
	"github.com/lowl11/boost/data/interfaces"
	"github.com/lowl11/boost/pkg/io/file"
	"github.com/lowl11/boost/pkg/io/folder"
	"github.com/lowl11/boost/pkg/io/paths"
	"github.com/lowl11/boost/pkg/system/object"
	"strings"
	"sync"
)

type Explorer struct {
	path string

	threadSafe bool
	mutex      sync.Mutex
}

func New(path string) interfaces.Explorer {
	return &Explorer{
		path: path,
	}
}

func (explorer *Explorer) ThreadSafe() interfaces.Explorer {
	explorer.threadSafe = true
	return explorer
}

func (explorer *Explorer) Empty() bool {
	if !folder.Exist(explorer.path) {
		return true
	}

	return folder.Empty(explorer.path)
}

func (explorer *Explorer) Sync() error {
	if !folder.Exist(explorer.path) {
		return folder.ErrorNotFound()
	}

	return nil
}

func (explorer *Explorer) Restore() error {
	if !folder.Exist(explorer.path) {
		return folder.Create(paths.GetFolderName(explorer.path))
	}

	return nil
}

func (explorer *Explorer) FileByPath(path string) (interfaces.File, error) {
	explorer.lock()
	defer explorer.unlock()

	filePath := paths.Build(explorer.path, path)

	content, err := file.Read(filePath)
	if err != nil {
		return nil, err
	}

	pathArray := strings.Split(path, "/")
	var name string
	if len(pathArray) > 1 {
		name = pathArray[len(pathArray)-1]
	} else {
		name = path
	}

	return file.NewFile(name, content, filePath), nil
}

func (explorer *Explorer) AddFileByPath(path string, content []byte) error {
	explorer.lock()
	defer explorer.unlock()

	filePath := paths.Build(explorer.path, path)

	if file.Exist(filePath) {
		return file.ErrorNotFound()
	}

	return file.New(filePath, content)
}

func (explorer *Explorer) UpdateFileByPath(path string, content []byte) error {
	explorer.lock()
	defer explorer.unlock()

	return file.Update(path, content)
}

func (explorer *Explorer) DeleteFileByPath(path string) error {
	explorer.lock()
	defer explorer.unlock()
	return file.Delete(paths.Build(explorer.path, path))
}

func (explorer *Explorer) FolderByPath(path string) (interfaces.Folder, error) {
	explorer.lock()
	defer explorer.unlock()

	folderPath := paths.Build(explorer.path, path)

	if !folder.Exist(folderPath) {
		return nil, folder.ErrorNotFound()
	}

	newFolder := New(folderPath)
	if explorer.threadSafe {
		newFolder.ThreadSafe()
	}

	return newFolder, nil
}

func (explorer *Explorer) AddFolderByPath(path, name string) error {
	explorer.lock()
	defer explorer.unlock()

	folderPath := paths.Build(explorer.path, path)

	if folder.Exist(folderPath) {
		return folder.ErrorNotFound()
	}

	return folder.Create(folderPath, name)
}

func (explorer *Explorer) DeleteFolderByPath(path string, force bool) error {
	explorer.lock()
	defer explorer.unlock()

	return folder.Delete(path, force)
}

func (explorer *Explorer) Path() string {
	return explorer.path
}

func (explorer *Explorer) Name() string {
	pathArray := strings.Split(explorer.path, "/")
	if len(pathArray) == 1 {
		return explorer.path
	}

	return pathArray[len(pathArray)-1]
}

func (explorer *Explorer) List() ([]object.Object, error) {
	explorer.lock()
	defer explorer.unlock()

	return folder.Objects(explorer.path)
}

func (explorer *Explorer) FileList() ([]interfaces.File, error) {
	explorer.lock()
	defer explorer.unlock()

	objects, err := folder.Objects(explorer.path)
	if err != nil {
		return nil, err
	}

	list := make([]interfaces.File, 0)
	for _, obj := range objects {
		if obj.IsFolder {
			continue
		}

		fileContent, err := file.Read(paths.Build(explorer.path, obj.Name))
		if err != nil {
			return nil, err
		}

		list = append(list, file.NewFile(obj.Name, fileContent, explorer.path))
	}

	return list, nil
}

func (explorer *Explorer) FolderList() ([]interfaces.Folder, error) {
	explorer.lock()
	defer explorer.unlock()

	objects, err := folder.Objects(explorer.path)
	if err != nil {
		return nil, err
	}

	list := make([]interfaces.Folder, 0)
	for _, obj := range objects {
		if !obj.IsFolder {
			continue
		}

		directory, err := explorer.getFolder(obj.Name)
		if err != nil {
			return nil, err
		}

		list = append(list, directory)
	}

	return list, nil
}

func (explorer *Explorer) File(name string) (interfaces.File, error) {
	explorer.lock()
	defer explorer.unlock()

	content, err := file.Read(paths.Build(explorer.path, name))
	if err != nil {
		return nil, err
	}

	return file.NewFile(name, content, explorer.path), nil
}

func (explorer *Explorer) AddFile(name string, content []byte) error {
	explorer.lock()
	defer explorer.unlock()

	return file.New(paths.Build(explorer.path, name), content)
}

func (explorer *Explorer) UpdateFile(name string, content []byte) error {
	explorer.lock()
	defer explorer.unlock()

	return file.Update(paths.Build(explorer.path, name), content)
}

func (explorer *Explorer) DeleteFile(name string) error {
	explorer.lock()
	defer explorer.unlock()

	return file.Delete(paths.Build(explorer.path, name))
}

func (explorer *Explorer) Folder(name string) (interfaces.Folder, error) {
	explorer.lock()
	defer explorer.unlock()

	return explorer.getFolder(name)
}

func (explorer *Explorer) AddFolder(name string) (interfaces.Folder, error) {
	explorer.lock()
	defer explorer.unlock()

	if err := folder.Create(explorer.path, name); err != nil {
		return nil, err
	}

	newFolder := New(paths.Build(explorer.path, name))
	if explorer.threadSafe {
		newFolder.ThreadSafe()
	}

	return newFolder, nil
}

func (explorer *Explorer) DeleteFolder(name string, force bool) error {
	explorer.lock()
	defer explorer.unlock()

	return folder.Delete(paths.Build(explorer.path, name), force)
}

func (explorer *Explorer) lock() {
	if !explorer.threadSafe {
		return
	}

	explorer.mutex.Lock()
}

func (explorer *Explorer) unlock() {
	if !explorer.threadSafe {
		return
	}

	explorer.mutex.Unlock()
}

func (explorer *Explorer) getFolder(name string) (interfaces.Folder, error) {
	folderName := paths.Build(explorer.path, name)

	if !folder.Exist(folderName) {
		return nil, folder.ErrorNotFound()
	}

	newFolder := New(folderName)
	if explorer.threadSafe {
		newFolder.ThreadSafe()
	}

	return newFolder, nil
}
