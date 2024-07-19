package handlers

import (
	"net/http"

	"github.com/RMarmolejo90/cooking_app/internal/database"
	"github.com/RMarmolejo90/cooking_app/internal/models"
	"github.com/gin-gonic/gin"
)

func CreateMessage(c *gin.Context) {
	var message models.Message

	if err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error binding data to json"})
		return
	}

	if result := db.Create(&message); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error while saving to the database"})
		return
	}

	c.JSON(http.StatusOK, message)
}
