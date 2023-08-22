package main

// the OS gdbm *developer* package must be installed (and build-essential)

import (
	"fmt"
	"time"

	"github.com/cfdrake/go-gdbm"
)

const (
	db_filename = "demo.gdbm"

	demo_key = "demo-key"
)

func main() {
	db, err := gdbm.Open(db_filename, "c")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if val, err := db.Fetch(demo_key); err == nil {
		fmt.Printf("Prev Value:  %s\n", val)
	}

	newVal := time.Now().UTC().Format(time.RFC3339)
	if err := db.Replace(demo_key, newVal); err != nil {
		panic(err)
	}
	db.Sync()
}
