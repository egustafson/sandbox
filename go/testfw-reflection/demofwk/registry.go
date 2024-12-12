package demofwk

import "sync"

type Registry struct {
	testSuites []any
	mu         sync.Mutex
}

var registry *Registry = newRegistry()

func Register(suite any) {
	registry.mu.Lock()
	defer registry.mu.Unlock()

	registry.testSuites = append(registry.testSuites, suite)
}

func newRegistry() *Registry {
	return &Registry{
		testSuites: make([]any, 0),
	}
}
