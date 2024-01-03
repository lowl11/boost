package boost

import (
	"github.com/lowl11/boost/pkg/io/explorer"
)

// NewExplorer returns object of "Explorer" - File System Management service
func NewExplorer(root string) Explorer {
	return explorer.New(root)
}
