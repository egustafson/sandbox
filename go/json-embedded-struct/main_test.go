package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	a := Advanced{
		Basic:  Basic{Name: "my-name", Value: 21},
		Secret: "a-secret",
	}
	fmt.Printf("   a: %v\n", a)

	buf, _ := json.Marshal(a)
	fmt.Printf("json: %s\n", string(buf))

	newA := new(Advanced)
	json.Unmarshal(buf, newA)
	fmt.Printf("newA: %v\n", newA)
}
