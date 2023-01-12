package main

import (
	"context"
	"errors"
	"sync/atomic"

	"github.com/Masterminds/log-go"
)

type Pool interface {
	DoWork(ctx context.Context, wd *WorkDescription) (r *WorkResult, err error)
	//Acquire(ctx context.Context) (w Worker, err error)
	//Return(w Worker)
	Close()
}

type lazyInitWorkerPool struct {
	ID            string
	workerFactory WorkerFactory
	pool          chan Worker
	ctx           context.Context
	cancel        context.CancelFunc
	poolMax       int
	poolMin       int
	numWorkers    *atomic.Int32
	idleWorkers   *atomic.Int32
}

// static check: simplePool implements Pool
var _ Pool = new(lazyInitWorkerPool)

func NewLazyInitWorkerPool(factory WorkerFactory, id string, max, min int) Pool {
	p := &lazyInitWorkerPool{
		ID:            id,
		workerFactory: factory,
		pool:          make(chan Worker, max+1),
		poolMax:       max,
		poolMin:       min,
		numWorkers:    new(atomic.Int32),
		idleWorkers:   new(atomic.Int32),
	}
	p.ctx, p.cancel = context.WithCancel(context.Background())
	p.numWorkers.Store(0)
	p.idleWorkers.Store(0)

	//
	// TODO: do something to ensure there are workers in the pool
	//
	log.Infof("pool created: %s", p.ID)
	return p
}

func (p *lazyInitWorkerPool) Close() {
	p.cancel()
	// TODO: complete Close() implementation
}

func (p *lazyInitWorkerPool) DoWork(ctx context.Context, wd *WorkDescription) (r *WorkResult, err error) {
	if p.ctx.Err() != nil {
		return nil, p.ctx.Err()
	}

	w, err := p.acquireWorker(ctx)
	if err != nil {
		return nil, err
	}

	req := &WorkRequest{
		Work:       wd,
		ResponseCh: make(chan *WorkResponse, 1),
	}
	go func() {
		defer p.returnWorker(w)
		w.DoRequest(req)
	}()

	select {
	case resp, ok := <-req.ResponseCh:
		if !ok {
			return nil, errors.New("result canceled: response chan closed")
		}
		return resp.Result, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// congested returns true if the pool is in use "enough" that workers above the
// low-water mark should be retained.  When the pool is _not_ congested then
// returning a worker to a pool that already has at least low-water-mark workers
// in it will cause that worker to be decommissioned.  This function is a
// heuristic; treat it as such.
func (p *lazyInitWorkerPool) congested() bool {
	return false // <-- the 'dumb' heuristic.
}

func (p *lazyInitWorkerPool) acquireWorker(ctx context.Context) (w Worker, err error) {
	// Logic:  (not in implementation order)
	// if worker available in queue then pull it and use it.
	// if numWorkers < highWaterWorkers then allocate another and use it.
	// block waiting for a worker to be returned to the pool.

	for {
		// if the pool is empty and we're not at the max allowed workers
		if int(p.idleWorkers.Load()) < 1 && int(p.numWorkers.Load()) < p.poolMax {
			p.addNewWorker()
		}

		var ok bool
		select {
		case w, ok = <-p.pool:
			if !ok {
				return nil, errors.New("worker pool closed")
			}
			p.idleWorkers.Add(-1)
		case <-ctx.Done():
			return nil, ctx.Err()
		}

		if w.IsHealthy() {
			return w, nil // <-- happy path
		}
		go func() {
			p.decommissionWorker(w)
		}()
	}
}

func (p *lazyInitWorkerPool) returnWorker(w Worker) {

	// if idleWorkers < lowWaterWorkers then return worker to pool
	// if "congested" then return worker to pool
	// retire worker and down active pool size

	if !w.IsHealthy() {
		p.decommissionWorker(w)
	}
	if !p.congested() && p.idleWorkers.Load() >= int32(p.poolMin) {
		p.decommissionWorker(w)
	}

	p.idleWorkers.Add(1)
	p.pool <- w
}

// addNewWorker attempts to add a new worker to the pool, IIF the pool would not
// go over the limit of poolMax.  This function is intended to be called
// concurrently by multiple threads of execution and guarantees that the pool
// will never go over limit.
func (p *lazyInitWorkerPool) addNewWorker() {
	if p.ctx.Err() != nil {
		return
	}
	if p.numWorkers.Load() >= int32(p.poolMax) {
		return // someone beat us to adding a worker, we're at max
	}

	w, err := p.workerFactory()
	if err != nil {
		// if not an example, do something to count or indicate a failure
		return // <-- don't loop trying to create, someone else will try later
	}

	numWorkers := p.numWorkers.Load() // establish the current truth
	if numWorkers >= int32(p.poolMax) {
		w.Close() // someone beat us to the max limit, throw this one away
		return
	}
	if p.numWorkers.CompareAndSwap(numWorkers, numWorkers+1) {
		// success!  add the worker to the pool, the count has been increased
		p.idleWorkers.Add(1)
		p.pool <- w
	} else {
		// unsuccessful add, throw away worker, someone else will try later
		w.Close()
	}
}

func (p *lazyInitWorkerPool) decommissionWorker(w Worker) {
	p.numWorkers.Add(-1)
	w.Close()
}
