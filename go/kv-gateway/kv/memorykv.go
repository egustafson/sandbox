package kv

import "strings"

type memoryKV struct {
	kv map[string]Value
}

var _ KV = (*memoryKV)(nil)

func NewMemoryKV() KV {
	return &memoryKV{
		kv: make(map[string]Value),
	}
}

func (mkv *memoryKV) Close() {
	mkv.kv = nil
}

func (mkv *memoryKV) Put(k Key, v Value) error {
	mkv.kv[string(k)] = v
	return nil
}

func (mkv *memoryKV) Get(k Key) (v Value, err error) {
	var ok bool
	v, ok = mkv.kv[string(k)]
	if !ok {
		return nil, noSuchKeyError()
	}
	return v, nil
}

func (mkv *memoryKV) GetPrefix(k Key) (kvs []KeyValue, err error) {
	prefix := string(k)
	kvs = make([]KeyValue, 0)
	for k, v := range mkv.kv {
		if strings.HasPrefix(k, prefix) {
			kvs = append(kvs, KeyValue{K: []byte(k), V: v})
		}
	}
	return kvs, nil
}

func (mkv *memoryKV) Del(k Key) (err error) {
	if _, ok := mkv.kv[string(k)]; !ok {
		return noSuchKeyError()
	}
	delete(mkv.kv, string(k))
	return nil
}

func (mkv *memoryKV) DelPrefix(k Key) (keys []Key, err error) {
	prefix := string(k)
	keys = make([]Key, 0)
	for k, _ := range mkv.kv {
		if strings.HasPrefix(k, prefix) {
			keys = append(keys, []byte(k))
			delete(mkv.kv, k)
		}
	}
	return keys, nil
}
