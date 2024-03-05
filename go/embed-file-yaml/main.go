package main

import (
	_ "embed"
	"fmt"
)

//go:embed embedded.yaml
var embedded_yaml string // or []byte, or a list of files, see godoc:embed

func main() {
	fmt.Println(embedded_yaml)
}
