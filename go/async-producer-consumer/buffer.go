package main

import (
	"sync"
)

type Buffer interface {
	Push(v interface{})
	Pull() (v interface{}, more bool)
	Size() int
	Close()
}

type element struct {
	v    interface{} // value
	next *element
	prev *element
}

type buffer struct {
	size   int
	head   *element
	tail   *element
	cond   *sync.Cond
	mutex  *sync.Mutex
	closed bool
}

func NewBuffer() Buffer {
	m := new(sync.Mutex)
	return &buffer{
		cond:   sync.NewCond(m),
		mutex:  m,
		closed: false,
	}

}

func (b *buffer) Push(v interface{}) {
	if b.closed {
		return
	}
	newElem := &element{v: v}

	b.mutex.Lock()
	defer b.mutex.Unlock()
	defer b.cond.Signal()

	if b.head != nil {
		newElem.next = b.head
		b.head.prev = newElem
		b.head = newElem
	} else { // buffer is empty - initialize
		b.head = newElem
		b.tail = newElem
	}
	b.size++
}

func (b *buffer) Pull() (v interface{}, more bool) {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	for b.size <= 0 && !b.closed {
		b.cond.Wait()
	}
	if b.size > 0 {
		v = b.tail.v // return value
		b.tail = b.tail.prev
		if b.tail != nil {
			b.tail.next = nil
		} else {
			b.head = nil
		}
		b.size--
	}
	more = !b.closed || b.size > 0
	return v, more
}

func (b *buffer) Size() int {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	return b.size
}

func (b *buffer) Close() {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	b.closed = true
	b.cond.Signal()
}
