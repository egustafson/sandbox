package main

import (
	"errors"

	"github.com/spf13/cobra"
)

// Resource is just an arbitrary struct that holds some arbitrary
// information and is used as a placeholder for a resource to be
// attached to the cobra.Command.
type Resource struct {
	Name string
	Id   int
}

// ResourceHolder is a holder of a resource.  It should be attached to
// the Command Context at Execute() time.  cobra does not expose the
// Command's context in a way that the context can be updated as it
// passes through each Command in the (cobra) chain of execution.
type ResourceHolder struct {
	Resource *Resource
}

var rootCmd = &cobra.Command{
	Use:   "demo",
	Short: "root command for the example",
	// no Run - this is just a trunk in the tree of commands.
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		res := &Resource{
			Name: "random-resource-name",
			Id:   42,
		}
		return setResource(cmd, res)
	},
}

func setResource(cmd *cobra.Command, res *Resource) error {
	ctx := cmd.Context()
	if ctx == nil {
		return errors.New("cmd.Context() is nil")
	}
	if rh, ok := ctx.Value("rh").(*ResourceHolder); ok {
		rh.Resource = res
		return nil
	}
	return errors.New("ResourceHolder not present on context")
}
