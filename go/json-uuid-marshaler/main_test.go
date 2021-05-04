package main

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/google/uuid"
)

func TestMain(t *testing.T) {
	d := &Demo{
		ID:   21,
		UUID: uuid.New(), // this can panic()
	}
	fmt.Printf("   d: %v\n", d)

	buf, _ := json.Marshal(d)
	fmt.Printf("json: %s\n", string(buf))

	newD := new(Demo)
	json.Unmarshal(buf, newD)
	fmt.Printf("newD: %v\n", d)
}
