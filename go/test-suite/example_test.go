package testexample

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ExampleTestSuite struct {
	suite.Suite
	SuiteScopedVar int
}

// SetupSuite will be run before the Suite begins
func (s *ExampleTestSuite) SetupSuite() {
	fmt.Println("- invoking SetupSuite()")
}

// SetupTest will be run before every test in the Suite
func (s *ExampleTestSuite) SetupTest() {
	fmt.Println("-- invoking SetupTest")
	s.SuiteScopedVar = 21
}

// TestExampleOne is one of two test cases
func (s *ExampleTestSuite) TestExampleOne() {
	s.Equal(21, s.SuiteScopedVar)
	fmt.Println("--- invoking TestExampleOne")
}

func (s *ExampleTestSuite) TestExampleTwo() {
	s.Equal(21, s.SuiteScopedVar)
	fmt.Println("--- invoking TestExampleTwo")
}

// TearDownTest will be run after every test in the Suite
func (s *ExampleTestSuite) TearDownTest() {
	fmt.Println("-- invoking TearDownTest")
}

// TearDownSuite will be run after the entire Suite completes
func (s *ExampleTestSuite) TearDownSuite() {
	fmt.Println("- invoking TearDownSuite")
}

// TestExampleTestSuite invokes the test Suite and is how the golang testing
// framework hooks into the Suite.
//
// Run this function in the debugger or `go test -v` to observe the print output
// from each func.
func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(ExampleTestSuite))
}
