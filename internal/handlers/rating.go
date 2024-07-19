package handlers

import (
	"net/http"

	"github.com/RMarmolejo90/cooking_app/internal/models"
	"github.com/gin-gonic/gin"
)

func CreateRating(c *gin.Context) {
	var rating models.Rating
	if err := c.ShouldBindJSON(&rating); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error parsing JSON: " + err.Error()})
		return
	}

	if result := db.Create(&rating); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error saving the rating to the database"})
		return
	}

	c.JSON(http.StatusOK, rating)
}

func GetRating(c *gin.Context) {
	id := c.Param("id")
	var rating models.Rating
	if err := db.First(&rating, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Rating not found"})
		return
	}
	c.JSON(http.StatusOK, rating)
}

func UpdateRating(c *gin.Context) {
	id := c.Param("id")
	var rating models.Rating

	if err := db.First(&rating, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Rating not found"})
		return
	}

	if err := c.ShouldBindJSON(&rating); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error parsing JSON: " + err.Error()})
		return
	}

	if result := db.Save(&rating); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error saving to the database"})
		return
	}

	c.JSON(http.StatusOK, rating)
}

func DeleteRating(c *gin.Context) {
	id := c.Param("id")
	var rating models.Rating

	if err := db.First(&rating, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Rating not found"})
		return
	}

	if result := db.Delete(&rating); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting from the database"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully deleted!"})
}
