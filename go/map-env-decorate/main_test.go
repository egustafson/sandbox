package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func StubMapFunc(in string) string {
	if strings.HasPrefix(in, "$") {
		if replacement, ok := envvars[in[1:]]; ok {
			return replacement
		}
	}
	return in
}

// TestStruct must be formally declared in order for the fields, (Name, Val) to
// be settable (reflect.Value.CanSet()).  This was discovered empirically; I
// believe inline struct definitions are qualified as unexported stucts which
// has a cascade effect on the fields within the unexported struct.
//
// In short, the rules on "exported" elements are not always obvious.  Inlining
// this stuct in the test doesn't work.
type TestStruct struct {
	Name string
	Val  int
}

func TestDecorateStruct(t *testing.T) {

	// test decorating a ptr

	testval := &TestStruct{Name: "$USERNAME", Val: 21}

	if err := Decorate(StubMapFunc, testval); err != nil {
		t.Errorf("decorate returned error: %s", err)
	}
	assert.Equal(t, "bogus-user", testval.Name) // decoration should translate
	assert.Equal(t, 21, testval.Val)

	// test decorating a value struct

	t2val := TestStruct{Name: "$USERNAME", Val: 21} // decorate a value

	if err := Decorate(StubMapFunc, &t2val); err != nil { // MUST pass in as a ref (i.e. ptr)
		t.Errorf("decorate returned error: %s", err)
	}
	assert.Equal(t, "bogus-user", t2val.Name)
	assert.Equal(t, 21, t2val.Val)
}
