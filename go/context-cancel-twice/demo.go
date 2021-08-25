package main

import (
	"context"
	"fmt"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	fmt.Println("context created")
	cancel()
	cancel()
	cancel()
	cancel()
	cancel()
	cancel()
	cancel()
	//
	// The above works.  It appears to be safe to double-cancel a context
	//

	select {
	case <-ctx.Done():
	}
	fmt.Println("context Done")
}
