package models

import (
	"gorm.io/gorm"
)

type Gallery struct {
	ID       uint   `gorm:"primaryKey"`
	ImageURL string `gorm:"type:TEXT;not null"`
	RecipeID uint   `gorm:"not null"` // The recipe this image is associated with
	Recipe   Recipe `gorm:"foreignKey:RecipeID"`
	UserID   uint   `gorm:"not null"` // The user who posted the image
	User     User   `gorm:"foreignKey:UserID"`
}
