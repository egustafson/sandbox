package main

import (
	"encoding/json"
	"fmt"
	"log"

	"gopkg.in/yaml.v2"
)

// type HealthState defined in health_state.go

type Health struct {
	Id     string            `yaml:"id" json:"id"`
	Health HealthState       `yaml:"health" json:"health"`
	Meta   map[string]string `yaml:",inline,omitempty" json:"meta,omitempty"`
	Child  []Health          `yaml:"components,omitempty" json:"components,omitempty"`
}

// This example sets the 'Meta' field to 'inline' for YAML.  The
// effect is that any entries in Meta are rendered as if they are part
// of the Health struct.

func main() {

	h := &Health{
		Id:     "daemon",
		Health: Unhealthy,
		Meta:   map[string]string{"version": "v1.2.3"},
		Child: []Health{
			Health{Id: "child", Health: Ok},
			Health{Id: "child2", Health: Unknown},
		},
	}

	y, err := yaml.Marshal(&h)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- h dump:\n%s...\n", string(y))

	j, err := json.MarshalIndent(&h, "", "  ")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("%s\n", string(j))

	fmt.Println("done.")
}
