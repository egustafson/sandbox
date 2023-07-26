package main

import "fmt"

type Iface interface {
	Behavior() string
}

type Base struct {
	id  string
	val int
}

type Child1 struct{ Base }

// static check: *Child1 struct isA Iface
var _ Iface = (*Child1)(nil)

type Child2 struct{ Base }

// static check: *Child2 struct isA Iface
var _ Iface = (*Child2)(nil)

func NewChild1(id string, val int) Iface {
	return &Child1{Base: Base{id, val}}
}

func (*Child1) Behavior() string {
	return "type-child-1"
}

func NewChild2(id string, val int) Iface {
	//return &Child2{Base: Base{id: id, val: val}}
	return &Child2{Base: Base{id, val}}
}

func (*Child2) Behavior() string {
	return "type-child-2"
}

// ----------------------------------------------
func main() {
	fmt.Println("done - see test case for example")
}
