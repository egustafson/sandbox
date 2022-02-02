package main

import (
	"context"
	"log"
	"sync"
	"time"
)

// Calculate is the function abstraction of using the worker pool
func Calculate(in int, d time.Duration) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	result := make(chan JobResp)
	j := &JobReq{
		InputData:  in,
		Delay:      d,
		Ctx:        ctx,
		ResultChan: result,
	}
	JobQueue <- j
	select {
	case r := <-result:
		return r.Result, nil
	case <-ctx.Done():
		return 0, ctx.Err()
	}
}

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	go Dispatcher(ctx, WorkerPool, JobQueue)

	time.Sleep(250 * time.Millisecond) // cause some timeout(s) waiting for work

	log.Print("--- submit some work serially")
	for ii := 1; ii < 10; ii++ {
		if ii == 2 {
			time.Sleep(250 * time.Millisecond) // cause timeout(s) waiting for work
		}
		if ii > 5 {
			time.Sleep(90 * time.Millisecond)
		}
		r, err := Calculate(ii, time.Millisecond)
		if err != nil {
			log.Printf("Calculate(%d) <- err: %v", ii, err)
		} else {
			log.Printf("Calculate(%d) = %d", ii, r)
		}
	}
	log.Print("--- now submit work in parallel")
	var wg sync.WaitGroup
	for ii := 1; ii < 13; ii++ {
		wg.Add(1)
		go func(idx int) {
			r, err := Calculate(idx, 150*time.Millisecond)
			if err != nil {
				log.Printf("Calculate(%d) <- err: %v", idx, err)
			} else {
				log.Printf("Calculate(%d) = %d", idx, r)
			}
			wg.Done()
		}(ii)
	}

	wg.Wait()
	log.Print("shutting down")
	cancel()
	time.Sleep(250 * time.Millisecond)
	log.Print("done.")
}
