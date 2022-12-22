package main

import (
	"fmt"
	"strings"

	"github.com/spf13/pflag"
)

func init() {
	CmdHandlerMap["subcmd"] = Handler{
		"subcmd",
		SubCmdHandler,
	}
}

func SubCmdHandler(args []string) (any, error) {
	fmt.Printf("SUBCMD: %s\n", strings.Join(args, " "))

	fs := pflag.NewFlagSet("subcmd-flags", pflag.ContinueOnError)
	fs.SetInterspersed(false)
	sf := fs.Bool("subflag", false, "subflag usage")

	if err := fs.Parse(args); err != nil {
		return nil, err
	}

	fmt.Printf("SUBCMD: --subflag=%v\n", *sf)

	result := make(map[string]any)
	result["subflag"] = *sf

	return result, nil
}
