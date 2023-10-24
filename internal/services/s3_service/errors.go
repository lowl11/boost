package s3_service

import (
	"github.com/lowl11/boost/data/errors"
)

func ErrorConnect(err error) error {
	return errors.
		New("Connect to S3 error").
		SetType("ConnectS3Error").
		SetError(err)
}

func ErrorCreateFolder(err error, path string) error {
	return errors.
		New("Create folder error").
		SetType("S3_CreateFolderError").
		SetError(err).
		AddContext("path", path)
}

func ErrorCreateFile(err error, path string) error {
	return errors.
		New("Create file error").
		SetType("S3_CreateFileError").
		SetError(err).
		AddContext("path", path)
}

func ErrorDeleteObject(err error, path string) error {
	return errors.
		New("Delete object error").
		SetType("S3_DeleteObjectError").
		SetError(err).
		AddContext("path", path)
}

func ErrorCopyObject(err error, old, new string) error {
	return errors.
		New("Copy object error").
		SetType("S3_CopyObjectError").
		SetError(err).
		AddContext("old", old).
		AddContext("new", new)
}

func ErrorGetObjects(err error, path *string) error {
	return errors.
		New("Get objects error").
		SetType("S3_GetObjectsError").
		SetError(err).
		AddContext("path", path)
}
