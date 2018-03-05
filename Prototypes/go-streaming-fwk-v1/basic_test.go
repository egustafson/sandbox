package gostr

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

type TestValue struct {
	val int
}

// Test Skeleton (type: simpleNode)

type simpleNode struct {
	id       string
	shutdown bool
}

func (n *simpleNode) Id() string {
	return n.id
}

func (n *simpleNode) Stop() {
	n.shutdown = true
}

// Test Producer

type producer struct {
	simpleNode
	inc    int
	upTo   int
	output chan<- TestValue
}

func MakeProducer(step int, upTo int) producer {
	return producer{
		simpleNode: simpleNode{
			id:       "prod1",
			shutdown: false,
		},
		inc:  step,
		upTo: upTo,
	}
}

func (p *producer) Depends() []string {
	return []string{} // empty list
}

func (p *producer) Attach(ch chan<- TestValue, param interface{}) error {
	if p.output != nil {
		return errors.New("producer already has its only attachment")
	}
	p.output = ch
	return nil
}

func (p *producer) Start() {
	if p.output == nil {
		return
	}
	ii := 0
	for { // loop forever (until shutdown)
		if p.shutdown || ii > p.upTo {
			if p.shutdown {
				fmt.Println("received shutdown.")
			}
			close(p.output)
			return
		}
		p.output <- TestValue{val: ii}
		ii += p.inc
	}
}

// Test Consumer

type consumer struct {
	simpleNode
	in chan TestValue
}

func MakeConsumer(p *producer) consumer {
	ch := make(chan TestValue)
	c := consumer{
		simpleNode: simpleNode{"consumer1", false},
		in:         ch,
	}
	p.Attach(ch, nil)
	return c
}

func (c *consumer) Depends() []string {
	return []string{} // this is a lie, but its not being tested
}

func (c *consumer) Attach(ch chan<- Value, param Parameters) error {
	// there's no attaching to a consumer, it doesn't output anything.
	return errors.New("can not attach to this consumer")
}

func (c *consumer) Start() {
	for { // loop until input chan is closed
		v, ok := <-c.in
		if !ok {
			fmt.Println("consumer stopping.")
			return // chan closed, we're done
		}
		fmt.Printf("%d\n", v.val)
	}
}

// The test case.

func TestOne(t *testing.T) {
	p := MakeProducer(1, 100000) // set an upper bound so we don't run-away.
	c := MakeConsumer(&p)
	go c.Start()
	go p.Start()
	time.Sleep(20 * time.Microsecond) // should complete before upper bound.
	p.Stop()
}
