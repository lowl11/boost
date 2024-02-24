package sql

import (
	"github.com/lowl11/boost/storage"
)

var (
	_repo storage.Repository
)

func getRepo() storage.Repository {
	if _repo != nil {
		return _repo
	}

	_repo = storage.NewRepo()
	return _repo
}
