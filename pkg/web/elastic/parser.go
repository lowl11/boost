package elastic

import (
	"github.com/google/uuid"
	"github.com/lowl11/boost/errors"
	"github.com/lowl11/boost/log"
	"github.com/lowl11/boost/pkg/io/flex"
	"reflect"
	"strings"
)

type mappingField struct {
	Type       string                  `json:"type"`
	Properties map[string]mappingField `json:"properties,omitempty"`
}

func parseObject(object any) (map[string]mappingField, error) {
	if !flex.Type(object).Unwrap().IsStruct() {
		return nil, errors.New("Given object is not struct").
			SetType("ELK_ObjectIsNotStruct")
	}

	fxObject, err := flex.Struct(object)
	if err != nil {
		return nil, err
	}

	fields := fxObject.FieldsRow()

	mappings := make(map[string]mappingField)
	for _, field := range fields {
		name := field.Tag("json")
		if len(name) == 0 {
			continue
		}

		fieldName := field.Type().String()
		fxFieldType := flex.Type(field.Type())
		if fxFieldType.IsSlice() {
			fxFieldType = flex.Type(field.Type().Elem())
		}

		if fieldName != "time.Time" && fieldName != "uuid.UUID" {
			if !fxFieldType.IsPrimitive() && (fxFieldType.IsStruct() || fxFieldType.IsSlice()) && fieldName != "time.Time" && fieldName != "uuid.UUID" {
				fxFieldType = flex.Type(field.Type().Elem())

				props, err := parseObject(reflect.New(fxFieldType.Type()).Elem().Interface())
				if err != nil {
					log.Error("Parse nested document error:", err)
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
			Type: convertTypeToMapping(flex.Type(field.Type), name),
		}
	}

	return mappings, nil
}

func getID(object any) (string, error) {
	if !flex.Type(object).Unwrap().IsStruct() {
		return "", errors.New("Given object is not struct").
			SetType("ELK_ObjectIsNotStruct")
	}

	id, ok := reflect.ValueOf(object).FieldByName("ID").Interface().(uuid.UUID)
	if ok {
		return id.String(), nil
	}

	return "", nil
}

func convertTypeToMapping(t flex.ObjectType, tags []string) string {
	for _, tag := range tags {
		if strings.Contains(tag, "custom") {
			_, after, found := strings.Cut(tag, ":")
			if found && len(after) > 0 {
				return after
			}
		}
	}

	if t.IsTime() {
		return "date"
	} else if t.IsUUID() {
		return "text"
	} else if t.IsBool() {
		return "boolean"
	} else if t.IsNumeric() {
		return "integer"
	}

	return "text"
}
