package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJwtClaim(t *testing.T) {

	jwt1 := MakeJwtClaim1("name-1")
	jwt2 := MakeJwtClaim2("name-2", "other-name")

	claimA := Validate1(jwt1) // should succeed
	assert.NotNil(t, claimA)

	claimB := Validate1(jwt2) // expect to pass, jwt2 is a superset of jwt1
	assert.NotNil(t, claimB)

	claimC := Validate2(jwt2) // should succeed
	assert.NotNil(t, claimC)

	// This passes
	claimD := Validate2(jwt1) // expect to fail, jwt1 is only a SUB-set of jwt2
	assert.NotNil(t, claimD)
}
