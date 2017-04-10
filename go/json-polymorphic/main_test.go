package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

type SimpleJson struct {
	Key string `json:"key"`
	Val string `json:"value"`
}

func TestSimpleJson(t *testing.T) {
	v := SimpleJson{Key: "key-string", Val: "val-sting"}
	j, _ := json.Marshal(v)
	fmt.Printf("%s\n", j)
	v.Key = ""
	v.Val = ""
	json.Unmarshal(j, &v)
	fmt.Printf("%+v\n", v)
}

// ////////////////////////////////////////////////////////////

type OneBodyStruct struct {
	OneField string `json:"onef"`
}

type TwoBodyStruct struct {
	TwoVersion int `json:"twov"`
}

type Message struct {
	Cmd string         `json:"cmd"`
	Ver int            `json:"version"`
	One *OneBodyStruct `json:"one"`
	Two *TwoBodyStruct `json:"two"`
}

func TestPolyJsonOne(t *testing.T) {
	v := Message{Cmd: "one", Ver: 21}
	v.One = &OneBodyStruct{OneField: "field-value"}
	j, _ := json.Marshal(v)
	fmt.Printf("%s\n", j)

	var v2 Message
	json.Unmarshal(j, &v2)
	fmt.Printf("%+v \\ %+v\n", v2, v2.One)
}

func TestPolyJsonTwo(t *testing.T) {
	v := Message{Cmd: "random", Ver: 49}
	v.Two = &TwoBodyStruct{TwoVersion: 5}
	j, _ := json.Marshal(v)
	fmt.Printf("%s\n", j)

	var v2 Message
	json.Unmarshal(j, &v2)
	fmt.Printf("%+v \\ %+v\n", v2, *v2.Two)
}
