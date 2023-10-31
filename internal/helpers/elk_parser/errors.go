package elk_parser

import "github.com/lowl11/boost/data/errors"

func ErrorObjectIsNotStruct() error {
	return errors.
		New("Given object is not struct").
		SetType("ELK_ObjectIsNotStruct")
}
