package tunable_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/egustafson/sandbox/go/tunable-prototype/tunable"
)

func TestTunableBool(t *testing.T) {
	var boolFlag bool = false

	tuneset := tunable.NewTunableSet("test-tunable-set")
	assert.NotNil(t, tuneset)

	tuneset.DefineBoolVar(&boolFlag, "test-bool-tunable", true)
	assert.True(t, boolFlag)
}

func TestTunableInt(t *testing.T) {
	var intFlag int = 21

	tuneset := tunable.NewTunableSet("test-tunable-set")
	assert.NotNil(t, tuneset)

	tuneset.DefineIntVar(&intFlag, "test-int-tunable", 42)
	assert.Equal(t, 42, intFlag)
}
