package main

import (
	"context"
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

	// Close is the means of closing, and orderly destructing the
	// Middleware component.
	Close()
}

// mimpl is an implementation of Middleware
type mimpl struct {
	b      Backend            // Backend
	ctx    context.Context    // (internal) context of the Middleware component
	cancel context.CancelFunc // (internal) cancelation of ctx
}

// enforce static check: struct mimpl implements Middleware
var _ Middleware = (*mimpl)(nil)

func NewMiddleware() (Middleware, Backend) {
	b := NewBackend()
	m := &mimpl{
		b: b,
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
	go func() { // spin a 'clean-up' goroutine
		select {
		case <-m.ctx.Done():
			cancel() // cancel this watcher -> backend
		case <-ctx.Done():
		}
	}()
	return m.b.Watch(ctx, watchID)
}

func (m *mimpl) Close() {
	m.cancel()
}
