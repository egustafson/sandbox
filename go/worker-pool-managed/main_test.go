package main

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

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
