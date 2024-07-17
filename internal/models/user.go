package models

import "gorm.io/gorm"

type User struct {
	Username     string `gorm:"unique;not null"`
	Email        string `gorm:"unique;not null"`
	ProfileImage string `gorm:"type:TEXT"`
	Profile      string `gorm:"type:TEXT"` // Profile description - short paragraph
	gorm.Model
}
