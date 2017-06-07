package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

const (
	dbFilename = "./stub.db"
)

func main() {
	os.Remove(dbFilename)

	db, err := sql.Open("sqlite3", dbFilename)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := `
      create table foo (id integer not null primary key, name text);
      delete from foo;
    `

	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}

	printDrivers()

	fmt.Printf("db: %V\n", db)

	fmt.Println("--")
	fmt.Println("done.")
}

func printDrivers() {
	fmt.Println("sql Drivers:")
	for d := range sql.Drivers() {
		fmt.Printf("  %v\n", d)
	}
	fmt.Println("--")
}
