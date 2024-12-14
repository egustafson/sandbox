package demofwk

import (
	"log/slog"
	"sync"
)

type Suite struct {
	t *T
	s SuiteHolder

	mu sync.RWMutex
}

type SuiteHolder interface {
	T() *T
	SetT(t *T)
	SetS(suite SuiteHolder)
}

var _ SuiteHolder = (*Suite)(nil) // type check that a *Suite isA SuiteHolder

func RunSuite(t *T, s SuiteHolder) {
	s.SetT(t)
	s.SetS(s)
	//
	// TODO
	//
	slog.Info("running suite", slog.String("name", t.Name))
}

func (s *Suite) T() *T {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.t
}

func (s *Suite) SetT(t *T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.t = t
	//
	// TODO: additional initialization of a *Suite object
	//
}

func (suite *Suite) SetS(s SuiteHolder) {
	suite.s = s
}
