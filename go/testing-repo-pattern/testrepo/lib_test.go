package testrepo_test

import (
	"testing"

	"github.com/egustafson/sandbox/go/testing-repo-pattern/testlib"
)

/* This is a proto-example of a package that ONLY has golang tests.  It is
 * intended to be used and demonstrated in conjunction with the 'testlib'
 * package.
 */

func TestNothing(t *testing.T) {
	// This test always passes, its just a placeholder.
}

func TestPositive(t *testing.T) {
	r := testlib.Positive(1, 2)
	if r < 0 {
		t.Fail()
	}
}
