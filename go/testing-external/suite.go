package main

import (
	"fmt"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/suite"
)

type DemoTestSuite struct {
	suite.Suite
}

func TestDemoTestSuite(t *testing.T) {
	suite.Run(t, new(DemoTestSuite))
}

func (s *DemoTestSuite) TestExamplePassing() {
	slog.Info("running suite test", slog.String("name", s.T().Name()))
}

func (s *DemoTestSuite) TestExampleFailing() {
	slog.Info("running suite test", slog.String("name", s.T().Name()))
	s.Fail("force failure")
}

func (s *DemoTestSuite) TestExampleNested() {
	slog.Info("running suite test", slog.String("name", s.T().Name()))
	for i := 0; i < 3; i++ {
		s.T().Run(fmt.Sprintf("nested-%d", i), func(t *testing.T) {
			s.Greater(i, 0, "expected failure") // the 0'th one should fail
		})
	}
}
