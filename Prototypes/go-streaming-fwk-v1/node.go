package gostr

import (
	"errors"
	"fmt"
)

type Parameters interface{}

type Value interface{}

type Node interface {
	Id() string
	Depends() []string
	Attach(ch chan<- Value, param Parameters) error
	Start() // blocking, and recursively calls dependent Node.Start()
	Stop()
}

type Registry interface {
	Register(n Node) error
	Get(id string) (Node, error)
}

type SimpleRegistry struct {
	nodes map[string]Node
}

func (r *SimpleRegistry) Register(n Node) error {
	if _, ok := r.nodes[n.Id()]; ok {
		return errors.New(fmt.Sprintf("[%s] node previously registered, no duplicates", n.Id()))
	}
	r.nodes[n.Id()] = n
	return nil
}

func (r *SimpleRegistry) Get(id string) (Node, error) {
	n, ok := r.nodes[id]
	if !ok {
		return nil, errors.New(fmt.Sprintf("no node named [%s] registered.", id))
	}
	return n, nil
}

// //////////   package *default* registry   //////////

var defaultRegistry = SimpleRegistry{
	nodes: make(map[string]Node),
}

func Register(n Node) error {
	return defaultRegistry.Register(n)
}

func Get(id string) (Node, error) {
	return defaultRegistry.Get(id)
}
