package kv

import (
	"bytes"
	"context"
	"fmt"
	"sort"
	"strings"
	"sync"
)

type memoryKV struct {
	kv     map[string]Value
	lock   sync.RWMutex
	alive  context.Context
	cancel context.CancelFunc
}

var _ KV = (*memoryKV)(nil)

func NewMemoryKV() KV {
	ctx, cancel := context.WithCancel(context.Background())
	return &memoryKV{
		kv:     make(map[string]Value),
		alive:  ctx,
		cancel: cancel,
	}
}

func (mkv *memoryKV) Close() {
	mkv.lock.Lock()
	defer mkv.lock.Unlock()
	mkv.cancel()
	mkv.kv = nil
}

type sortedkeys []string

func (k sortedkeys) Len() int           { return len(k) }
func (k sortedkeys) Swap(i, j int)      { k[i], k[j] = k[j], k[i] }
func (k sortedkeys) Less(i, j int) bool { return string(k[i]) < string(k[j]) }

func (mkv *memoryKV) Dump() string {
	mkv.lock.RLock()
	defer mkv.lock.RUnlock()
	// Dump ignores the alive context; it's a debug function
	//
	keys := make([]string, 0, len(mkv.kv))
	for k := range mkv.kv {
		keys = append(keys, k)
	}
	sort.Sort(sortedkeys(keys))

	var strBuf bytes.Buffer
	for _, k := range keys {
		strBuf.WriteString(fmt.Sprintf("%s %s\n", k, string(mkv.kv[k])))
	}
	return strBuf.String()
}

func (mkv *memoryKV) Put(k Key, v Value) error {
	mkv.lock.Lock()
	defer mkv.lock.Unlock()
	if mkv.alive.Err() != nil {
		return closedError()
	}
	mkv.kv[string(k)] = v
	return nil
}

func (mkv *memoryKV) Get(k Key) (v Value, err error) {
	mkv.lock.RLock()
	defer mkv.lock.RUnlock()
	if mkv.alive.Err() != nil {
		return nil, closedError()
	}
	var ok bool
	v, ok = mkv.kv[string(k)]
	if !ok {
		return nil, noSuchKeyError()
	}
	return v, nil
}

func (mkv *memoryKV) GetPrefix(k Key) (kvs []KeyValue, err error) {
	mkv.lock.RLock()
	defer mkv.lock.RUnlock()
	if mkv.alive.Err() != nil {
		return nil, closedError()
	}
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
	mkv.lock.Lock()
	defer mkv.lock.Unlock()
	if mkv.alive.Err() != nil {
		return closedError()
	}
	if _, ok := mkv.kv[string(k)]; !ok {
		return noSuchKeyError()
	}
	delete(mkv.kv, string(k))
	return nil
}

func (mkv *memoryKV) DelPrefix(k Key) (keys []Key, err error) {
	mkv.lock.Lock()
	defer mkv.lock.Unlock()
	if mkv.alive.Err() != nil {
		return nil, closedError()
	}
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
