package context

import (
	"strconv"
	"strings"
)

type Param struct {
	value string
}

func NewParam(value string) *Param {
	return &Param{
		value: value,
	}
}

func (param Param) String() string {
	return param.value
}

func (param Param) Int() int {
	intValue, err := strconv.Atoi(param.value)
	if err != nil {
		return 0
	}

	return intValue
}
func (param Param) Bool() bool {
	return strings.ToLower(param.value) == "true"
}
