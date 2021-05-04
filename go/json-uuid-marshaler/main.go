package main

import (
	"fmt"

	"github.com/google/uuid"
)

type Demo struct {
	ID   int       `json:"id"`
	UUID uuid.UUID `json:"uuid"` // implementes Text.Marshaler interface
}

func (d *Demo) String() string {
	return fmt.Sprintf("{ id: %d, uuid: %s }", d.ID, d.UUID)
}
