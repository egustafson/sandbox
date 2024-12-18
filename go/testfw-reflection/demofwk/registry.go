package demofwk

import (
	"context"
	"log/slog"
	"reflect"
	"runtime"
	"sync"
)

type RunnableRec struct {
	Name string
	R    Runnable
}

type Registry struct {
	runnables []*RunnableRec
	mu        sync.Mutex
}

var registry *Registry = newRegistry()

func Register(runnable Runnable) {
	rr := &RunnableRec{
		Name: runtime.FuncForPC(reflect.ValueOf(runnable).Pointer()).Name(),
		R:    runnable,
	}
	registry.mu.Lock()
	defer registry.mu.Unlock()

	registry.runnables = append(registry.runnables, rr)
}

func newRegistry() *Registry {
	return &Registry{
		runnables: make([]*RunnableRec, 0),
	}
}

func newRootT() *T {
	t := &T{
		Name:     "",
		children: make([]*T, 0),
	}
	t.ctx, t.cancel = context.WithCancel(context.Background())
	return t
}

func Run() {
	slog.Info("Run.demofwk - start")
	rootT := newRootT()
	for _, rr := range registry.runnables {
		name := rr.Name
		rootT.Run(name, rr.R)
	}
	slog.Info("Run.demofwk - finished")
}
