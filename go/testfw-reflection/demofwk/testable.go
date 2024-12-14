package demofwk

import (
	"context"
	"fmt"
)

type Runnable func(t *T)

type T struct {
	Name     string
	ctx      context.Context
	cancel   context.CancelFunc
	parent   *T
	children []*T
	failed   bool
}

func (t *T) Run(name string, fn Runnable) {
	child := &T{
		Name:   fmt.Sprintf("%s:%s", t.Name, name),
		parent: t,
	}
	child.ctx, child.cancel = context.WithCancel(t.ctx)

	child.doRun(fn)
}

func (t *T) doRun(fn Runnable) {
	running := make(chan struct{})
	go func() {
		fn(t)
		close(running)
	}()
	<-running
}

func newT(parent *T, name string) *T {
	t := &T{
		Name:   name,
		parent: parent,
	}
	t.ctx, t.cancel = context.WithCancel(parent.ctx)
	return t
}
