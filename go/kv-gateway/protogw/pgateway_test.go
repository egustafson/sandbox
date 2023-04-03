package protogw_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/egustafson/sandbox/go/kv-gateway/kv"
	"github.com/egustafson/sandbox/go/kv-gateway/protogw"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

func TestProtoGwTestSuite(t *testing.T) {
	suite.Run(t, new(ProtoGwTestSuite))
}

type ProtoGwTestSuite struct {
	suite.Suite
	kvdb    kv.KV
	protoGw protogw.ProtoGateway
}

func (s *ProtoGwTestSuite) SetupTest() {
	s.kvdb = kv.NewMemoryKV()
	s.protoGw = protogw.MakeProtoGateway("test-proto-gw-test-suite", s.kvdb)
}

func (s *ProtoGwTestSuite) TearDownTest() {
	s.protoGw.Close()
	s.kvdb = nil
	s.protoGw = nil
}

func (s *ProtoGwTestSuite) TestProtoGateway_Close() {

	s.protoGw.Close()

	_, err := s.kvdb.GetPrefix(kv.Key(string("")))
	closedError := kv.ClosedError(nil)
	s.ErrorAs(err, &closedError)
}

func (s *ProtoGwTestSuite) TestProtoGateway_GetProducer() {
	p := protogw.Producer{
		ID:    uuid.New(),
		Name:  "test-producer",
		Value: 123,
	}

	err := s.protoGw.SetProducer(p)
	s.Nil(err)

	result, err := s.protoGw.GetProducer(p.ID)
	if s.Nil(err) {
		s.Equal(p, result)
	}
}

func (s *ProtoGwTestSuite) TestProtoGateway_GetProducers() {
	producers := make(map[int]protogw.Producer)
	for i := 0; i < 10; i++ {
		p := protogw.Producer{
			ID:    uuid.New(),
			Name:  fmt.Sprintf("test-producer-%d", i),
			Value: i,
		}
		producers[i] = p
		err := s.protoGw.SetProducer(p)
		s.Nil(err)
	}

	results, err := s.protoGw.GetProducers()
	if s.Nil(err) {
		s.Equal(len(producers), len(results))
		for _, r := range results {
			p, ok := producers[r.Value]
			if s.True(ok) {
				s.Equal(p, r)
			}
		}
	}
}

func (s *ProtoGwTestSuite) TestProtoGateway_GetAllLogs() {
	logs := make(map[string]protogw.LogRecord)
	for i := 0; i < 10; i++ {
		lr := protogw.LogRecord{
			TimeStamp: time.Now().UTC(),
			Producer:  uuid.New(),
			Message:   fmt.Sprintf("log-message-%d", i),
		}
		logs[lr.Message] = lr
		err := s.protoGw.SendLogRecord(lr)
		s.Nil(err)
	}

	results, err := s.protoGw.GetAllLogs()
	if s.Nil(err) {
		s.Equal(len(logs), len(results))
		for _, r := range results {
			lr, ok := logs[r.Message]
			if s.True(ok) {
				s.Equal(lr, r)
			}
		}
	}
}
