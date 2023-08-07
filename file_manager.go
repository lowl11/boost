package boost

import "github.com/lowl11/lazyfile/fmanager"

func NewFileManager(root string) FileManager {
	return fmanager.New(root)
}
