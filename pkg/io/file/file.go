package file

import (
	"github.com/lowl11/boost/pkg/io/paths"
)

type File struct {
	name        string
	content     []byte
	path        string
	isDestroyed bool
}

func NewFile(name string, content []byte, path string) *File {
	return &File{
		name:    name,
		content: content,
		path:    path,
	}
}

func (file *File) Name() string {
	if file.isDestroyed {
		return ""
	}

	return file.name
}

func (file *File) Bytes() []byte {
	if file.isDestroyed {
		return nil
	}

	return file.content
}

func (file *File) String() string {
	if file.isDestroyed {
		return ""
	}

	return string(file.content)
}

func (file *File) Path() string {
	if file.isDestroyed {
		return ""
	}

	return file.path
}

func (file *File) Update(content []byte) error {
	if file.isDestroyed {
		return ErrorAlreadyDestroyed()
	}

	if err := Update(paths.Build(file.path, file.name), content); err != nil {
		return err
	}

	file.content = content
	return nil
}

func (file *File) Delete() error {
	if err := Delete(paths.Build(file.path, file.name)); err != nil {
		return err
	}

	file.isDestroyed = true
	return nil
}

func (file *File) IsDestroyed() bool {
	return file.isDestroyed
}

func (file *File) Sync() error {
	if !Exist(paths.Build(file.path, file.name)) {
		return ErrorNotFound()
	}

	return nil
}

func (file *File) Restore() error {
	return New(paths.Build(file.path, file.name), file.content)
}
