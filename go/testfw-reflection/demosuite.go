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

func DemoTestSuiteRunner(t *demofwk.T) {
	demofwk.RunSuite(t, new(DemoTestSuite))
}

func (s *DemoTestSuite) TestDemo1() {
	slog.Info("DemoTestSuite running", slog.String("name", s.T().Name))
}
