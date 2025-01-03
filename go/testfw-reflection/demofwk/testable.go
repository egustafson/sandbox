package demofwk

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type T struct {
	Name      string
	lock      sync.Mutex
	ctx       context.Context
	cancel    context.CancelFunc
	parent    *T
	children  []*T
	failed    bool
	panicked  bool
	done      bool
	start     time.Time
	duration  time.Duration
	completed chan struct{}
}

func newT(parent *T, name string) *T {
	t := &T{
		Name:      name,
		parent:    parent,
		children:  make([]*T, 0),
		completed: make(chan struct{}),
	}
	if parent != nil {
		t.ctx, t.cancel = context.WithCancel(parent.ctx)
	} else {
		t.ctx, t.cancel = context.WithCancel(context.Background())
	}
	return t
}

func (t *T) Run(name string, fn TestFn) {
	if len(t.Name) > 0 {
		name = fmt.Sprintf("%s:%s", t.Name, name)
	}
	child := newT(t, name)
	child.doRun(fn)
}

func (t *T) doRun(fn TestFn) {
	go func() {
		// panic handler
		defer func() {
			t.lock.Lock()
			defer t.lock.Unlock()

			t.duration = time.Since(t.start)

			// handle any panic()'s
			if err := recover(); err != nil {
				t.failed = true
				t.panicked = true
			}

			t.done = true
			close(t.completed)
		}()

		t.start = time.Now()
		fn(t)
	}()
	<-t.completed
}
