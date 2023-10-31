package elk_service

import "github.com/lowl11/boost/internal/boosties/errors"

func ErrorGetAllIndices(err error) error {
	return errors.
		New("Get all indices error").
		SetType("ELK_GetAllIndicesError").
		SetError(err)
}

func ErrorCreateIndex(err error) error {
	return errors.
		New("Create new index error").
		SetType("ELK_CreateIndexError").
		SetError(err)
}

func ErrorDeleteIndex(err error) error {
	return errors.
		New("Delete index error").
		SetType("ELK_DeleteIndexError").
		SetError(err)
}

func ErrorBindAlias(err error) error {
	return errors.
		New("Bind alias error").
		SetType("ELK_BindAliasError").
		SetError(err)
}

func ErrorInsertData(err error) error {
	return errors.
		New("Insert data error").
		SetType("ELK_InsertDataError").
		SetError(err)
}

func ErrorGetAllDocuments(err error) error {
	return errors.
		New("Get all documents error").
		SetType("ELK_GetAllDocsError").
		SetError(err)
}
