package demofwk

import (
	"context"
	"log/slog"
	"reflect"
	"runtime"
	"sync"
)

type Registry struct {
	runnables []Runnable
	mu        sync.Mutex
}

var registry *Registry = newRegistry()

func Register(runnable Runnable) {
	registry.mu.Lock()
	defer registry.mu.Unlock()

	registry.runnables = append(registry.runnables, runnable)
}

func newRegistry() *Registry {
	return &Registry{
		runnables: make([]Runnable, 0),
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
	for _, r := range registry.runnables {
		name := runtime.FuncForPC(reflect.ValueOf(r).Pointer()).Name()
		rootT.Run(name, r)
	}
	slog.Info("Run.demofwk - finished")
}
