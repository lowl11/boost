package text

import "github.com/mehanizm/iuliia-go"

func CyrillicToCode(text string) string {
	return iuliia.Wikipedia.Translate(text)
}
