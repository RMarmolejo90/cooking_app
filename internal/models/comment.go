package models

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	ID        uint      `gorm:"primaryKey"`
	RecipeID  uint      `gorm:"not null"` // The recipe this comment is associated with
	UserID    uint      `gorm:"not null"` // The user who posted the comment
	Content   string    `gorm:"type:TEXT;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	Recipe    Recipe    `gorm:"foreignKey:RecipeID"`
	User      User      `gorm:"foreignKey:UserID"`
}
