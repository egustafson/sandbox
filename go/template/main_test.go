package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeRE(t *testing.T) {
	domain := "elfwerks"
	re, err := MakeRE(domain)
	if assert.Nil(t, err) {
		source := "cn=elves,ou=workers,dc=elfwerks,dc=org"

		m := re.FindStringSubmatch(source)
		if assert.Len(t, m, 2) {
			assert.Equal(t, "elves", m[1])
		}
	}
}
