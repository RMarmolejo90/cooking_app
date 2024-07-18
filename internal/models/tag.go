package models

type Tag struct {
	ID      uint     `gorm:"primaryKey"`
	TagName string   `gorm:"unique;not null"`
	Posts   []Post   `gorm:"many2many:post_tags;"`
	Recipes []Recipe `gorm:"many2many:recipe_tags;"`
}
