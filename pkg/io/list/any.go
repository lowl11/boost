package list

import "github.com/lowl11/boost/pkg/io/flex"

func Any(slice any) []any {
	fx, err := flex.Slice(slice)
	if err != nil {
		return []any{}
	}
	return fx.Elements()
}
