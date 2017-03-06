package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
)

var port = flag.Int("port", 9900, "listen on port (default: 9900)")

type result struct {
	Status string  `json:"status"`
	Remote string  `json:"remote-addr"`
	Method string  `json:"method"`
	URI string     `json:"uri"`
}

func demoHandler(w http.ResponseWriter, r *http.Request) {
	jResult := result {
		Status: "Ok",
		Remote: r.RemoteAddr,
		Method: r.Method,
		URI:    r.RequestURI,
	}
	resp, _ := json.Marshal(jResult)
	// print request
	log.Printf("%s: %s %s %s", r.RemoteAddr, r.Method, r.RequestURI, r.Proto)
	// print request headers
	var hdrs bytes.Buffer
	err := r.Header.Write(&hdrs)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Headers:\n%s\n", &hdrs)
	// print & send response body
	log.Printf("<- %s", resp)
	fmt.Fprintf(w, "%s", resp)  // send response body
}


func main() {
	flag.Parse()
	lport := fmt.Sprintf(":%d", *port)
	s := &http.Server{
		Addr        : lport,
		Handler     : http.HandlerFunc(demoHandler),
	}
	log.Printf("starting HTTP server on %s", s.Addr)
	log.Fatal(s.ListenAndServe())
}
