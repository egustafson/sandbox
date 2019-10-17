package main

import (
	"fmt"
	"github.com/codahale/hdrhistogram"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	h := hdrhistogram.New(0, 1000000, 2)

	// Create 2 distributions in the Histogram

	// Synthesize a small cluster 0-1000
	//
	for ii := 0; ii < 20; ii++ {
		h.RecordValue(rand.Int63n(1000))
	}

	// Synthesize a larger, normal distribution 0-max
	//
	for ii := 0; ii < 10; ii++ {
		h.RecordValue(rand.Int63n(1000000))
	}

	fmt.Printf("mean: %f\n", h.Mean())

	for _, bar := range h.Distribution() {
		if bar.Count > 0 {
			fmt.Printf("%d - %d : %d\n", bar.From, bar.To, bar.Count)
		}
	}
}
