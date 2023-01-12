package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComponent(t *testing.T) {
	c := NewComponent("test-component")

	assert.Equal(t, c.GetID(), 123)

	c.InitID(987)
	assert.Equal(t, c.GetID(), 987)
}
