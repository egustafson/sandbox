package main

// An example of how to run an HTTP server (http.ListenAndServe) in a secondary
// thread and do something "useful" (i.e. print 'tick') in the primary thread.

import (
	"log"
	"net/http"
	"strconv"
	"time"
)

type countHandler struct {
	count int
}

func (ch *countHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ch.count += 1
	c := strconv.Itoa(ch.count)
	w.Write([]byte("Count: " + c + "\n"))
	log.Printf("count: %s\n", c)
}

func launchHTTP(handler http.Handler) {
	log.Println("Listening...")
	http.ListenAndServe(":3000", handler)
}

func main() {
	mux := http.NewServeMux()

	ch := &countHandler{}
	mux.Handle("/count", ch)

	go launchHTTP(mux)
	for {
		time.Sleep(5 * time.Second)
		log.Println("tick.")
	}
}
