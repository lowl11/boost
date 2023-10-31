package elk_parser

import (
	"github.com/google/uuid"
	"github.com/lowl11/flex"
	"reflect"
)

type MappingField struct {
	Type string `json:"type"`
}

func ParseObject(object any) (map[string]MappingField, error) {
	fxType := flex.Type(reflect.TypeOf(object))
	fxType.Reset(fxType.Unwrap())
	if !fxType.IsStruct() {
		return nil, ErrorObjectIsNotStruct()
	}

	fxObject := flex.Struct(object)
	fields := fxObject.FieldsRow()

	mappings := make(map[string]MappingField)
	for _, field := range fields {
		name := flex.Field(field).Tag("elk")
		if len(name) == 0 {
			continue
		}

		mappings[name[0]] = MappingField{
			Type: convertTypeToMapping(field.Type),
		}
	}

	return mappings, nil
}

func GetID(object any) (string, error) {
	fxType := flex.Type(reflect.TypeOf(object))
	fxType.Reset(fxType.Unwrap())

	if !fxType.IsStruct() {
		return "", ErrorObjectIsNotStruct()
	}

	id, ok := reflect.ValueOf(object).FieldByName("ID").Interface().(uuid.UUID)
	if ok {
		return id.String(), nil
	}

	return "", nil
}
