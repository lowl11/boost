package domain

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

	return createdWithIDString{
		ID: id.(string),
	}
}
