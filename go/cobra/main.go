package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "demo",
	Short: "A demo program that uses Cobra.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cobra ROOT command.")
	},
}

var childCmd = &cobra.Command{
	Use:   "child",
	Short: "A child command for the demo Cobra app.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("CHILD cobra command.")
	},
}

func init() {
	RootCmd.AddCommand(childCmd)
}

func main() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
