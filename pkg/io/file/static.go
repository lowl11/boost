package file

import (
	"bytes"
	"github.com/lowl11/lazyfile/data/errors"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// New file in given path.
// If file already exist does nothing
func New(path string, body []byte) error {
	if Exist(path) {
		return errors.FileAlreadyExist
	}

	return ioutil.WriteFile(path, body, os.ModePerm)
}

// Update updates file content
func Update(path string, body []byte) error {
	if !Exist(path) {
		return errors.FileNotFound
	}

	if err := os.Truncate(path, 0); err != nil {
		return err
	}

	file, err := os.OpenFile(path, os.O_RDWR, 0644)

	if err != nil {
		return err
	}
	defer file.Close()

	if _, err = file.Write(body); err != nil {
		return err
	}

	return nil
}

// Append add body to the already existing content & file
func Append(path string, body []byte) error {
	if !Exist(path) {
		return New(path, bytes.TrimSpace(body))
	}

	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Println("Close file error: ", path)
		}
	}()

	if _, err = file.Write(body); err != nil {
		return err
	}

	return nil
}

// Replace change content inside to another one
func Replace(path string, newContent []byte) error {
	if !Exist(path) {
		return nil
	}

	if err := Delete(path); err != nil {
		return err
	}

	return New(path, newContent)
}

// CreateFromFile create file.
// Takes content from one file and create new with given path.
// If source file does not exist returns error.
// If destination path already exist does nothing.
func CreateFromFile(source, destination string) error {
	if NotExist(source) {
		return errors.FileSourceNotFound
	}

	if Exist(destination) {
		return nil
	}

	sourceBody, err := ioutil.ReadFile(source)
	if err != nil {
		return err
	}

	return New(destination, sourceBody)
}

// Delete file by given path
func Delete(path string) error {
	return os.Remove(path)
}

// Rename file
func Rename(oldPath, newName string) error {
	newPath := strings.ReplaceAll(oldPath, filepath.Base(oldPath), newName)
	return os.Rename(oldPath, newPath)
}

// Exist check if file exist
func Exist(filePath string) bool {
	_, err := os.Stat(filePath)
	if err != nil {
		return false
	}

	return !os.IsNotExist(err)
}

// NotExist like Exist but opposite
func NotExist(filePath string) bool {
	stat, err := os.Stat(filePath)
	if err != nil {
		return false
	}

	if stat == nil {
		return os.IsNotExist(err)
	}

	return stat == nil
}

// Read get content of file
func Read(path string) ([]byte, error) {
	if !Exist(path) {
		return nil, errors.FileNotFound
	}

	stat, err := os.Stat(path)
	if err == nil && stat.IsDir() {
		return nil, errors.FileIsFolder
	}

	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return content, nil
}

// Empty is file content empty
func Empty(path string) bool {
	content, _ := Read(path)
	if content == nil {
		return true
	}

	return len(content) == 0
}
