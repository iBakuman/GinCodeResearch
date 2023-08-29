package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mattn/go-colorable"
	"log"
	"net/http"
	"os"
)

func main() {
	gin.DefaultWriter = colorable.NewColorableStdout()
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	// Starts a new Gin instance with no middle-ware
	r := gin.New()

	// Define handlers
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	})
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Listen and serve on defined port
	log.Printf("Listening on port %s", port)
	r.Run(":" + port)
}
