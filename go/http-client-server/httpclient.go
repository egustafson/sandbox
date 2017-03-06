package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var host = flag.String("host", "localhost:9900", "connect to (default: localhost:9900)")
var path = flag.String("path", "/", "GET path (default: '/')")

func main() {
	flag.Parse()
	url := fmt.Sprintf("http://%s%s", *host, *path)
	// make request
	log.Printf("GET %s", url)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	// print response code
	fmt.Printf("%s\n", resp.Status)
	// print headers
	var hdrs bytes.Buffer
	err = resp.Header.Write(&hdrs)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", &hdrs)
	// print body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", body)
}
