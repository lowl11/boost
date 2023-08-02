package interfaces

import "time"

type CacheRepository interface {
	All() map[string]any
	Set(string, any, ...time.Duration)
	Get(string) any
	Delete(string)
}
