package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

const (
	IterCount = 100000
	FanOut    = 4000
	KillProb  = 99998 // (0.99998) 1 in 100k
)

func TestDemo(t *testing.T) {
	p := NewProducer(IterCount)

	for ii := 0; ii < FanOut; ii++ {
		w := p.Watch()
		go func(id int) {
			count := 0
			//fmt.Printf("start id: %d\n", id)
			for n := range w.Chan() {
				// process n
				if rand.Intn(100000) > KillProb {
					// 					fmt.Printf("w(%d) - closing\n", id)
					go func() {
						w.Close() // <-- close the chan before it's complete
					}()
					break
				}
				_ = n
				count++
			}
			fmt.Printf("w(%d) exited at count: %d\n", id, count)
		}(ii)
	}

	startWatcherCount := p.NumWatchers()
	fmt.Printf("%d watchers at the beginnig.\n", startWatcherCount)
	p.Start() // synchronous run
	fmt.Printf("Production complete.")
	fmt.Printf("%d watchers at the beginnig.\n", startWatcherCount)
	fmt.Printf("%d watchers at the end.\n", p.NumWatchers())
	time.Sleep(2 * time.Second)
	fmt.Printf("exiting.")
}
