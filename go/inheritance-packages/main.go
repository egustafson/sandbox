package main

import (
	"fmt"

	"github.com/egustafson/sandbox/go/inheritance-packages/circle"
	"github.com/egustafson/sandbox/go/inheritance-packages/square"
)

func main() {

	c := circle.NewCircle(5)
	s := square.NewSquare(10)

	fmt.Printf("Circle width %d\n", c.GetWidth())
	fmt.Printf("Square width %d\n", s.GetWidth())

	s.GetShapeRec().Name = "square2"

	fmt.Printf("Square name:  %s\n", s.GetShapeRec().Name)

	fmt.Println("done.")
}
