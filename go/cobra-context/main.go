package main

import (
	"context"
	"os"
)

func main() {
	rh := new(ResourceHolder)
	ctx := context.WithValue(context.Background(), "rh", rh)
	if err := rootCmd.ExecuteContext(ctx); err != nil {
		os.Exit(1)
	}
}
