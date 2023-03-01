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
	testHandlers = cli.NewCliHandler()
	ctx          = context.Background()
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

func StringResponseTestHandler(ctx context.Context, args []string) *cli.Result {
	return &cli.Result{
		Body: "testing-result",
	}
}

func YamlResponseTestHandler(ctx context.Context, args []string) *cli.Result {
	return &cli.Result{
		Body: testYamlResponse{
			ID:  "id-1",
			Val: 123,
			Msg: "testing-yaml-handler-message",
		},
	}
}

func ArgsResponseTestHandler(ctx context.Context, args []string) *cli.Result {
	if len(args) < 1 {
		return &cli.Result{
			Body: args,
			Err:  errors.New("no arguments present"),
		}
	}
	return &cli.Result{Body: args}
}

func HeaderZeroBodyResponseTestHandler(ctx context.Context, args []string) *cli.Result {
	r := &cli.Result{Body: ""}
	r.Headers.Set("test-header", "test-value")
	return r
}

// --  Test Cases  ----------------------------------------

func TestExecuteStringResponseCmd(t *testing.T) {
	r, err := testHandlers.Execute(ctx, "testing")
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
	r, err := testHandlers.Execute(ctx, "tyaml")
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
	r, err := testHandlers.Execute(ctx, "args a b c d")
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

	_, err = testHandlers.Execute(ctx, "args") // no args == error
	assert.NotNil(t, err)
}

func TestExecuteHeaderZeroBodyCmd(t *testing.T) {
	r, err := testHandlers.Execute(ctx, "zero-body")
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
