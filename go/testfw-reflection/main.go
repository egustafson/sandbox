package main

import (
	"fmt"

	"github.com/egustafson/go/testfw-reflection/demofwk"
)

func main() {
	initLogging()

	demofwk.Run()
	fmt.Println("done.")
}
