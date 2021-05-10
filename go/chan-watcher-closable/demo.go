package main

import (
	"fmt"
	"math/rand"
	_ "time"
)

// Give the consumer access to:
//  1. the channel (read only)
//  2. a closer()
type IntWatcher interface {
	Chan() <-chan int
	Close()
}

// --  Watcher Implementation  -----------------------------

// Hold internally
//  1. the channel (read/write)
//  2. a channel to indicate being complete by closing
type intWatch struct {
	ch   chan int
	done chan struct{}
}

func newIntWatch() *intWatch {
	return &intWatch{
		ch:   make(chan int, 2),
		done: make(chan struct{}),
	}
}

func (w *intWatch) Chan() <-chan int { return w.ch }

func (w *intWatch) Close() { close(w.done) }

// --  Producer that uses intWatcher  ----------------------

// An object with a bunch of watchers
type producable struct {
	count    int // how many ints to send
	watchers map[int]*intWatch
}

func NewProducer(count int) *producable {
	return &producable{
		count:    count,
		watchers: make(map[int]*intWatch),
	}
}

func (p *producable) Start() {
	fmt.Println("start producer.")
	for ii := 0; ii < p.count; ii++ {
		//time.Sleep(100 * time.Millisecond)
		//fmt.Printf("iteration: %d\n", ii)
		n := rand.Intn(100)
		toClose := make([]int, 0)
		for id, w := range p.watchers { // not thread safe
			select {
			case w.ch <- n:
			case <-w.done:
				close(w.ch)
				toClose = append(toClose, id)
			}
		}
		// remove the closed watchers
		for _, id := range toClose { // not thread safe
			delete(p.watchers, id)
			//fmt.Printf("count(%d), removed: %d\n", ii, id)
		}
	}
}

func (p *producable) Watch() IntWatcher {
	id := len(p.watchers)
	w := newIntWatch()
	p.watchers[id] = w
	return w
}

func (p *producable) NumWatchers() int {
	return len(p.watchers)
}

// --  main() func - not used  -----------------------------

func main() {
	// do nothing
}
