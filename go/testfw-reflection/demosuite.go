package main

import (
	"log/slog"

	"github.com/egustafson/go/testfw-reflection/demofwk"
)

type DemoTestSuite struct {
	demofwk.Suite

	suiteCounter     int
	suiteInitialized bool
}

func init() {
	demofwk.Register(DemoTestSuiteRunner)
}

func DemoTestSuiteRunner(t *demofwk.T) {
	demofwk.RunSuite(t, new(DemoTestSuite))
}

func (s *DemoTestSuite) Test1() {
	slog.Info("  --> inside suite DemoTestSuite", slog.String("name", s.T().Name))
}

func (s *DemoTestSuite) Test2() {
	slog.Info("  --> inside suite DemoTestSuite", slog.String("name", s.T().Name))
}

func (s *DemoTestSuite) TestPanic() {
	slog.Info("  --> inside suite DemoTestSuite", slog.String("name", s.T().Name))
	panic("cause a panic")
}

func (s *DemoTestSuite) TestNested() {
	slog.Info("  --> inside suite DemoTestSuite", slog.String("name", s.T().Name))
	s.T().Run("nested-1", func(t *demofwk.T) {
		slog.Info("    --> nested test", slog.String("name", t.Name))
	})
	s.T().Run("nested-2", func(t *demofwk.T) {
		slog.Info("    --> nested test", slog.String("name", t.Name))
	})
	slog.Info("  --> ending nested suite test", slog.String("name", s.T().Name))
}
