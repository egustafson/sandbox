package main

import (
	"fmt"
)

func main() {
	frames := make([]string, 2)
	frames = append(frames, "first-append")
	fmt.Printf("len(frames) = %d: %v\n", len(frames), frames)
	frames[0] = "0"
	frames[1] = "1"
	fmt.Printf("len(frames) = %d: %v\n", len(frames), frames)
}
