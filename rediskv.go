package gokv

// Redis string, key
var key = "k"

// redis list
var lists = "l"

// redis hash
var hash = "h"

// redis set
var set = "s"

// redis sorted set
var sSet = "ss"

type RedisIns interface {
	GetKey(arr ...string) string
	GetListKey(arr ...string) string
	GetHashKey(arr ...string) string
	GetSetKey(arr ...string) string
	GetSSetKey(arr ...string) string
}

type RedisKv struct {
	Kv
}

func (r RedisKv) GetKey(arr ...string) string {
	return r.CompileKey(arr, key)
}

func (r RedisKv) GetListKey(arr ...string) string {
	return r.CompileKey(arr, lists)
}

func (r RedisKv) GetHashKey(arr ...string) string {
	return r.CompileKey(arr, hash)
}

func (r RedisKv) GetSetKey(arr ...string) string {
	return r.CompileKey(arr, set)
}

func (r RedisKv) GetSSetKey(arr ...string) string {
	return r.CompileKey(arr, sSet)
}
