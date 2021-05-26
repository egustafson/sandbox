package main

import "fmt"

type Child interface {
	Op() string
	Close() string
}

type child struct{}

func (c *child) Op() string {
	return "message from child.Op()"
}

func (c *child) Close() string {
	return "child.Close()'ing"
}

type Parent struct {
	Child
	data int
}

func (p *Parent) Close() string {
	str := p.Child.Close() + "\n" + "Parent.Close()'ing"
	return str
}

func main() {
	p := &Parent{Child: &child{}, data: 1}

	fmt.Println(p.Op())
	fmt.Println(p.Close())

	fmt.Println("done.")
}
