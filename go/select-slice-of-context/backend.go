package main

import (
	"context"
	"sync"
)

type Backend interface {
	Watch(ctx context.Context, watchID int) WatchCh
	Message(m string)
	Fail(watchID int)
}

type bwatch struct {
	id     int
	ch     chan<- string
	ctx    context.Context
	cancel context.CancelFunc
}

type bimpl struct {
	watches map[int]*bwatch
	lock    sync.Mutex
}

// enforce static check: struct bimpl implements Backend
var _ Backend = (*bimpl)(nil)

func NewBackend() Backend {
	b := &bimpl{
		watches: make(map[int]*bwatch),
	}
	return b
}

func (b *bimpl) Watch(ctx context.Context, watchID int) WatchCh {
	ch := make(chan string, 1)
	w := &bwatch{
		id: watchID,
		ch: ch,
	}
	w.ctx, w.cancel = context.WithCancel(ctx)

	b.lock.Lock()
	defer b.lock.Unlock()

	b.watches[watchID] = w
	// _cleanup goroutine_: Normally this cleanup with be based on the
	// underlying I/O system, so we'll simulate it here.
	go func() {
		<-w.ctx.Done() // block for watcher context to be canceled
		close(w.ch)
		b.lock.Lock()
		defer b.lock.Unlock()
		delete(b.watches, watchID)
	}()
	return ch
}

func (b *bimpl) Message(m string) {
	b.lock.Lock()
	defer b.lock.Unlock()

	for _, w := range b.watches {
		if w.ctx.Err() != nil {
			continue
		}
		select {
		case <-w.ctx.Done():
			// guard against write on close
		case w.ch <- m:
			// just send
		}
	}
}

func (b *bimpl) Fail(watchID int) {
	b.lock.Lock()
	defer b.lock.Unlock()

	if w, ok := b.watches[watchID]; ok {
		w.cancel()
		// closing the channel and removing it from watchers is
		// handled in the cleanup goroutine
	}
}
