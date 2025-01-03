package demofwk

import (
	"log/slog"
	"reflect"
	"runtime"
	"sync"
)

type TestFn func(t *T)

type TestRec struct {
	Name string
	Fn   TestFn
}

type Registry struct {
	tests []*TestRec
	lock  sync.Mutex
}

var registry *Registry = newRegistry()

func Register(testfn TestFn) {
	rr := &TestRec{
		Name: runtime.FuncForPC(reflect.ValueOf(testfn).Pointer()).Name(),
		Fn:   testfn,
	}
	registry.lock.Lock()
	defer registry.lock.Unlock()

	registry.tests = append(registry.tests, rr)
}

func newRegistry() *Registry {
	return &Registry{
		tests: make([]*TestRec, 0),
	}
}

func newRootT() *T {
	return newT(nil, "")
}

func Run() {
	slog.Info("Run.demofwk - start")
	rootT := newRootT()
	for _, tr := range registry.tests {
		name := tr.Name
		rootT.Run(name, tr.Fn)
	}
	slog.Info("Run.demofwk - finished")
}
