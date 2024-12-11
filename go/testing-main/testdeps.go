package main

import (
	"io"
	"reflect"
	"time"
)

type corpusEntry = struct {
	Parent     string
	Path       string
	Data       []byte
	Values     []any
	Generation int
	IsSeed     bool
}

type mockTestDeps struct{}

func (td mockTestDeps) ImportPath() string                                       { return "" }
func (td mockTestDeps) MatchString(pat, str string) (bool, error)                { return true, nil }
func (td mockTestDeps) SetPanicOnExit0(bool)                                     {}
func (td mockTestDeps) StartCPUProfile(w io.Writer) error                        { return nil }
func (td mockTestDeps) StopCPUProfile()                                          {}
func (td mockTestDeps) StartTestLog(io.Writer)                                   {}
func (td mockTestDeps) StopTestLog() error                                       { return nil }
func (td mockTestDeps) WriteProfileTo(string, io.Writer, int) error              { return nil }
func (td mockTestDeps) RunFuzzWorker(func(corpusEntry) error) error              { return nil }
func (td mockTestDeps) ReadCorpus(string, []reflect.Type) ([]corpusEntry, error) { return nil, nil }
func (td mockTestDeps) CheckCorpus([]any, []reflect.Type) error                  { return nil }
func (td mockTestDeps) ResetCoverage()                                           {}
func (td mockTestDeps) SnapshotCoverage()                                        {}
func (td mockTestDeps) InitRuntimeCoverage() (string, func(string, string) (string, error), func() float64) {
	return "", nil, nil
}
func (td mockTestDeps) CoordinateFuzzing(
	time.Duration, int64, time.Duration, int64, int, []corpusEntry,
	[]reflect.Type, string, string) error {
	return nil
}
