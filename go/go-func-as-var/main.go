package main

import (
	"fmt"
	"log"
)

var output func(msg string) (err error)

func fmtOutput(msg string) (err error) {
	_, err = fmt.Print(msg)
	return
}

func logOutput(msg string) (err error) {
	log.Print(msg)
	return nil
}

func init() {
	output = logOutput
}

func main() {
	output("Hola, Mundo.")
}
