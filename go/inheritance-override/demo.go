package main

import (
	"fmt"
)

type Base struct {
	Name string
}

type Super struct {
	Base
}

type Super2 struct {
	Base
	Ver int
}

func (b *Base) Operation() {
	fmt.Printf("Base.Operation() on %s\n", b.Name)
}

func (s *Super) Operation() {
	fmt.Printf("Super.Operation() on %s\n", s.Name)
	fmt.Print("  ")      // indent 2 spaces
	(s.Base).Operation() // invoke base method
}

// --  Demo  --

func main() {

	b := Base{
		Name: "Base-1",
	}
	s := Super{
		Base: Base{Name: "Super-1"},
	}
	s2 := Super2{
		Base: Base{Name: "Super-2"},
		Ver:  2,
	}

	fmt.Println("----")
	b.Operation()
	fmt.Println("")
	s.Operation()
	fmt.Println("")
	fmt.Println("expect Base.Operation() to be invoked")
	s2.Operation()
	fmt.Println("")

	fmt.Println("done.")
}
