package context

import "errors"

func ErrorUnknownContentType(contentType string) error {
	return errors.New("unknown content-type: " + contentType)
}
