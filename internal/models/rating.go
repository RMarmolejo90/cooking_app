package models

import (
	"time"
)

type Rating struct {
	ID        uint      `gorm:"primaryKey"`
	RecipeID  uint      `gorm:"not null"` // The recipe this rating is associated with
	UserID    uint      `gorm:"not null"` // The user who gave the rating
	Rating    int       `gorm:"not null"` // The rating value (e.g., 1-5)
	CreatedAt time.Time `gorm:"autoCreateTime"`
	Recipe    Recipe    `gorm:"foreignKey:RecipeID"`
	User      User      `gorm:"foreignKey:UserID"`
}
