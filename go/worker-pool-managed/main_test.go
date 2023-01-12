package main

import (
	"context"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/Masterminds/log-go"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	log.Current = &log.StdLogger{
		Level: log.DebugLevel,
	}
	os.Exit(m.Run())
}

func TestNewPoolSimpleWorker(t *testing.T) {

	p := NewLazyInitWorkerPool(SimpleWorkerFactory, "pool-1", 2, 1)
	assert.NotNil(t, p)
	p.Close()
}

func TestSimpleWorkerPool(t *testing.T) {

	p := NewLazyInitWorkerPool(SimpleWorkerFactory, "pool-simple-worker-test", 2, 1)

	ctx, cancel := context.WithTimeout(context.Background(), 600*time.Second)
	defer cancel()

	r, err := p.DoWork(ctx, &WorkDescription{1})
	assert.Nil(t, err)
	assert.True(t, r.Success)
}

func TestScaledSimpleWorkerPool(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	p := NewLazyInitWorkerPool(SimpleWorkerFactory, "scale-test-simple-worker", 5, 2)

	var wg sync.WaitGroup
	groupSize := 100

	for ii := 0; ii < groupSize; ii++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			r, err := p.DoWork(ctx, &WorkDescription{idx})
			if assert.Nil(t, err) {
				assert.True(t, r.Success)
			}
		}(ii)
	}
	wg.Wait()
	p.Close()
}
