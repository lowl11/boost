package types

import (
	"github.com/lowl11/boost/internal/helpers/type_helper"
	"time"
)

func StringToTime(value, format string) time.Time {
	return type_helper.StringToTime(value, format)
}
