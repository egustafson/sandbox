package tunables_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/egustafson/sandbox/go/tunable2/tunables"
)

func TestTunable(t *testing.T) {
	baseIntValue := 2
	v := tunables.IntegerVal(baseIntValue)
	assert.Equal(t, baseIntValue, int(v))
	assert.True(t, int(v) == baseIntValue)
	assert.Equal(t, "2", v.String())

	v2 := tunables.StringVal("hello")
	assert.Equal(t, "hello", string(v2))
	assert.True(t, v2 == "hello")
	assert.Equal(t, "hello", v2.String())
}

func TestPolymorphic(t *testing.T) {
	var v tunables.Value = tunables.IntegerVal(123)
	// i, ok := v.(int)
	// assert.True(t, v.(int) == 123)
	// assert.True(t, int(v) == 123)
	assert.NotNil(t, v)
	assert.Equal(t, "123", v.String())

	v.Set("321")
	assert.NotNil(t, v)
	assert.Equal(t, "321", v.String())
}
