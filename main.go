package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

func setupRouter() *gin.Engine {

	r := gin.Default()

	// Ping test
	r.GET("/", func(c *gin.Context) {
		log.Printf("\n\t--- Server is Running ---\n")
		c.String(http.StatusOK, "server running")
	})
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
