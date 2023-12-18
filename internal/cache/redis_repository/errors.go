package redis_repository

import (
	"github.com/lowl11/boost/internal/boosties/errors"
)

func ErrorURLRequired() error {
	return errors.
		New("URL is required").
		SetType("Redis_URLRequired")
}

func ErrorRedisUnknownType(key, redisType string) error {
	return errors.
		New("Unknown Redis type").
		SetType("Redis_UnknownType").
		AddContext("redis_type", redisType).
		AddContext("key", key)
}

func ErrorRedisPing(err error) error {
	return errors.
		New("Ping Redis error").
		SetType("Redis_Ping").
		SetError(err)
}

func ErrorRedisGetType(err error) error {
	return errors.
		New("Get key type error").
		SetType("Redis_GetType").
		SetError(err)
}

func ErrorRedisParseValue(err error) error {
	return errors.
		New("Redis parse value error").
		SetType("Redis_ParseValue").
		SetError(err)
}

func ErrorSearchKeys(pattern string, err error) error {
	return errors.
		New("Search keys by pattern error").
		SetType("Redis_SearchKeys").
		SetError(err).
		AddContext("pattern", pattern)
}

func ErrorGetAllKeys(err error) error {
	return errors.
		New("Get all keys error").
		SetType("Redis_GetAllKeys").
		SetError(err)
}

func ErrorGetCacheByKey(key string, err error) error {
	return errors.
		New("Get cache by key error").
		SetType("Redis_GetCacheByKey").
		SetError(err).
		AddContext("key", key)
}

func ErrorSetCache(key string, value any, err error) error {
	return errors.
		New("Set cache error").
		SetType("Redis_SetCache").
		SetError(err).
		SetContext(map[string]any{
			"key":   key,
			"value": value,
		})
}

func ErrorDeleteCache(key string, err error) error {
	return errors.
		New("Delete cache error").
		SetType("Redis_DeleteCache").
		SetError(err).
		AddContext("key", key)
}
