package elk_parser

import (
	"github.com/lowl11/flex"
	"reflect"
	"strings"
)

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
