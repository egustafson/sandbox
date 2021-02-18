// +build integration

// this IS an Integration test.  It is only run with `--tags=integration`

package main_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntegrationTestExample(t *testing.T) {
	assert.True(t, true)
}
