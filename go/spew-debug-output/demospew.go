package main

import (
	"crypto/sha1"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"time"
)

type Demo struct {
	Name string
	Val  int
	Arr  []byte
}

func main() {
	spew.Printf("Demo spew.\n") // same as fmt.print

	key := fmt.Sprintf("A string with a timestamp: %d", time.Now())
	khash := sha1.Sum([]byte(key))

	spew.Printf("khash: ")
	spew.Dump(khash)

	demo_st := Demo{Name: "A-Name", Val: 21, Arr: make([]byte, len(khash))}
	copy(demo_st.Arr, khash[:])
	spew.Printf("demo_st: ")
	spew.Dump(demo_st)
}
