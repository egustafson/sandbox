package main

import (
	"fmt"
	"sync/atomic"
	"time"

	"github.com/Masterminds/log-go"
)

type WorkDescription struct {
	Effort int
}

type WorkResult struct {
	Success bool
}

type WorkRequest struct {
	Work       *WorkDescription
	ResponseCh chan *WorkResponse
	//
	// TODO: add more meaningful fields
}
type WorkResponse struct {
	Result *WorkResult
	Err    error
	//
	// TODO: add more meaningful fields
}

type Worker interface {
	ID() string
	DoRequest(req *WorkRequest) // result passed through req.ResultCh
	IsHealthy() bool
	Close()
}

type WorkerFactory func() (Worker, error)

// --  Simple Worker  ------------------------------------------------
//

var simpleWorkerIDCounter = new(atomic.Uint32)

func init() {
	simpleWorkerIDCounter.Store(0)
}

type simpleWorker struct {
	id string
}

// static check:  simpleWorker implements Worker
var _ Worker = new(simpleWorker)

func SimpleWorkerFactory() (Worker, error) {
	id := simpleWorkerIDCounter.Add(1)
	w := &simpleWorker{
		id: fmt.Sprintf("simple-worker-%d", id),
	}
	log.Debugf("worker created: %s", w.ID())
	return w, nil
}

func (w *simpleWorker) DoRequest(req *WorkRequest) {

	// delay for a short time
	<-time.After(time.Duration(req.Work.Effort) * time.Millisecond)

	req.ResponseCh <- &WorkResponse{
		Result: &WorkResult{true},
		Err:    nil,
	}
}

func (w *simpleWorker) ID() string      { return w.id }
func (w *simpleWorker) IsHealthy() bool { return true }
func (w *simpleWorker) Close() {
	log.Debugf("worker closed: %s", w.ID())
}

// --  Mock Unstable Worker  -----------------------------------------
//

type mockUnstableWorker struct {
	id string
}

// static check:  mockUnstableWorker implements Worker
var _ Worker = new(mockUnstableWorker)

func MockUnstableWorkerFactory() (Worker, error) {
	w := &mockUnstableWorker{}
	return w, nil
}

func (w *mockUnstableWorker) DoRequest(req *WorkRequest) {
	req.ResponseCh <- &WorkResponse{
		Result: &WorkResult{true},
		Err:    nil,
	}
}

func (w *mockUnstableWorker) ID() string      { return w.id }
func (w *mockUnstableWorker) IsHealthy() bool { return true }
func (w *mockUnstableWorker) Close()          {}
