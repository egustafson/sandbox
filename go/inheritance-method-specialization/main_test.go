package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChild1_Beahvior(t *testing.T) {
	var c Iface = NewChild1("child-1", 1)
	assert.Equal(t, "type-child-1", c.Behavior())
}

func TestChild2_Behavior(t *testing.T) {
	var c Iface = NewChild2("child-2", 2)
	assert.Equal(t, "type-child-2", c.Behavior())
}
