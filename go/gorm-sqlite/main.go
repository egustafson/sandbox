package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Rec struct {
	gorm.Model
	Key string
	Val string
}

func main() {

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect to database")
	}
	defer db.Close()

	db.AutoMigrate(&Rec{})

	db.Create(&Rec{Key: "key-1", Val: "value-1"})

	var r Rec
	db.First(&r, 1)
	db.First(&r, "key = ?", "key-1")

	db.Model(&r).Update("Value", "value-updated")

	db.Delete(&r)
}
