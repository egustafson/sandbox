package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func healthzHandler(c *gin.Context) {
	hn, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, gin.H{"hostname": hn})
}

func mdlGetHandler(c *gin.Context) {
	typeStr := c.Param("type")
	idStr := c.Param("id")

	resp := gin.H{"type": typeStr}
	if len(idStr) > 0 {
		resp["id"] = idStr
	}
	c.JSON(http.StatusOK, resp)
}

func main() {
	router := gin.Default()
	router.GET("/healthz", healthzHandler)
	router.GET("/mdl/:type/*id", mdlGetHandler)

	router.Run() // listen & serve 0.0.0.0:8080
}
