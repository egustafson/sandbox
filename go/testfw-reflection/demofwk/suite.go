package demofwk

import (
	"log/slog"
	"reflect"
	"regexp"
	"sync"
)

var (
	defaultMethodPrefix = "^Test"
	methodRE            = regexp.MustCompile(defaultMethodPrefix)
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
	slog.Info("starting suite", slog.String("name", t.Name))
	s.SetT(t)
	s.SetS(s)
	//
	runnables := []Runnable{}
	methodFinder := reflect.TypeOf(s)
	suiteName := methodFinder.Elem().Name()
	for i := 0; i < methodFinder.NumMethod(); i++ {
		method := methodFinder.Method(i)
		if methodRE.MatchString(method.Name) {
			slog.Info("method", slog.String("name", method.Name), slog.String("suite", suiteName))
			r := func(t *T) {
				defer func() { // teardown hook
					// add hooks to run teardown hook if it exists
					slog.Info("- teardown hook", slog.String("method", method.Name))
				}()

				// setup hook
				slog.Info(" - setup hook", slog.String("method", method.Name))
				method.Func.Call([]reflect.Value{reflect.ValueOf(s)})
			}
			runnables = append(runnables, r)
		}
	}

	for _, r := range runnables {
		t.Run("child-need-name", r)
	}
	//
	// TODO
	//
	slog.Info("completed suite", slog.String("name", t.Name))
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
