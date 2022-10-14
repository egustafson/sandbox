package main

// Test (validate) Backend interface + impl

import (
	"context"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tmsg = "test-message"

func TestBackend(t *testing.T) {
	b := NewBackend()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ch1 := b.Watch(ctx, 1)

	b.Message(tmsg)
	result := <-ch1
	assert.Equal(t, result, tmsg)
}

func TestCanceledWatch(t *testing.T) {
	b := NewBackend()

	ctx, cancel := context.WithCancel(context.Background())
	// will manually invoke cancel - no defer
	ch1 := b.Watch(ctx, 1)

	cancel()          // <- should close the backend, but will not be caught until a message is sent.
	runtime.Gosched() // force context switch -> clean-up on context.Cancel()

	b.Message(tmsg)
	_, ok := <-ch1
	assert.False(t, ok) // indicating the channel was closed (by cancel())
}

func TestFailBackend(t *testing.T) {
	b := NewBackend()

	ctx := context.Background()
	ch1 := b.Watch(ctx, 1)

	b.Fail(1)
	runtime.Gosched() // force context switch -> clean-up on context.Cancel()
	b.Message(tmsg)
	_, ok := <-ch1
	assert.False(t, ok)
}
