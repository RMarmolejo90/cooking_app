package handlers

import (
	"net/http"

	"github.com/RMarmolejo90/cooking_app/internal/database"
	"github.com/RMarmolejo90/cooking_app/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus" // or any logger you prefer
)

var db = database.DB

func CreateRecipe(c *gin.Context) {
	var recipe models.Recipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		logrus.WithError(err).Error("Failed to bind JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate required fields
	if recipe.Title == "" || recipe.UserID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title and UserID are required"})
		return
	}

	if result := db.Create(&recipe); result.Error != nil {
		logrus.WithError(result.Error).Error("Failed to create recipe")
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, recipe)
}
