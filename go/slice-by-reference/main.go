package main

import (
	"fmt"
)

func crunch(vec []string) []string {
	// vec is passed by value.  Modifying inside crunch() does not
	// modify the variable passed as vec.
	return vec[1:]
}

func addto(vec []string) []string {
	// vec is passed by value.  Appending to it does not modify the
	// variable passed in.
	vec = append(vec, "a5")
	return vec
}

func main() {
	ovec := []string{"a1", "a2", "a3", "a4"}

	fmt.Printf("orig vec: %v\n", ovec)

	fmt.Println("crunch(ovec)")
	nvec := crunch(ovec)
	fmt.Printf("orig vec: %v\n", ovec)
	fmt.Printf("new  vec: %v\n", nvec)

	fmt.Println("addto(ovec)")
	nvec = addto(ovec)
	fmt.Printf("orig vec: %v\n", ovec)
	fmt.Printf("new  vec: %v\n", nvec)

	fmt.Println("done.")
}
