package errors

import (
	"github.com/lowl11/boost/data/interfaces"
)

func equals(left, right interfaces.Error) bool {
	if left.HttpCode() == right.HttpCode() && left.Type() == right.Type() &&
		left.Error() == right.Error() {
		return true
	}

	return false
}
