package main

import (
	// point browser at /debug/vars
	"expvar"
	"flag"
	"fmt"
	"log"
	"net/http"
)

var port = flag.Int("port", 9000, "listen on port (default: 9000)")

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola, Mundo!")
}

func expfunc() interface{} {
	return 123
}

func main() {
	flag.Parse()
	lport := fmt.Sprintf(":%d", *port)

	ver := expvar.NewString("version")
	ver.Set("v1.0.1")

	var fv expvar.Func = expfunc
	expvar.Publish("intfn", fv)

	http.HandleFunc("/", helloHandler)
	log.Printf("starting HTTP server on %s", lport)
	log.Fatal(http.ListenAndServe(lport, nil))
}
