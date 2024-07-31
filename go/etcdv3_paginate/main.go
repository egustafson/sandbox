package main

// This is a COBRA CLI example with two principal commands:  `demo` and
// `populate`

import (
	"os"

	"github.com/spf13/cobra"
)

const (
	matchPrefix = "paginate-demo/" // prefix for matching entries
	matchCount  = 20               // how many matching entries to create
)

var (
	rootCmd = &cobra.Command{
		Use: "paginate <sub-command>",
	}

	demoCmd = &cobra.Command{
		Use:  "paginate",
		RunE: doPaginate,
	}

	populateCmd = &cobra.Command{
		Use:  "populate",
		RunE: doPopulate,
	}

	config = initConfig()

	// CLI Flags
	verboseFlag bool = false
)

func init() {
	rootCmd.AddCommand(demoCmd)
	rootCmd.AddCommand(populateCmd)

	rootCmd.PersistentFlags().BoolVarP(&verboseFlag, "verbose", "v", false,
		"verbose output")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		// cobra will print an error message
		os.Exit(1)
	}
}
