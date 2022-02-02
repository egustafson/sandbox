package main

import (
	"context"
	"log"
	"time"
)

// This package implements the components to implement the pattern
// described in the README.md
//
//  * Job Struct
//  * Work Queue
//  * Worker Pool (Queue)
//  * Worker "object"
//  * Dispatcher

type JobReq struct {
	InputData  int
	Delay      time.Duration
	Ctx        context.Context
	ResultChan chan JobResp // unbuffered
}

type JobResp struct {
	Result int
	Err    error
}

type Worker struct {
	ID         int
	StartTime  time.Time
	workerPool WorkerPoolT
}

type WorkerPoolT chan *Worker
type JobQueueT chan *JobReq

const (
	JobQueueSize   = 10 // intentionally on the small size
	WorkerPoolSize = 3
	MaxWorkerAge   = time.Second // retire old workers
)

var (
	WorkerPool   WorkerPoolT = make(WorkerPoolT, WorkerPoolSize)
	JobQueue     JobQueueT   = make(JobQueueT, JobQueueSize)
	nextWorkerID int         = 1 // there's a race on this variable -- this is a demo
)

func init() {
	// initial population of worker pool
	for ii := 0; ii < WorkerPoolSize; ii++ {
		worker := &Worker{
			ID:         nextWorkerID, // RACE CONDITION -- this is a demo
			StartTime:  time.Now(),
			workerPool: WorkerPool,
		}
		nextWorkerID = nextWorkerID + 1 // RACE CONDITION - this is a demo
		log.Printf("INFO:  adding worker(%d) to the pool", worker.ID)
		WorkerPool <- worker
	}
}

// DoWork implements the actual act of doing the work, returning the
// work and re-enqueuing oneself into the WorkerPool.
func (w *Worker) DoWork(j *JobReq) {
	if j.Ctx.Err() != nil { // context canceled
		j.ResultChan <- JobResp{
			Result: 0,
			Err:    j.Ctx.Err(),
		}
		close(j.ResultChan)
		return
	}
	log.Printf("INFO:  worker(%d) doing work for %v", w.ID, j.Delay)
	time.Sleep(j.Delay)
	//
	response := JobResp{ // "do work" and create response
		Result: j.InputData + w.ID,
		Err:    nil,
	}
	j.ResultChan <- response // deliver result
	close(j.ResultChan)      // and we're done with the channel
	w.workerPool <- w        // re-enqueue myself to do more work
}

func Dispatcher(ctx context.Context, wPool WorkerPoolT, jobQ JobQueueT) { // long running, background goroutine
	const (
		workerWaitTime = 100 * time.Millisecond // should not wait long
	)
	defer log.Print("Dispatcher exiting.")
	for ctx.Err() == nil { // loop forever (until context canceled)
		var worker *Worker
		// Get a worker
		select {
		case worker = <-wPool:
		case <-time.After(workerWaitTime):
			log.Print("WARN:  timeout waiting on worker")
			continue // just enter the loop again
		}
		// Check worker's age
		if time.Now().Sub(worker.StartTime) > MaxWorkerAge {
			Retire(worker)
			continue // loop and get a fresh worker
		}
		// Get some work, and do it
		select {
		case job := <-JobQueue:
			worker.DoWork(job) // fire and forget
		case <-time.After(workerWaitTime):
			log.Printf("DBUG:  timeout waiting for work, returning worker(%d) to pool", worker.ID)
			wPool <- worker
		}
		// and do it all over again
	}
}

func Retire(w *Worker) {
	// create a new worker
	newID := nextWorkerID
	nextWorkerID = nextWorkerID + 1 // Giant RACE CONDITION -- this is a demo.
	newW := &Worker{
		ID:         newID,
		StartTime:  time.Now(),
		workerPool: w.workerPool,
	}
	log.Printf("INFO:  new worker(%d) -> in-service", newW.ID)
	w.workerPool <- newW // add new worker to the pool
	log.Printf("INFO:  worker(%d) retired", w.ID)
	//
	// Let old worker, 'w', quietly slip into the darkness and be
	// reaped by garbage collection.
}
