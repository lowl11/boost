package domain

import (
	"github.com/lowl11/boost/pkg/system/types"
)

type createdWithIDNumber struct {
	ID int `json:"id"`
}

type createdWithIDString struct {
	ID string `json:"id"`
}

func NewCreatedWithID(id any) any {
	numberID, ok := id.(int)
	if ok {
		return createdWithIDNumber{
			ID: numberID,
		}
	}

	stringID, ok := id.(string)
	if ok {
		return createdWithIDString{
			ID: stringID,
		}
	}

	return createdWithIDString{
		ID: types.ToString(id),
	}
}
