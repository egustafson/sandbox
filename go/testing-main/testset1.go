package main

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

// DemoTestSet1 is the simplest example of a test set (suite).
type DemoTestSet1 struct {
	suite.Suite
}

func RunDemoTestSet1(t *testing.T) {
	suite.Run(t, new(DemoTestSet1))
}

func (suite *DemoTestSet1) TestOne_Passing() {
	value := true
	suite.True(value)
}

func (suite *DemoTestSet1) TestTwo_Failing() {
	value := false
	suite.True(value) // this is intended to fail
}
