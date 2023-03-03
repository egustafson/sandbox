package cli_test

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"

	"github.com/egustafson/sandbox/go/nested-command-args/cli"
)

func init() {
	testHandlers.Register(cli.Handler{
		Command:   "top-cmd",
		HandlerFn: DiagnosticResponseHandler,
	})
	testHandlers.Register(cli.Handler{
		Command:   "nest1:nest2",
		HandlerFn: DiagnosticResponseHandler,
	})
	testHandlers.Register(cli.Handler{
		Command:   "middle-layer",
		HandlerFn: MiddleLayerHandler,
	})
	testHandlers.Register(cli.Handler{
		Command:   "middle-layer:child",
		HandlerFn: DiagnosticResponseHandler,
	})
}

type diagnosticYamlResponse struct {
	Command string   `yaml:"cmd"`
	Args    []string `yaml:"args"`
}

func DiagnosticResponseHandler(ctx context.Context, resp *cli.Response, req *cli.Request) {
	resp.Body = diagnosticYamlResponse{
		Command: req.Command,
		Args:    req.Args,
	}
}

func MiddleLayerHandler(ctx context.Context, resp *cli.Response, req *cli.Request) {
	// Add a header with the timestamp.
	resp.Headers.Append("middle-layer", time.Now().Format(time.RFC3339))
	if len(req.Args) > 0 {
		req.CmdHandler.DescendHandlers(ctx, resp, req)
	} else { // mimic DiagnosticResponseHandler to make test work
		DiagnosticResponseHandler(ctx, resp, req)
	}
}

// --  Test Cases  ----------------------------------------------------------

func TestCmdLineInOut(t *testing.T) {
	tests := []struct {
		name    string
		inCmd   string
		outCmd  string
		outArgs []string
	}{
		{
			name:    "single top level cmd",
			inCmd:   "top-cmd",
			outCmd:  "top-cmd",
			outArgs: []string{},
		}, {
			name:    "simple nested cmd",
			inCmd:   "nest1 nest2",
			outCmd:  "nest1:nest2",
			outArgs: []string{},
		}, {
			name:    "simple arg, translated",
			inCmd:   "top-cmd +long_arg ++short ++x +long=arg ++f=filename",
			outCmd:  "top-cmd",
			outArgs: []string{"--long_arg", "-short", "-x", "--long=arg", "-f=filename"},
		}, {
			name:    "middle handler with no further descent",
			inCmd:   "middle-layer",
			outCmd:  "middle-layer",
			outArgs: []string{},
		}, {
			name:    "middle handler to diagnostic handler, no args",
			inCmd:   "middle-layer child",
			outCmd:  "middle-layer:child",
			outArgs: []string{},
		}, {
			name:    "middle handler to diagnostic handler with args",
			inCmd:   "middle-layer child +argument value ++flags",
			outCmd:  "middle-layer:child",
			outArgs: []string{"--argument", "value", "-flags"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, err := testHandlers.Execute(testCtx, tt.inCmd)
			assert.Nil(t, err)
			if assert.NotNil(t, r) {
				var diag diagnosticYamlResponse
				err = yaml.Unmarshal([]byte(r.Body()), &diag)
				if assert.Nil(t, err) {
					assert.Equal(t, diag.Command, tt.outCmd)
					if assert.Equal(t, len(diag.Args), len(tt.outArgs)) {
						for ii := range diag.Args {
							assert.Equal(t, diag.Args[ii], tt.outArgs[ii])
						}
					}
				}
			}
		})
	}
}

func TestMissingNestedHandler(t *testing.T) {
	badCmd := "bogus nested cmd"
	_, err := testHandlers.Execute(testCtx, badCmd)
	if assert.NotNil(t, err) {
		// The returned error should cite the command it can not find
		assert.Contains(t, err.Error(), strings.ReplaceAll(badCmd, " ", ":"))
	}
}

func TestMiddleHandlerHeader(t *testing.T) {
	testCmd := "middle-layer child +argument value"
	r, err := testHandlers.Execute(testCtx, testCmd)
	assert.Nil(t, err)
	if assert.NotNil(t, r) {
		headers := cli.Headers(r.Headers())
		assert.True(t, len(headers) > 0)
		assert.True(t, len(headers.GetAll("middle-layer")) == 1)
	}
}
