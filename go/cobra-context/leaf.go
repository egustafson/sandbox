package main

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
)

var leafCmd = &cobra.Command{
	Use:   "leaf",
	Short: "demo 'leaf' command that receives resource",
	RunE:  doLeaf,
}

func init() {
	rootCmd.AddCommand(leafCmd)
}

func doLeaf(cmd *cobra.Command, args []string) error {
	res, err := getResource(cmd)
	if err != nil {
		return err
	}

	fmt.Printf("Resource: %+v\n", *res)

	return nil
}

func getResource(cmd *cobra.Command) (*Resource, error) {
	ctx := cmd.Context()
	if ctx == nil {
		return nil, errors.New("cmd.Context() is nil")
	}
	if rh, ok := ctx.Value("rh").(*ResourceHolder); ok {
		return rh.Resource, nil
	}
	return nil, errors.New("ResourceHolder not present on context")
}
