package main

import (
	_ "embed" // see sandbox/go/embed-file
	"fmt"
	"log"

	"gopkg.in/yaml.v3"
)

//go:embed embedded.yaml
var embedded_yaml []byte

func main() {
	var doc map[string]any = make(map[string]any)

	err := yaml.Unmarshal(embedded_yaml, doc)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	fmt.Printf("doc is a %T\n", doc)
	demo := doc["demo"].(map[string]any)
	fmt.Printf("  demo is a %T\n", demo)
	for k, v := range demo {
		fmt.Printf("    %s: %v (is a %T)\n", k, v, v)
	}
}
