package main

import (
	"fmt"
	"strings"

	"github.com/spf13/pflag"
)

func init() {
	RootHandlerMap["cmd"] = Handler{
		"cmd",
		CmdHandler,
	}
}

var CmdHandlerMap = make(HandlerMap)

func CmdHandler(args []string) (any, error) {
	fmt.Printf("CMD: %s\n", strings.Join(args, " "))

	fs := pflag.NewFlagSet("cmd-flags", pflag.ContinueOnError)
	fs.SetInterspersed(false)
	a := fs.String("arg", "default", "arg usage")

	if err := fs.Parse(args); err != nil {
		return nil, err
	}

	fmt.Printf("CMD: --arg=\"%s\"\n", *a)

	result := make(map[string]any)
	result["arg"] = *a

	if fs.NArg() > 0 {
		sub, err := HandleCommand(CmdHandlerMap, fs.Args())
		if err != nil {
			return nil, err
		}
		result[fs.Arg(0)] = sub
	}
	return result, nil
}
