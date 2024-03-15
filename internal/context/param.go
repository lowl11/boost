package context

import (
	"github.com/google/uuid"
	"github.com/lowl11/boost/pkg/io/types"
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

func (param Param) Strings() []string {
	return strings.Split(param.value, ",")
}

func (param Param) Int() (int, error) {
	intValue, err := strconv.Atoi(param.value)
	if err != nil {
		return 0, ErrorParseIntParam(err, param.value)
	}

	return intValue, nil
}

func (param Param) MustInt() int {
	intValue, err := strconv.Atoi(param.value)
	if err != nil {
		return 0
	}

	return intValue
}

func (param Param) Bool() bool {
	return strings.ToLower(param.value) == "true"
}

func (param Param) UUID() (uuid.UUID, error) {
	uuidValue, err := uuid.Parse(param.value)
	if err != nil {
		return uuid.UUID{}, ErrorParseUUIDParam(err, param.value)
	}

	return uuidValue, nil
}

func (param Param) MustUUID() uuid.UUID {
	uuidValue, err := uuid.Parse(param.value)
	if err != nil {
		return uuid.UUID{}
	}

	return uuidValue
}

func (param Param) Bytes() []byte {
	return types.ToBytes(param.value)
}
