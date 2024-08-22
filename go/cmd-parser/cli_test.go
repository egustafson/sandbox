package command

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCommandRegistry(t *testing.T) {
	cr := NewCommandRegistry()
	assert.NotNil(t, cr.CommandMap)
	assert.Zero(t, len(cr.CommandMap))
}

func TestSingleRegister(t *testing.T) {
	cr := NewCommandRegistry()
	ch := &CommandHandler{
		Use:     "base sub cmd [arg]",
		Handler: "base-sub-cmd",
	}
	err := cr.Register(ch)
	assert.Nil(t, err)
}

func TestDuplicateRegisterError(t *testing.T) {
	cr := NewCommandRegistry()
	ch := &CommandHandler{
		Use:     "base sub cmd [arg]",
		Handler: "base-sub-cmd",
	}
	err := cr.Register(ch) // first registration
	assert.Nil(t, err)     // <-- success
	err = cr.Register(ch)  // SECOND registration
	assert.Error(t, err)
}

func TestLookup(t *testing.T) {
	cr := NewCommandRegistry()
	cr.Register(&CommandHandler{
		Use:     "a b c [d] <e> -f",
		Handler: "a-b-c",
	})
	tests := []struct {
		in   string
		pass bool
		out  string
	}{{
		in:   "a b c -f arg",
		pass: true,
		out:  "a-b-c",
	}, {
		in:   "a b c d e -f", // exactly the cmd line as speced by the a-b-c handler
		pass: true,
		out:  "a-b-c",
	}, {
		in:   "x b c d e -f", // first cmd fails, requires iteration
		pass: false,
		out:  "",
	}, {
		in:   "", // empty cmd-line
		pass: false,
		out:  "",
	}, {
		in:   "  ", // whitespace filled cmd-line (i.e. empty)
		pass: false,
		out:  "",
	}}
	for _, tt := range tests {
		t.Run(tt.out, func(t *testing.T) {
			r, ok := cr.Lookup(tt.in)
			assert.Equal(t, tt.pass, ok)
			if tt.pass {
				assert.Equal(t, tt.out, r.Handler)
			}

		})
	}
}

func TestParseCmdLine(t *testing.T) {
	tests := []struct {
		in   string
		cmd  []string
		args []string
	}{{
		in:   "a",
		cmd:  []string{"a"},
		args: []string{},
	}, {
		in:   "a -f",
		cmd:  []string{"a"},
		args: []string{"-f"},
	}, {
		in:   "", // empty
		cmd:  []string{},
		args: []string{},
	}, {
		in:   "  ", // whitespace => empty
		cmd:  []string{},
		args: []string{},
	}, {
		in:   "a b",
		cmd:  []string{"a", "b"},
		args: []string{},
	}, {
		in:   "a b <p>",
		cmd:  []string{"a", "b"},
		args: []string{"<p>"},
	}, {
		in:   "a b [p]",
		cmd:  []string{"a", "b"},
		args: []string{"[p]"},
	}, {
		in:   "a very long sub command with -flags [opt-param]",
		cmd:  []string{"a", "very", "long", "sub", "command", "with"},
		args: []string{"-flags", "[opt-param]"},
	}, {
		in:   "short <param> [opt-param] -flag <flag-param> [opt2]",
		cmd:  []string{"short"},
		args: []string{"<param>", "[opt-param]", "-flag", "<flag-param>", "[opt2]"},
	}, {
		in:   "-f no-command",
		cmd:  []string{},
		args: []string{"-f", "no-command"},
	}}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("'%s'", tt.in), func(t *testing.T) {
			c, a := parseCmdLine(tt.in)
			assert.Equal(t, tt.cmd, c, "command sequence")
			assert.Equal(t, tt.args, a, "arguments")
		})
	}
}
