package swagger

import (
	"github.com/lowl11/boost/pkg/io/file"
	"github.com/lowl11/boost/pkg/system/types"
	"sync"
)

var (
	_cache = sync.Map{}
)

func ReadFile(name string) string {
	cacheValue, ok := _cache.Load(name)
	if ok {
		return cacheValue.(string)
	}

	content, err := file.Read("docs/" + name)
	if err != nil {
		return ""
	}

	contentStr := types.ToString(content)
	_cache.Store(name, contentStr)
	return contentStr
}
