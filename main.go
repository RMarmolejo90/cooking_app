package main

import (
	"log"
	"net/http"

	"github.com/RMarmolejo90/cooking_app/internal/database"
	"github.com/gin-gonic/gin"
)

func main() {

	// attempt database connection
	database.ConnectDB()

	r := gin.Default()

	// Ping test
	r.GET("/", func(c *gin.Context) {
		log.Printf("\n\t--- Server is Running ---\n")
		c.String(http.StatusOK, "server running")
	})
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
