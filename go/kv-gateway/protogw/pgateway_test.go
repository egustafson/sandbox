package protogw_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/egustafson/sandbox/go/kv-gateway/protogw"
	"github.com/egustafson/werks/kv"
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

func (s *ProtoGwTestSuite) TestProtoGateway_GetComponent() {
	c := protogw.Component{
		ID:   uuid.New(),
		Name: "test-component",
	}

	err := s.protoGw.SetComponent(c)
	s.Nil(err)

	result, err := s.protoGw.GetComponent(c.ID)
	if s.Nil(err) {
		s.Equal(c, result)
	}
}

func (s *ProtoGwTestSuite) TestProtoGateway_GetAllComponents() {
	components := make(map[string]protogw.Component)
	for i := 0; i < 10; i++ {
		p := protogw.Component{
			ID:   uuid.New(),
			Name: fmt.Sprintf("test-component-%d", i),
		}
		components[p.Name] = p
		err := s.protoGw.SetComponent(p)
		s.Nil(err)
	}

	results, err := s.protoGw.GetAllComponents()
	if s.Nil(err) {
		s.Equal(len(components), len(results))
		for _, r := range results {
			p, ok := components[r.Name]
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
			Sender:    uuid.New(),
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

func observeLogOrTimeout(logCh <-chan protogw.LogRecord) (protogw.LogRecord, bool) {
	select {
	case lr := <-logCh:
		return lr, true
	case <-time.After(time.Millisecond):
		return protogw.LogRecord{}, false
	}
}

func (s *ProtoGwTestSuite) TestProtoGateway_ObserveLogs() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	logWatcher, err := s.protoGw.ObserveLogs(ctx)
	s.Nil(err)

	logUUID := uuid.New()
	for i := 0; i < 10; i++ {
		sendLR := protogw.LogRecord{
			TimeStamp: time.Now().UTC(),
			Sender:    logUUID,
			Message:   fmt.Sprintf("message-%d", i),
		}
		err = s.protoGw.SendLogRecord(sendLR)
		s.Nil(err)

		recvLR, ok := observeLogOrTimeout(logWatcher)
		if s.True(ok) {
			s.Equal(sendLR, recvLR)
		}
	}
}

func observeGaugeOrTimeout(gCh <-chan protogw.GaugeState) (int, bool) {
	select {
	case gauge := <-gCh:
		return gauge.Value, true
	case <-time.After(time.Millisecond):
		return 0, false
	}
}

func (s *ProtoGwTestSuite) TestProtoGateway_Gauge() {
	gID := uuid.New()
	err := s.protoGw.CreateGauge(gID, "gauge-under-test") // validate CreateGauge
	s.Nil(err)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	gWatcher, err := s.protoGw.ObserveGauge(ctx, gID) // validate ObserveGauge
	s.Nil(err)

	for i := 10; i < 20; i++ {
		err = s.protoGw.SetGauge(gID, i) // validate SetGauge
		s.Nil(err)

		r, ok := observeGaugeOrTimeout(gWatcher) // validate observer respondes
		s.True(ok)
		s.Equal(i, r)
	}

	// Validate that observers buffer (some) in order.
	for i := 50; i < 55; i++ {
		err = s.protoGw.SetGauge(gID, i) // validate SetGauge
		s.Nil(err)
	}

	for i := 50; i < 55; i++ {
		r, ok := observeGaugeOrTimeout(gWatcher) // validate observer respondes
		s.True(ok)
		s.Equal(i, r)
	}
}
