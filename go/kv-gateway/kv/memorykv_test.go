package kv_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/egustafson/sandbox/go/kv-gateway/kv"
)

func TestKVTestSuite(t *testing.T) {
	suite.Run(t, new(KVTestSuite))
}

type KVTestSuite struct {
	suite.Suite
	kvs kv.KV
}

func (s *KVTestSuite) SetupTest() {
	s.kvs = kv.NewMemoryKV()
}

func (s *KVTestSuite) TearDownTest() {
	s.kvs.Close()
	s.kvs = nil
}

func (s *KVTestSuite) TestPutGetDelGet() {
	k := kv.Key([]byte("key"))
	v := kv.Value([]byte("value"))

	err := s.kvs.Put(k, v)
	s.Nil(err)

	val, err := s.kvs.Get(k)
	s.Nil(err)
	s.Equal(v, val)

	err = s.kvs.Del(k)
	s.Nil(err)

	_, err = s.kvs.Get(k)
	nskErr := kv.NoSuchKeyError(nil)
	s.ErrorAs(err, &nskErr)
}

func createKVList(prefix string, count int) (kvs []kv.KeyValue) {
	for ii := 0; ii < count; ii++ {
		k := fmt.Sprintf("%s-%d", prefix, ii)
		v := fmt.Sprintf("val-%s", k)
		kvs = append(kvs, kv.KeyValue{K: []byte(k), V: []byte(v)})
	}
	return kvs
}

func loadKV(kvs kv.KV, kvlist []kv.KeyValue) {
	for _, kv := range kvlist {
		kvs.Put(kv.K, kv.V)
	}
}

func (s *KVTestSuite) TestGetPrefix() {
	p := "key-b"
	loadKV(s.kvs, createKVList(p, 10))
	loadKV(s.kvs, createKVList("key-a", 10))
	loadKV(s.kvs, createKVList("key-c", 10))

	bKeys, err := s.kvs.GetPrefix(kv.Key([]byte(p)))
	s.Nil(err)
	s.True(len(bKeys) == 10)
	for _, kv := range bKeys {
		s.True(strings.HasPrefix(string(kv.K), p))
	}

}

func (s *KVTestSuite) TestDelPrefix() {
	p := "key-b"
	loadKV(s.kvs, createKVList(p, 10))
	loadKV(s.kvs, createKVList("key-a", 10))
	loadKV(s.kvs, createKVList("key-c", 10))

	bkeys, err := s.kvs.DelPrefix(kv.Key([]byte(p)))
	s.Nil(err)
	s.True(len(bkeys) == 10)
	for _, k := range bkeys {
		s.True(strings.HasPrefix(string(k), p))
	}

	remainingKeys, err := s.kvs.GetPrefix(kv.Key([]byte("")))
	s.Nil(err)
	for _, kv := range remainingKeys {
		s.False(strings.HasPrefix(string(kv.K), p))
	}
}
