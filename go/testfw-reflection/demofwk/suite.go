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

type SuiteSetupAll interface{ SetupAll() }
type SuiteSetupTest interface{ SetupTest() }
type SuiteTeardownAll interface{ TeardownAll() }
type SuiteTeardownTest interface{ TeardownTest() }

func RunSuite(t *T, s SuiteHolder) {
	slog.Info("starting suite", slog.String("name", t.Name))
	s.SetT(t)
	s.SetS(s)
	//
	runnables := []*RunnableRec{}
	methodFinder := reflect.TypeOf(s)
	suiteName := methodFinder.Elem().Name()
	t.Name = suiteName // overwrite name of T -- probably needs more robust handling.
	for i := 0; i < methodFinder.NumMethod(); i++ {
		method := methodFinder.Method(i)
		if methodRE.MatchString(method.Name) {
			rr := &RunnableRec{
				Name: method.Name,
				R: func(t *T) {
					defer func() { // teardown hook
						if suiteTeardown, ok := s.(SuiteTeardownTest); ok {
							suiteTeardown.TeardownTest()
						}
					}()

					if suiteSetup, ok := s.(SuiteSetupTest); ok {
						suiteSetup.SetupTest()
					}
					method.Func.Call([]reflect.Value{reflect.ValueOf(s)})
				},
			}
			runnables = append(runnables, rr)
		}
	}

	if suiteSetup, ok := s.(SuiteSetupAll); ok {
		suiteSetup.SetupAll()
	}
	for _, rr := range runnables {
		t.Run(rr.Name, rr.R)
	}
	if suiteTeardown, ok := s.(SuiteTeardownAll); ok {
		suiteTeardown.TeardownAll()
	}
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
