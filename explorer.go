package boost

import (
	"github.com/lowl11/boost/pkg/io/explorer"
)

func NewExplorer(root string) Explorer {
	return explorer.New(root)
}
