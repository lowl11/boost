package redis_repository

import (
	"github.com/lowl11/boost/internal/boosties/errors"
	"net/http"
)

const (
	typeErrorURLRequired      = "URLRequired"
	typeErrorUnknownValueType = "RedisUnknownType"
	typeErrorPing             = "RedisPing"
	typeErrorParseValue       = "RedisParseValue"
	typeErrorSearchKeys       = "SearchKeys"
	typeErrorGetType          = "RedisGetType"
	typeErrorGetAllKeys       = "GetAllKeys"
	typeErrorGetCacheByKey    = "GetCacheByKey"
	typeErrorSetCache         = "SetCache"
	typeErrorDeleteCache      = "DeleteCache"
)

func ErrorURLRequired() error {
	return errors.
		New("URL is required").
		SetType(typeErrorURLRequired).
		SetHttpCode(http.StatusInternalServerError)
}

func ErrorRedisUnknownType(key, redisType string) error {
	return errors.
		New("Unknown Redis type").
		SetType(typeErrorUnknownValueType).
		SetHttpCode(http.StatusInternalServerError).
		AddContext("redis_type", redisType).
		AddContext("key", key)
}

func ErrorRedisPing(err error) error {
	return errors.
		New("Ping Redis error").
		SetType(typeErrorPing).
		SetHttpCode(http.StatusInternalServerError).
		SetError(err)
}

func ErrorRedisGetType(err error) error {
	return errors.
		New("Get key type error").
		SetType(typeErrorGetType).
		SetHttpCode(http.StatusInternalServerError).
		SetError(err)
}

func ErrorRedisParseValue(err error) error {
	return errors.
		New("Redis parse value error").
		SetType(typeErrorParseValue).
		SetHttpCode(http.StatusInternalServerError).
		SetError(err)
}

func ErrorSearchKeys(pattern string, err error) error {
	return errors.
		New("Search keys by pattern error").
		SetType(typeErrorSearchKeys).
		SetHttpCode(http.StatusInternalServerError).
		SetError(err).
		AddContext("pattern", pattern)
}

func ErrorGetAllKeys(err error) error {
	return errors.
		New("Get all keys error").
		SetType(typeErrorGetAllKeys).
		SetHttpCode(http.StatusInternalServerError).
		SetError(err)
}

func ErrorGetCacheByKey(key string, err error) error {
	return errors.
		New("Get cache by key error").
		SetType(typeErrorGetCacheByKey).
		SetHttpCode(http.StatusInternalServerError).
		SetError(err).
		AddContext("key", key)
}

func ErrorSetCache(key string, value any, err error) error {
	return errors.
		New("Set cache error").
		SetType(typeErrorSetCache).
		SetHttpCode(http.StatusInternalServerError).
		SetError(err).
		SetContext(map[string]any{
			"key":   key,
			"value": value,
		})
}

func ErrorDeleteCache(key string, err error) error {
	return errors.
		New("Delete cache error").
		SetType(typeErrorDeleteCache).
		SetHttpCode(http.StatusInternalServerError).
		SetError(err).
		AddContext("key", key)
}
