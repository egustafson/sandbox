package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

const (
	DSN = "file:demo.db"
)

func main() {
	db, err := sqlx.Connect("sqlite3", DSN)
	if err != nil {
		panic("failed to connect to the database: " + err.Error())
	}
	err = db.Ping()
	if err != nil {
		panic("failed to ping the database: " + err.Error())
	}

	fmt.Println("done.")
}
