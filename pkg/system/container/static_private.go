package container

import "sync"

var (
	_container sync.Map
)

func init() {
	_container = sync.Map{}
}
