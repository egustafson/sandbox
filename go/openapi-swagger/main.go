package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	middleware "github.com/oapi-codegen/gin-middleware"

	"openapi-swagger/api" // <-- this repository
)

func initStoreAPI(r gin.IRouter) error {

	oapiSpec, err := api.GetSwagger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get OpenAPI spec: %v\n", err)
		return err
	}

	// Clear servers field; this causes the validator to ignore server URL
	// checks.
	//oapiSpec.Servers = nil

	r.Use(middleware.OapiRequestValidator(oapiSpec))

	api.RegisterHandlers(r, api.NewItemStore()) // <-- register the API handler
	return nil
}

func main() {
	router := gin.Default()

	router.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	// bind the OpenAPI (oapi-codegen) generated API to the path /api.  This
	// coresponds to the `servers.url` field in the API spec, (openapi.yaml in
	// the root of this project).
	theAPI := router.Group("/api")
	initStoreAPI(theAPI)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	fmt.Println("Starting server on :8080")
	if err := srv.ListenAndServe(); err != nil {
		fmt.Fprintf(os.Stderr, "Server failed: %v\n", err)
	}
	// ListenAndServe is blocking; server runs until interrupted
}
