package oldhandler

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRootHandler(t *testing.T) {
	response, err := RootHandler([]string{
		"cmd",
		"--arg", "val",
		"subcmd",
		"--subflag",
	})
	if assert.Nil(t, err) {
		fmt.Print(response)
	}
}
