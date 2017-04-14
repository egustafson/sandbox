// main.go
//
// lifted from blog
//  - https://semaphoreci.com/community/tutorials/building-go-web-applications-and-microservices-using-gin
//
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var router *gin.Engine

func main() {
	router = gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(
			http.StatusOK,
			"index.gohtml",
			gin.H{
				"title": "Gin Demo - Home Page",
			},
		)
	})
	router.Run()
}
