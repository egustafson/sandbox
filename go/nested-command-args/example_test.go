package main_test

import (
	"context"
	"fmt"

	"github.com/egustafson/sandbox/go/nested-command-args/cli"
	"github.com/spf13/pflag"
)

func demoParentHandler(ctx context.Context, resp *cli.Response, req *cli.Request) {
	// Process possible flags (with pflag)
	fs := pflag.NewFlagSet("parent-handler-flags", pflag.ContinueOnError)
	fs.SetInterspersed(false)
	parentFlag := fs.String("parent-flag", "", "usage: demo flag")
	if err := fs.Parse(req.Args); err != nil {
		resp.Err = err
		return
	}
	req.Args = fs.Args() // update req's Args

	// set header values
	resp.Headers.Append("handler", "parent-handler")
	if len(*parentFlag) > 0 {
		resp.Headers.Append("parent-flag", *parentFlag)
	}

	// if args remain it must be a sub-command
	if len(req.Args) > 0 {
		req.CmdHandler.DescendHandlers(ctx, resp, req)
	} else {
		resp.Body = "simple text response"
	}
}

func demoLeafHandler(ctx context.Context, resp *cli.Response, req *cli.Request) {

	// As this Handler does not descend, it is safe to ignore flags and further
	// arguments.  Best practice would be to process them.

	resp.Headers.Append("handler", "leaf-handler")
	resp.Body = struct {
		ID  int
		Val string
	}{
		ID:  123,
		Val: "leaf-response-value",
	}
}

var (
	demoHandlers = cli.NewCliHandler()
	demoCtx      = context.Background()
)

func init() {
	// Register the two handlers in this example.  'leaf' is a child of 'demo'
	demoHandlers.Register(cli.Handler{
		Command:   "demo",
		HandlerFn: demoParentHandler,
	})
	demoHandlers.Register(cli.Handler{
		Command:   "demo:leaf", // called as:  'demo leaf -flags'
		HandlerFn: demoLeafHandler,
	})
}

func RunCmdline(cmdline string) {
	r, err := demoHandlers.Execute(demoCtx, cmdline)
	if err != nil {
		panic("example failed")
	}

	for _, h := range r.Headers() {
		fmt.Printf("%s: %s\n", h.Key, h.Value)
	}
	fmt.Println("")
	fmt.Printf("%s\n", r.Body())
	fmt.Println("")
}

func Example_parent_only() {
	cmdline := "demo +parent-flag parent-flag-value"
	RunCmdline(cmdline)
	// Output:
	// handler: parent-handler
	// parent-flag: parent-flag-value
	// Content-Type: text/plain
	//
	// simple text response
	//
}

func Example_leaf() {
	cmdline := "demo leaf +flag flag-value"
	RunCmdline(cmdline)
	// Output:
	// handler: parent-handler
	// handler: leaf-handler
	// Content-Type: text/yaml
	//
	// id: 123
	// val: leaf-response-value
	//
}
