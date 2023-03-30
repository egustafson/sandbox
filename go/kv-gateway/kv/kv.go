package kv

import "errors"

type Key []byte
type Value []byte

type KeyValue struct {
	K Key
	V Value
}

type KV interface {
	Close()
	Put(k Key, v Value) error
	Get(k Key) (v Value, err error)
	GetPrefix(k Key) (kvs []KeyValue, err error)
	Del(k Key) (err error)
	DelPrefix(k Key) (keys []Key, err error)
}

type NoSuchKeyError interface{ error }

func noSuchKeyError() NoSuchKeyError {
	e := errors.New("no such key")
	return NoSuchKeyError(e)
}
