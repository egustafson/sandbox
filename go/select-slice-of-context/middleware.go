package main

import (
	"context"
	"sync"
)

// WatchCh is the prototypical channel of watch events.  In this case,
// they're just strings.
type WatchCh <-chan string

// Middleware is the prototypical middle layer which provides a
// Watch() func where the implementatoin of Middleware want's to
// interceed on shutdown.
type Middleware interface {

	// Watch is the func that provides a watchable channel which can
	// be closed when the Middleware is Closed()
	Watch(ctx context.Context, watchID int) WatchCh

	// NumWatchers returns the current number of watchers passing
	// through the Middleware.  (i.e. non-canceled contexts)
	NumWatchers() int

	// Close is the means of closing, and orderly destructing the
	// Middleware component.
	Close()
}

type mwatch struct {
	id     int
	ctx    context.Context
	cancel context.CancelFunc
}

// mimpl is an implementation of Middleware
type mimpl struct {
	b        Backend            // Backend
	ctx      context.Context    // (internal) context of the Middleware component
	cancel   context.CancelFunc // (internal) cancelation of ctx
	watchers map[int]*mwatch
	lock     sync.Mutex
}

// enforce static check: struct mimpl implements Middleware
var _ Middleware = (*mimpl)(nil)

func NewMiddleware() (Middleware, Backend) {
	b := NewBackend()
	m := &mimpl{
		b:        b,
		watchers: make(map[int]*mwatch),
	}
	m.ctx, m.cancel = context.WithCancel(context.Background())
	return m, b
}

func (m *mimpl) Watch(ctx context.Context, watchID int) WatchCh {
	if m.ctx.Err() != nil {
		return nil // the middleware is closed
	}
	//
	// There's a race with Close() -> ctx.Cancel() in here.
	//   (this is an example, we'll live with it)
	//
	ctx, cancel := context.WithCancel(ctx)
	mw := &mwatch{
		id:     watchID,
		ctx:    ctx,
		cancel: cancel,
	}
	func() {
		m.lock.Lock()
		defer m.lock.Unlock()

		m.watchers[watchID] = mw
		go func() { // spin a 'clean-up' goroutine
			select {
			case <-m.ctx.Done():
				mw.cancel() // cancel this watcher -> backend
			case <-mw.ctx.Done():
			}
			m.lock.Lock() // we're in a new goroutine here
			defer m.lock.Unlock()
			delete(m.watchers, mw.id)
		}()
	}() // release lock here
	return m.b.Watch(ctx, watchID)
}

func (m *mimpl) NumWatchers() int {
	m.lock.Lock()
	defer m.lock.Unlock()

	return len(m.watchers)
}

func (m *mimpl) Close() {
	m.cancel()

	for _, mw := range m.watchers {
		mw.cancel()
		// monitoring goroutine will remove mw from m.watchers
	}
}
