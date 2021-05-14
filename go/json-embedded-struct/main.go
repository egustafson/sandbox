package main

import ()

type Basic struct {
	Name  string `json:"name"`
	Value int    `json:"val"`
}

type Advanced struct {
	Basic
	Secret string `json:"secret"`
}

func main() {
	// no-op
}
