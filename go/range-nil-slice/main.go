package main

import "fmt"

func main() {
	var x []string
	fmt.Printf("x = %v of type %T\n", x, x)
	for y := range x {
		fmt.Printf("element: %v\n", y) // this should never print
	}
}
