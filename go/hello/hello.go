package main

import (
	"fmt"
	// in the same GitHub repo as this program
	"github.com/egustafson/sandbox/go/demopkg"
)

func main() {
	var msg = demopkg.GetMessage()
	fmt.Printf(msg)
}
