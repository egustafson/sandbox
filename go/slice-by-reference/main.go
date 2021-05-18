package main

import (
	"fmt"
)

func crunch(vec []string) []string {
	// vec is passed by value.  Modifying inside crunch() does not
	// modify the variable passed as vec.
	return vec[1:]
}

func main() {
	vec := []string{"a1", "a2", "a3", "a4"}

	fmt.Printf("vec: %v\n", vec)
	vec = crunch(vec)
	fmt.Printf("vec: %v\n", vec)

	fmt.Println("done.")
}
