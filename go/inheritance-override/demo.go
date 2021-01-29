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

	fmt.Println("----")
	b.Operation()
	fmt.Println("")
	s.Operation()
	fmt.Println("")

	fmt.Println("done.")
}
