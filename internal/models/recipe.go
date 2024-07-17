package models

import (
	"gorm.io/gorm"
)

type Recipe struct {
	Title           string    `gorm:"not null"`
	Description     string    `gorm:"type:TEXT"`
	Ingredients     string    `gorm:"type:TEXT"` // List of ingredients
	Instructions    string    `gorm:"type:TEXT"` // Cooking instructions
	UserID          uint      `gorm:"not null"`  // The user who created the recipe
	User            User      `gorm:"foreignKey:UserID"`
	Gallery         []Gallery `gorm:"foreignKey:RecipeID"`
	Tags            []Tag     `gorm:"many2many:recipe_tags;"`
	AverageRating   float64   `gorm:"-"` // Field to store the average rating, not stored in DB
	NumberOfRatings int       `gorm:"-"` // Field to store the number of ratings, not stored in DB
	gorm.Model
}
