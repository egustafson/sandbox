package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	d := Demo{
		ID: 31,
		At: Now(),
	}
	fmt.Printf("   d: %s\n", d.String())

	buf, _ := json.Marshal(d)
	fmt.Printf("json: %s\n", string(buf))

	newD := new(Demo)
	json.Unmarshal(buf, newD)
	fmt.Printf("newD: %s\n", d.String())
}
