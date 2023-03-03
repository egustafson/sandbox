package cli_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"

	"github.com/egustafson/sandbox/go/nested-command-args/cli"
)

type testYamlResponse struct {
	ID  string `yaml:"id"`
	Val int    `yaml:"value"`
	Msg string `yaml:"message"`
}

var (
	testHandlers = cli.NewCliHandler()  // used by ALL cli_xxx_test.go tests
	testCtx      = context.Background() // used by ALL cli_xxx_test.go tests
)

func init() {
	testHandlers.Register(cli.Handler{
		Use:       "testing",
		Command:   "testing",
		Short:     "testing returns a simple result, as text",
		HandlerFn: StringResponseTestHandler,
	})
	testHandlers.Register(cli.Handler{
		Use:       "tyaml",
		Command:   "tyaml",
		Short:     "tyaml returns a stringified yaml response",
		HandlerFn: YamlResponseTestHandler,
	})
	testHandlers.Register(cli.Handler{
		Command:   "args",
		HandlerFn: ArgsResponseTestHandler,
	})
	testHandlers.Register(cli.Handler{
		Command:   "zero-body",
		HandlerFn: HeaderZeroBodyResponseTestHandler,
	})
}

func StringResponseTestHandler(ctx context.Context, resp *cli.Response, req *cli.Request) {
	resp.Body = "testing-result"
}

func YamlResponseTestHandler(ctx context.Context, resp *cli.Response, req *cli.Request) {
	resp.Body = testYamlResponse{
		ID:  "id-1",
		Val: 123,
		Msg: "testing-yaml-handler-message",
	}
}

func ArgsResponseTestHandler(ctx context.Context, resp *cli.Response, req *cli.Request) {
	resp.Body = req.Args
	if len(req.Args) < 1 {
		resp.Err = errors.New("no arguments present")
	}
}

func HeaderZeroBodyResponseTestHandler(ctx context.Context, resp *cli.Response, req *cli.Request) {
	resp.Body = ""
	resp.Headers.Set("test-header", "test-value")
}

// --  Test Cases  ----------------------------------------------------------

func TestEmptyCmdLine(t *testing.T) {
	emptyCmdLineError := cli.EmptyCommandLineError(nil)

	tests := []struct {
		name    string
		cmdline string
	}{
		{
			name:    "blank line",
			cmdline: "",
		}, {
			name:    "new-line",
			cmdline: "\n",
		}, {
			name:    "new-lines",
			cmdline: "\n\n\n\n",
		}, {
			name:    "all blank lines",
			cmdline: " \n \n \n ",
		}, {
			name:    "blank line with tab",
			cmdline: "\t\n \n\t \n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := testHandlers.Execute(testCtx, tt.cmdline)
			assert.True(t, errors.As(err, &emptyCmdLineError))
		})
	}
}

func TestUnknownHandler(t *testing.T) {
	unknownCommandError := cli.UnknownCommandError(nil)
	_, err := testHandlers.Execute(testCtx, "unknown-command")
	assert.True(t, errors.As(err, &unknownCommandError))
}

func TestExecuteStringResponseCmd(t *testing.T) {
	r, err := testHandlers.Execute(testCtx, "testing")
	assert.Nil(t, err)
	if assert.NotNil(t, r) {
		assert.Greater(t, len(r.Body()), 0)
		hdrs := cli.Headers(r.Headers())
		if assert.True(t, hdrs.Contains("Content-Type")) {
			assert.Equal(t, hdrs.Get("Content-Type"), "text/plain")
		}
	}
}

func TestExecuteYamlResponseCmd(t *testing.T) {
	r, err := testHandlers.Execute(testCtx, "tyaml")
	assert.Nil(t, err)
	if assert.NotNil(t, r) {
		assert.Greater(t, len(r.Body()), 0)
		var body testYamlResponse
		err = yaml.Unmarshal([]byte(r.Body()), &body)
		if assert.Nil(t, err) {
			assert.Equal(t, body.ID, "id-1")
			assert.Equal(t, body.Val, 123)
			assert.Equal(t, body.Msg, "testing-yaml-handler-message")
		}
		hdrs := cli.Headers(r.Headers())
		if assert.True(t, hdrs.Contains("Content-Type")) {
			assert.Equal(t, hdrs.Get("Content-Type"), "text/yaml")
		}
	}
}

func TestExecuteArgsResponseCmd(t *testing.T) {
	r, err := testHandlers.Execute(testCtx, "args a b c d")
	assert.Nil(t, err)
	if assert.NotNil(t, r) {
		assert.Greater(t, len(r.Body()), 0)
		var body []string
		err = yaml.Unmarshal([]byte(r.Body()), &body)
		if assert.Nil(t, err) {
			if assert.Equal(t, len(body), 4) {
				assert.Equal(t, body[0], "a")
				assert.Equal(t, body[1], "b")
				assert.Equal(t, body[2], "c")
				assert.Equal(t, body[3], "d")
			}
		}
	}

	_, err = testHandlers.Execute(testCtx, "args") // no args == error
	assert.NotNil(t, err)
}

func TestExecuteHeaderZeroBodyCmd(t *testing.T) {
	r, err := testHandlers.Execute(testCtx, "zero-body")
	assert.Nil(t, err)
	if assert.NotNil(t, r) {
		assert.Equal(t, len(r.Body()), 0)
		hdrs := cli.Headers(r.Headers())
		if assert.True(t, hdrs.Contains("test-header")) {
			assert.Equal(t, hdrs.Get("test-header"), "test-value")
		}
		if assert.True(t, hdrs.Contains("Content-Type")) {
			assert.Equal(t, hdrs.Get("Content-Type"), "text/plain")
		}
	}
}
