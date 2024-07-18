package handlers

import (
	"net/http"

	"github.com/RMarmolejo90/cooking_app/internal/database"
	"github.com/RMarmolejo90/cooking_app/internal/models"
	"github.com/gin-gonic/gin"
)

func CreateTag(c *gin.Context) {
	var tag models.Tag
	if err := c.ShouldBindJSON(&tag); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result := database.DB.Create(&tag); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, tag)
}

func GetTag(c *gin.Context) {
	id := c.Param("id")
	var tag models.Tag

	if err := database.DB.First(&tag, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tag not found"})
		return
	}

	c.JSON(http.StatusOK, tag)
}

func UpdateTag(c *gin.Context) {
	id := c.Param("id")
	var tag models.Tag

	if err := database.DB.First(&tag, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tag not found"})
		return
	}

	if err := c.ShouldBindJSON(&tag); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result := database.DB.Save(&tag); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, tag)
}

func DeleteTag(c *gin.Context) {
	id := c.Param("id")
	var tag models.Tag

	if err := database.DB.First(&tag, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tag not found"})
		return
	}

	if result := database.DB.Delete(&tag); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tag deleted successfully"})
}

func GetPostsByTag(c *gin.Context) {
	tagName := c.Param("tagName")
	var tag models.Tag

	if err := database.DB.Preload("Posts").Where("tag_name = ?", tagName).First(&tag).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tag not found"})
		return
	}

	c.JSON(http.StatusOK, tag.Posts)
}

func GetRecipesByTag(c *gin.Context) {
	tagName := c.Param("tagName")
	var tag models.Tag

	if err := database.DB.Preload("Recipes").Where("tag_name = ?", tagName).First(&tag).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tag not found"})
		return
	}

	c.JSON(http.StatusOK, tag.Recipes)
}
