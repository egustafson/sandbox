package main

import (
	"fmt"
	"io"
	"log/slog"
	"reflect"
	"testing"
	"time"
)

func main() {
	tests := []testing.InternalTest{
		{Name: "TestPassing", F: TestPrototypePass},
		{Name: "TestFailing", F: TestPrototypeFail},
		{Name: "TestNested", F: TestPrototypeNested},
		{Name: "TestSuite", F: TestDemoTestSuite},
	}

	//testing.Init()

	m := testing.MainStart(testDeps{}, tests, nil, nil, nil)
	success := m.Run()
	//	success := testing.RunTests(matcherPass, tests)

	fmt.Printf("success:  %v\n", success)
}

// func matcherPass(pat, str string) (bool, error) {
// 	return true, nil
// }

func TestPrototypePass(t *testing.T) {

	slog.Info("running prototype pass")
	// silently pass
}

func TestPrototypeFail(t *testing.T) {
	slog.Info("running prototype fail")
	t.Fail()
}

func TestPrototypeNested(t *testing.T) {
	slog.Info("running prototype nested")
	tests := []struct {
		Name string
		V    int
	}{
		{Name: "test 1", V: 10},
		{Name: "test 2", V: 200},
		{Name: "fail test 3", V: -1},
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			if tt.V < 0 {
				t.Errorf("test failed: v = %d", tt.V)
			}
		})
	}
}

// ------------------------------------------------------------------

type corpusEntry = struct {
	Parent     string
	Path       string
	Data       []byte
	Values     []any
	Generation int
	IsSeed     bool
}

type testDeps struct{}

func (f testDeps) MatchString(pat, str string) (bool, error)   { return true, nil }
func (f testDeps) StartCPUProfile(w io.Writer) error           { return nil }
func (f testDeps) StopCPUProfile()                             {}
func (f testDeps) WriteProfileTo(string, io.Writer, int) error { return nil }
func (f testDeps) ImportPath() string                          { return "" }
func (f testDeps) StartTestLog(io.Writer)                      {}
func (f testDeps) StopTestLog() error                          { return nil }
func (f testDeps) SetPanicOnExit0(bool)                        {}
func (f testDeps) CoordinateFuzzing(time.Duration, int64, time.Duration, int64, int, []corpusEntry, []reflect.Type, string, string) error {
	return nil
}
func (f testDeps) RunFuzzWorker(func(corpusEntry) error) error { return nil }
func (f testDeps) ReadCorpus(string, []reflect.Type) ([]corpusEntry, error) {
	return nil, nil
}
func (f testDeps) CheckCorpus([]any, []reflect.Type) error { return nil }
func (f testDeps) ResetCoverage()                          {}
func (f testDeps) SnapshotCoverage()                       {}

func (f testDeps) InitRuntimeCoverage() (mode string, tearDown func(string, string) (string, error), snapcov func() float64) {
	return
}
