package oldhandler

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

	// build a FlagSet
	fs := pflag.NewFlagSet("cmd-flags", pflag.ContinueOnError)
	fs.SetInterspersed(false)
	a := fs.String("arg", "default", "arg usage")

	// parse the Flags
	if err := fs.Parse(args); err != nil {
		return nil, err
	}

	// Process/execute behavior for this command
	fmt.Printf("CMD: --arg=\"%s\"\n", *a)
	result := make(map[string]any)
	result["arg"] = *a

	// if arguments remain, recurse:  sub-command
	if fs.NArg() > 0 {
		sub, err := HandleCommand(CmdHandlerMap, fs.Args())
		if err != nil {
			return nil, err
		}
		// append sub-command's results to this result
		result[fs.Arg(0)] = sub
	}

	// return the result(s)
	return result, nil
}
