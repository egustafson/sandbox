package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "demo",
		Short: "A pub and sub demo program for nsq",
	}

	aboutCmd = &cobra.Command{
		Use:   "about",
		Short: "information about demo",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("demo program for nsq pub and sub")
		},
	}
)

func init() {
	rootCmd.AddCommand(aboutCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
