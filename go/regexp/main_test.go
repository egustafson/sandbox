package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRE(t *testing.T) {
	cases := []struct {
		name  string
		input string
		id    string
		div   string
	}{
		{
			name:  "service account",
			input: "cn=group1,ou=Services,ou=division2,ou=Environments,dc=elfwerks,dc=org",
			id:    "group1",
			div:   "division2",
		},
		{
			name:  "service account with dash",
			input: "cn=group-1,ou=Services,ou=division2,ou=Environments,dc=elfwerks,dc=org",
			id:    "group-1",
			div:   "division2",
		},
		{
			name:  "service account with dash in division",
			input: "cn=group-1,ou=Services,ou=division-2,ou=Environments,dc=elfwerks,dc=org",
			id:    "group-1",
			div:   "division-2",
		},
		{
			name:  "person account",
			input: "uid=user3,ou=People,dc=elfwerks,dc=org",
			id:    "user3",
			div:   "",
		},
		{
			name:  "person account with dash",
			input: "uid=user-3,ou=People,dc=elfwerks,dc=org",
			id:    "user-3",
			div:   "",
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			id, div := Parse(tt.input)
			fmt.Printf("\n%s\n", tt.input)
			fmt.Printf("  ->match: '%s', '%s'\n", id, div)
			assert.Equal(t, tt.id, id)
			assert.Equal(t, tt.div, div)
		})
	}
	fmt.Println("")
}
