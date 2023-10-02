package explorer

import (
	"github.com/lowl11/boost/pkg/interfaces"
	"github.com/lowl11/boost/pkg/io/folder"
	"github.com/lowl11/boost/pkg/io/paths"
)

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

func (explorer *Explorer) getFolder(name string) (interfaces.IFolder, error) {
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
