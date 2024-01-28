package elk_parser

import (
	"github.com/google/uuid"
	"github.com/lowl11/boost/log"
	"github.com/lowl11/flex"
	"reflect"
)

type MappingField struct {
	Type       string                  `json:"type"`
	Properties map[string]MappingField `json:"properties,omitempty"`
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
		name := flex.Field(field).Tag("json")
		if len(name) == 0 {
			continue
		}

		if flex.Type(field.Type).IsStruct() {
			props, err := ParseObject(reflect.New(field.Type).Interface())
			if err != nil {
				log.Error(err, "Parse nested document error")
				continue
			}

			mappings[name[0]] = MappingField{
				Type:       "nested",
				Properties: props,
			}

			continue
		}

		mappings[name[0]] = MappingField{
			Type: convertTypeToMapping(field.Type, name),
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
