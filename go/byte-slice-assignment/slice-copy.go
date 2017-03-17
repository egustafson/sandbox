package main

import (
	"fmt"
)

func main() {
	// initialize a slice
	s1 := []byte{1, 2, 3, 4, 5, 6}
	s2 := make([]byte, len(s1))
	// "copy" s2 <- s1 (into s2)
	copy(s2, s1)
	// now zero the elements in s1
	for c, _ := range s2 {
		s1[c] = 0
	}
	// and print both
	fmt.Printf("s1: %v\n", s1)
	fmt.Printf("s2: %v\n", s2)
	//
	// what not to do
	//
	s3 := []byte{1, 2, 3, 4, 5, 6}
	s4 := s3 // <-- 2nd reference to the same underlying array
	for c, _ := range s4 {
		s3[c] = 0
	}
	println("-----------")
	fmt.Printf("s3: %v\n", s3)
	fmt.Printf("s4: %v\n", s4)
}
