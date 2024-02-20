package storages

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func ToString(anyValue any) string {
	if anyValue == nil {
		return ""
	}

	if _, ok := anyValue.(error); ok {
		return anyValue.(error).Error()
	}

	isNumber := func(value rune) bool {
		switch value {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			return true
		}
		return false
	}

	value := reflect.ValueOf(anyValue)

	switch value.Kind() {
	case reflect.String:
		stringValue := anyValue.(string)

		// if this is counter variable like $1
		if len(stringValue) > 1 && stringValue[0] == '$' && isNumber(rune(stringValue[1])) {
			return stringValue
		}

		// if this is variable like :id
		if stringValue != "" && stringValue[0] == ':' {
			return stringValue
		}

		// if this is variable for join fields
		if stringValue != "" && stringValue[0] == '$' {
			return stringValue[1:]
		}

		valueString := strings.Builder{}
		valueString.WriteString("'")
		if strings.Contains(stringValue, "'") {
			stringValue = strings.ReplaceAll(stringValue, "'", "''")
		}
		valueString.WriteString(stringValue)
		valueString.WriteString("'")
		return valueString.String()
	case reflect.Bool:
		return strconv.FormatBool(anyValue.(bool))
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(value.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(value.Uint(), 10)
	case reflect.Float32:
		return fmt.Sprintf("%f", value.Float())
	case reflect.Float64:
		return fmt.Sprintf("%g", value.Float())
	case reflect.Struct, reflect.Map, reflect.Slice, reflect.Array:
		valueInBytes, err := json.Marshal(anyValue)
		if err != nil {
			return ""
		}
		return strings.ReplaceAll(string(valueInBytes), "\"", "'")
	case reflect.Ptr:
		if value.IsZero() || value.Elem().IsZero() {
			return "NULL"
		}

		return ToString(value.Elem().Interface())
	default:
		return fmt.Sprintf("%v", value)
	}
}
