package main

import (
	"github.com/Masterminds/log-go"

	"github.com/egustafson/sandbox/go/mix-in-pvt-impl/mix"
)

type Component interface {
	mix.IDer
	GetName() string
}

type compImpl struct {
	mix.IDerMixin
	Name string
}

// static check: compImpl implements Component
var _ Component = new(compImpl)

func NewComponent(n string) Component {
	c := &compImpl{
		Name: n,
	}
	c.InitID(123)
	return c
}

func (c *compImpl) GetName() string { return c.Name }

func main() {
	log.Info("done.")
}
