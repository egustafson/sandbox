package main

import (
	"testing"
)

func TestWalk(t *testing.T) {
	r := &Root{
		RootID: "root-id",
		Value:  "value",
		Child1: &ChildA{
			Avalue:       "a-value",
			AnotherValue: "another-value",
			Children: []*ChildB{&ChildB{
				StrValue: "str-value",
				IntValue: 200,
			}},
		},
		Child2: &ChildB{
			StrValue: "str-value-child2",
			IntValue: 100,
		},
	}
	Print(r)
	Walk(r)
	Print(r)
}
