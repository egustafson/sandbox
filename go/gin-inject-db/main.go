package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Model struct {
	db *DB
}

func (mdl *Model) List(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, mdl.db.List())
}

func main() {
	mdl := &Model{ // initalize "DB" into Model "handler"
		db: PreloadDB(map[string]string{
			"a": "a-value",
			"b": "b-value",
		}),
	}

	r := gin.Default()
	r.GET("/", mdl.List)

	r.Run(":8080")
}
