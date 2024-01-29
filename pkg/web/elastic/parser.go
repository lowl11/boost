package elastic

import (
	"github.com/google/uuid"
	"github.com/lowl11/boost/errors"
	"github.com/lowl11/boost/log"
	"github.com/lowl11/flex"
	"reflect"
	"strings"
)

type mappingField struct {
	Type       string                  `json:"type"`
	Properties map[string]mappingField `json:"properties,omitempty"`
}

func parseObject(object any) (map[string]mappingField, error) {
	fxType := flex.Type(reflect.TypeOf(object))
	fxType.Reset(fxType.Unwrap())
	if !fxType.IsStruct() {
		return nil, errors.New("Given object is not struct").
			SetType("ELK_ObjectIsNotStruct")
	}

	fxObject := flex.Struct(object)
	fields := fxObject.FieldsRow()

	mappings := make(map[string]mappingField)
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

				props, err := parseObject(reflect.New(fxFieldType.Type()).Elem().Interface())
				if err != nil {
					log.Error(err, "Parse nested document error")
					continue
				}

				if len(props) == 0 {
					log.Error(errors.New("No props for nested document error"))
					continue
				}

				mappings[name[0]] = mappingField{
					Type:       "nested",
					Properties: props,
				}

				continue
			}
		}

		mappings[name[0]] = mappingField{
			Type: convertTypeToMapping(field.Type, name),
		}
	}

	return mappings, nil
}

func getID(object any) (string, error) {
	fxType := flex.Type(reflect.TypeOf(object))
	fxType.Reset(fxType.Unwrap())

	if !fxType.IsStruct() {
		return "", errors.New("Given object is not struct").
			SetType("ELK_ObjectIsNotStruct")
	}

	id, ok := reflect.ValueOf(object).FieldByName("ID").Interface().(uuid.UUID)
	if ok {
		return id.String(), nil
	}

	return "", nil
}

func convertTypeToMapping(t reflect.Type, tags []string) string {
	for _, tag := range tags {
		if strings.Contains(tag, "custom") {
			_, after, found := strings.Cut(tag, ":")
			if found && len(after) > 0 {
				return after
			}
		}
	}

	fxType := flex.Type(t)
	fxType.Reset(fxType.Unwrap())

	switch fxType.Type().String() {
	case "time.Time":
		return "date"
	case "uuid.UUID":
		return "text"
	}

	switch fxType.Type().Kind() {
	case reflect.Bool:
		return "boolean"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128:
		return "integer"
	default:
		return "text"
	}
}
