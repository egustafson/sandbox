package main

// Race Conditions: note that there are numerous points in the test
// below where either runtime.Gosched() is invoked, and/or
// <-time.After(time.Millisecond) is invoked.
//
// In these cases, the intent is to force the scheduler to context
// switch to perform actions that should occur in other goroutines.
// In NONE of these cases where the results tested _after_ the forced
// context switch is there a guarantee the event happens at any
// deadline time in the future.  The forced context switches are
// merely to speed up the test run.  Correctness is defined by the
// action after the forced context switch _eventually_ happening.

import (
	"context"
	"runtime"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type AppWatcher struct {
	id     int
	ch     WatchCh // <-chan string
	ctx    context.Context
	cancel context.CancelFunc
}

func TestMain(t *testing.T) {
	// create the Middleware => create backend
	m, b := NewMiddleware()

	w := make(map[int]AppWatcher) // watchers
	for ii := 0; ii < 5; ii++ {
		ctx, cancel := context.WithCancel(context.Background())
		wch := m.Watch(ctx, ii)
		w[ii] = AppWatcher{
			id:     ii,
			ch:     wch,
			ctx:    ctx,
			cancel: cancel,
		}
	}

	assert.Equal(t, m.NumWatchers(), 5)

	// Backend closes with nothing to be read in the queues
	//
	b.Fail(0)         // immediately fail watcher 0 from the backend
	runtime.Gosched() // yield so the canceled context can react
	select {
	case _, ok := <-w[0].ch:
		assert.False(t, ok) // indicating the channel was closed
		w[0].cancel()       // cancel just because
		delete(w, 0)        // remove the watcher, it's extinguished
	case <-time.After(time.Millisecond):
		assert.Fail(t, "w[0].ch should have been closed by b.Fail(0)")
	}

	runtime.Gosched()              // yield so the canceled context can react
	<-time.After(time.Millisecond) // yield hard to give other goroutines a chance to run

	assert.Equal(t, m.NumWatchers(), 4)

	//  Verify all (existing) watchers receive messages
	//
	b.Message("test-message")
	ReadAllWatchers(t, w)
	AllWatchersBlockOnRead(t, w)

	runtime.Gosched() // yield so the canceled context can react

	assert.Equal(t, m.NumWatchers(), 4)

	// Backend closes with messages in the channels
	//
	b.Message("test-message")
	b.Fail(1)             // from the backend: cancel/close after write
	runtime.Gosched()     // yield so the canceled context can react
	ReadAllWatchers(t, w) // all watchers should return the message, they all got it
	select {
	case _, ok := <-w[1].ch:
		assert.False(t, ok) // indicating the channel was closed
		w[1].cancel()       // clean-up
		delete(w, 1)        // remove the watcher, it's extinguished
	case <-time.After(time.Millisecond):
		assert.Fail(t, "w[1].ch should have been closed by b.Fail(1)")
	}
	AllWatchersBlockOnRead(t, w)

	runtime.Gosched() // yield so the canceled context can react

	assert.Equal(t, m.NumWatchers(), 3)

	// Frontend closes with messages in the channels
	//
	b.Message("test-message")
	w[2].cancel()         // from the frontend: cancel/close after write
	runtime.Gosched()     // yield so the canceled context can react
	ReadAllWatchers(t, w) // all watchers should return the message, they all got it
	select {
	case _, ok := <-w[2].ch:
		assert.False(t, ok) // indicating the channel was closed
		w[2].cancel()       // double-cancel, which should be ok
		delete(w, 2)        // remove the watcher, it's extinguished
	case <-time.After(time.Millisecond):
		assert.Fail(t, "w[2].ch should have been closed by w[2].Cancel()")
	}
	AllWatchersBlockOnRead(t, w)

	runtime.Gosched() // yield so the canceled context can react

	assert.Equal(t, m.NumWatchers(), 2)

	// Frontend closes with the channels empty
	w[3].cancel() // from the frontend: cancel/close before write
	// no goroutine yield, cancelation is strictly ordered before write
	b.Message("test-message")
	select {
	case _, ok := <-w[3].ch:
		assert.False(t, ok) // indicating the channel was closed
		w[3].cancel()       // double-cancel, which should be ok
		delete(w, 3)
	case <-time.After(time.Millisecond):
		assert.Fail(t, "w[3].ch should have been closed by w[3].cancel()")
	}
	ReadAllWatchers(t, w) // all (remaining) watchers should return a message

	runtime.Gosched()              // yield so the canceled context can react
	<-time.After(time.Millisecond) // yield hard to give other goroutines a chance to run

	assert.Equal(t, m.NumWatchers(), 1)
}

func ReadAllWatchers(t *testing.T, watchers map[int]AppWatcher) {
	for id, w := range watchers {
		select {
		// do not test for w.ch.Done() as the important part is the channel has data for consumption
		case <-w.ch:
			// good - this is what's supposed to happen
		case <-time.After(time.Millisecond):
			assert.Failf(t, "nothing to read, unexpected", "watcher[%d]", id)
		}
	}
}

func AllWatchersBlockOnRead(t *testing.T, watchers map[int]AppWatcher) {
	for id, w := range watchers {
		select {
		case <-w.ch:
			assert.Fail(t, "watcher had pending read", "watcher[%d]", id)
		case <-w.ctx.Done():
			assert.Fail(t, "watcher ctx canceled", "watcher[%d]", id)
		case <-time.After(time.Millisecond):
			// good - this is what's supposed to happen
		}
	}
}
