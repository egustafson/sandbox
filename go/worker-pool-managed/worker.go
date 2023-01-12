package main

import (
	"fmt"
	"sync/atomic"

	"github.com/Masterminds/log-go"
)

type WorkDescription struct {
	Scope int
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
	ID string
}

// static check:  simpleWorker implements Worker
var _ Worker = new(simpleWorker)

func SimpleWorkerFactory() (Worker, error) {
	id := simpleWorkerIDCounter.Add(1)
	w := &simpleWorker{
		ID: fmt.Sprintf("simple-worker-%d", id),
	}
	log.Debugf("worker created: %s", w.ID)
	return w, nil
}

func (w *simpleWorker) DoRequest(req *WorkRequest) {
	//
	// Simple worker:  just return a positive result immediately
	//
	req.ResponseCh <- &WorkResponse{
		Result: &WorkResult{true},
		Err:    nil,
	}
}

func (w *simpleWorker) IsHealthy() bool { return true }
func (w *simpleWorker) Close()          {}

// --  Mock Unstable Worker  -----------------------------------------
//

type mockUnstableWorker struct{}

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

func (w *mockUnstableWorker) IsHealthy() bool { return true }
func (w *mockUnstableWorker) Close()          {}
