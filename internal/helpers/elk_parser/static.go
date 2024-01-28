package elk_parser

import (
	"github.com/google/uuid"
	"github.com/lowl11/boost/internal/boosties/errors"
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

		fieldName := field.Type.String()
		fxFieldType := flex.Type(field.Type)
		if fxFieldType.IsSlice() {
			fxFieldType = flex.Type(field.Type.Elem())
		}

		if fieldName != "time.Time" && fieldName != "uuid.UUID" {
			if !fxFieldType.IsPrimitive() && (fxFieldType.IsStruct() || fxFieldType.IsSlice()) && fieldName != "time.Time" && fieldName != "uuid.UUID" {
				fxFieldType = flex.Type(field.Type.Elem())

				props, err := ParseObject(reflect.New(fxFieldType.Type()).Elem().Interface())
				if err != nil {
					log.Error(err, "Parse nested document error")
					continue
				}

				if len(props) == 0 {
					log.Error(errors.New("No props for nested document error"))
					continue
				}

				mappings[name[0]] = MappingField{
					Type:       "nested",
					Properties: props,
				}

				continue
			}
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
