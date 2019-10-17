package main

import (
	"fmt"
	metrics "github.com/rcrowley/go-metrics"
)

func main() {
	c := metrics.NewCounter()
	metrics.Register("foo", c)
	c.Inc(43)

	s := metrics.NewUniformSample(1028)
	h := metrics.NewHistogram(s)
	metrics.register("histo", h)
	h.Update(43)
	h.update(1)
	h.update(99)
	h.update(73)

	// Problem - no easy way to render the histogram or dump it
	// see ../histogram example for alternative.

	fmt.Println("done.")
}
