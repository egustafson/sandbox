package oldhandler

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

	// build a FlagSet
	fs := pflag.NewFlagSet("subcmd-flags", pflag.ContinueOnError)
	fs.SetInterspersed(false)
	sf := fs.Bool("subflag", false, "subflag usage")

	// parse the Flags
	if err := fs.Parse(args); err != nil {
		return nil, err
	}

	// Note: no check for sub-commands.  This handler assumes/knows there are no
	// sub-commands beneath it.

	// Process/execute behavior for this command
	fmt.Printf("SUBCMD: --subflag=%v\n", *sf)
	result := make(map[string]any)
	result["subflag"] = *sf

	// return the results
	return result, nil
}
