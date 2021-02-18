// +build !integration

// this is NOT an Integration test, it is run whenever `go test` is run.

package main_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnitTestingExample(t *testing.T) {
	assert.True(t, true)
}
