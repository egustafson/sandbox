package cli_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/slices"

	"github.com/egustafson/sandbox/go/nested-command-args/cli"
)

var (
	testHeaders = cli.Headers{
		{"k", "v"},
		{"Content-Type", "any/any"},
		{"k-many", "value-1"},
		{"k-many", "value-2"},
	}
)

func TestHeadersKeys(t *testing.T) {
	keys := testHeaders.Keys()
	for _, k := range keys {
		assert.True(t, testHeaders.Contains(k))
	}
}

func TestHeadersContains(t *testing.T) {
	assert.True(t, testHeaders.Contains("k"))
	assert.False(t, testHeaders.Contains("missing"))
}

func TestHeadersGet(t *testing.T) {
	assert.Equal(t, testHeaders.Get("k"), "v")
	assert.Equal(t, testHeaders.Get("missing"), "")
}

func TestHeadersGetAll(t *testing.T) {
	all := testHeaders.GetAll("k-many")
	assert.True(t, slices.Contains(all, "value-1"))
	assert.True(t, slices.Contains(all, "value-2"))
}

func TestHeadersSet(t *testing.T) {
	assert.False(t, testHeaders.Contains("k2"))
	testHeaders.Set("k2", "v1")
	assert.Equal(t, testHeaders.Get("k2"), "v1")
	testHeaders.Set("k2", "v2")
	assert.Equal(t, testHeaders.Get("k2"), "v2")
}

func TestHeadersAppend(t *testing.T) {
	key := "k-append"
	assert.False(t, testHeaders.Contains(key))
	testHeaders.Append(key, "val-1")
	testHeaders.Append(key, "val-2")
	values := testHeaders.GetAll(key)
	assert.True(t, len(values) == 2)
	testHeaders.Append(key, "val-3")
	values = testHeaders.GetAll(key)
	assert.True(t, slices.Contains(values, "val-1"))
	assert.True(t, slices.Contains(values, "val-2"))
	assert.True(t, slices.Contains(values, "val-3"))
}

func TestHeadersDelete(t *testing.T) {
	testKey := "key-add-remove"
	assert.False(t, testHeaders.Contains(testKey))
	testHeaders.Set(testKey, "value")
	assert.True(t, testHeaders.Contains(testKey))
	testHeaders.Delete(testKey)
	assert.False(t, testHeaders.Contains(testKey))
}

func TestHeadersDeleteAll(t *testing.T) {
	testKey := "key-add-remove-all"
	assert.False(t, testHeaders.Contains(testKey))

	testHeaders.Set(testKey, "val-1")
	testHeaders.Append(testKey, "val-2")
	testHeaders.Append(testKey, "val-3")
	assert.True(t, testHeaders.Contains(testKey))
	values := testHeaders.GetAll(testKey)
	assert.True(t, len(values) == 3)

	testHeaders.DeleteAll(testKey)
	assert.False(t, testHeaders.Contains(testKey))
	values = testHeaders.GetAll(testKey)
	assert.True(t, len(values) == 0)
}
