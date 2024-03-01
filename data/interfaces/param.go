package interfaces

import "github.com/google/uuid"

// Param is interface which returns after getting QueryParam or Param from Context
type Param interface {
	// String returns param value in string
	String() string

	// Strings return slice of param values
	Strings() []string

	// Bool returns param value in boolean
	Bool() bool

	// Int returns param value in int
	Int() (int, error)
	// MustInt returns param value in int, if throws error, will return 0
	MustInt() int

	// UUID returns parsed UUID value
	UUID() (uuid.UUID, error)
	// MustUUID returns parsed UUID value, if throws error will return empty [16]byte{}
	MustUUID() uuid.UUID

	Bytes() []byte
}
