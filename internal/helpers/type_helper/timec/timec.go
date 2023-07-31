package timec

import originalTime "time"

type Time originalTime.Time

func (time Time) IsZero() bool {
	return time.isZeroYear() && time.isZeroMonth() && time.isZeroDay() &&
		time.isZeroHour() && time.isZeroMinute() && time.izZeroSecond() &&
		time.isZeroNanosecond()
}

func (time Time) isZeroYear() bool {
	return time.original().Year() == 0
}

func (time Time) isZeroMonth() bool {
	return time.original().Month() == 0
}

func (time Time) isZeroDay() bool {
	return time.original().Day() == 0
}

func (time Time) isZeroHour() bool {
	return time.original().Hour() == 0
}

func (time Time) isZeroMinute() bool {
	return time.original().Minute() == 0
}

func (time Time) izZeroSecond() bool {
	return time.original().Second() == 0
}

func (time Time) isZeroNanosecond() bool {
	return time.original().Nanosecond() == 0
}

func (time Time) original() originalTime.Time {
	return originalTime.Time(time)
}
