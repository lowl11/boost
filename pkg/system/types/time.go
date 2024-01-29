package types

import (
	"github.com/lowl11/boost/pkg/system/types/timec"
	"time"
)

func StringToTime(value, format string) time.Time {
	parsedTime, err := time.Parse(format, value)
	if err != nil {
		return zeroTime()
	}

	return parsedTime
}

func IsZeroTime(timeValue timec.Time) bool {
	return timeValue.IsZero()
}

func zeroTime() time.Time {
	return time.Time{}
}
