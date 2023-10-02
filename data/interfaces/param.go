package interfaces

// Param is interface which returns after getting QueryParam or Param from Context
type Param interface {
	// String returns param value in string
	String() string
	// Int returns param value in int
	Int() int
	// Bool returns param value in boolean
	Bool() bool
}
