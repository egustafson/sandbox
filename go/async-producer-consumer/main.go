package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	b := NewBuffer()

	go produce(b)
	consume(b)

	fmt.Println("done.")
}

func produce(b Buffer) {
	rand.Seed(time.Now().UnixNano())
	for ii := 0; ii < 100; ii++ {
		time.Sleep(time.Duration(100*(rand.Intn(9)+1)) * time.Millisecond)
		b.Push(rand.Int())
		if s := b.Size(); s > 0 {
			fmt.Printf("size: %d\n", s)
		}
	}
	b.Close()
}

func consume(b Buffer) {
	var more bool = true
	for more {
		time.Sleep(time.Duration(100*(rand.Intn(9)+1)) * time.Millisecond)
		var v interface{}
		v, more = b.Pull()
		fmt.Printf("%v\n", v)
	}
}
