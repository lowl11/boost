package interfaces

import (
	"github.com/lowl11/boost/pkg/io/object"
)

type IExplorer interface {
	IFolder

	/*
		ThreadSafe turns on thread safe mode.
		Created folder objects by IManager inherit thread safe mode
	*/
	ThreadSafe() IExplorer

	/*
		FileByPath get IFile object by given path.
		Path is path inside given root path
	*/
	FileByPath(path string) (IFile, error)

	/*
		AddFileByPath creates new file by given path.
		Path is path inside given root path
	*/
	AddFileByPath(path string, content []byte) error

	/*
		UpdateFileByPath update file content by given path.
		Path is path inside given root path
	*/
	UpdateFileByPath(path string, content []byte) error

	/*
		DeleteFileByPath removes file by given path.
		Path is path inside given root path
	*/
	DeleteFileByPath(path string) error

	/*
		FolderByPath get IFolder object by give path.
		Path is path inside given root path
	*/
	FolderByPath(path string) (IFolder, error)

	/*
		AddFolderByPath	creates path by given path.
		Path is path inside given root path
	*/
	AddFolderByPath(path, name string) error

	/*
		DeleteFolderByPath removes folder by given path.
		Path is path inside given root path
	*/
	DeleteFolderByPath(path string, force bool) error
}

type IFolder interface {
	// Empty is folder does not contain objects
	Empty() bool

	// Name returns name of folder
	Name() string

	// Path returns path of folder located (without folder name)
	Path() string

	/*
		Sync synchronize folder state.
		If folder does not exist, will return errors.FolderNotExist
	*/
	Sync() error

	/*
		Restore get back folder (without content)
	*/
	Restore() error

	/*
		List returns list of objects which contains current folder (objects = files + folders)
	*/
	List() ([]object.Object, error)

	/*
		FileList returns list of IFile objects which contains current folder
	*/
	FileList() ([]IFile, error)

	/*
		FolderList returns list of IFolder objects which contains current folder
	*/
	FolderList() ([]IFolder, error)

	/*
		File returns IFile object by name (with extension)
	*/
	File(name string) (IFile, error)

	/*
		AddFile creates new file
	*/
	AddFile(name string, content []byte) error

	/*
		UpdateFile updates file content by given name
	*/
	UpdateFile(name string, content []byte) error

	/*
		DeleteFile removes file by given name
	*/
	DeleteFile(name string) error

	/*
		Folder returns IFolder object by name
	*/
	Folder(name string) (IFolder, error)

	/*
		AddFolder creates new folder by name
	*/
	AddFolder(name string) (IFolder, error)

	/*
		DeleteFolder removes folder by name
	*/
	DeleteFolder(name string, force bool) error
}

type IFile interface {
	// Name returns name of file
	Name() string

	// Bytes returns content of file in bytes
	Bytes() []byte

	// String returns content of file in string
	String() string

	// Path returns path of file which is located (without file name)
	Path() string

	// Update updates file content
	Update(content []byte) error

	// Delete removes file
	Delete() error

	/*
		Sync synchronize folder state.
		If folder does not exist, will return errors.FolderNotExist
	*/
	Sync() error

	/*
		Restore get back folder (without content)
	*/
	Restore() error

	/*
		IsDestroyed returns status is file destroyed or not (true if it is removed)
	*/
	IsDestroyed() bool
}
