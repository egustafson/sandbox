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
	t    *T
	s    SuiteHolder
	lock sync.RWMutex
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
	defer s.SetT(t) // ensure the suite's T record is reset and not a child.
	//
	tests := []*TestRec{}
	methodFinder := reflect.TypeOf(s)
	suiteName := methodFinder.Elem().Name()
	t.Name = suiteName // overwrite name of T -- probably needs more robust handling.
	for i := 0; i < methodFinder.NumMethod(); i++ {
		method := methodFinder.Method(i)
		if methodRE.MatchString(method.Name) {
			tr := &TestRec{
				Name: method.Name,
				Fn: func(t *T) {
					defer func() { // teardown hook
						if suiteTeardown, ok := s.(SuiteTeardownTest); ok {
							suiteTeardown.TeardownTest()
						}
					}()

					// make the suite's T == this test's T.  The suite's T is
					// reset with a deferred SetT() in RunSuite.
					s.SetT(t)

					if suiteSetup, ok := s.(SuiteSetupTest); ok {
						suiteSetup.SetupTest()
					}
					method.Func.Call([]reflect.Value{reflect.ValueOf(s)})
				},
			}
			tests = append(tests, tr)
		}
	}

	if suiteSetup, ok := s.(SuiteSetupAll); ok {
		suiteSetup.SetupAll()
	}
	for _, tr := range tests {
		t.Run(tr.Name, tr.Fn)
	}
	if suiteTeardown, ok := s.(SuiteTeardownAll); ok {
		suiteTeardown.TeardownAll()
	}
	slog.Info("completed suite", slog.String("name", t.Name))
}

func (s *Suite) T() *T {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.t
}

func (s *Suite) SetT(t *T) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.t = t
	//
	// TODO: additional initialization of a *Suite object
	//
}

func (suite *Suite) SetS(s SuiteHolder) {
	suite.s = s
}
