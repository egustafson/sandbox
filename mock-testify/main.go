package main

import (
	"fmt"
)

type ExampleService interface {
	Message(id int) string
}

type demoService struct{}

// static check: struct demoService implements ExampleService
var _ ExampleService = (*demoService)(nil)

func (d *demoService) Message(id int) string {
	return fmt.Sprintf("(%d) demo message: run tests to see actual demo\n", id)
}

func main() {
	d := new(demoService)
	fmt.Print(d.Message(1))
}
