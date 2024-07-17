package models

type Post struct {
	ID      uint   `gorm:"primaryKey"`
	Title   string `gorm:"not null"`
	Content string `gorm:"type:TEXT"`
	Tags    []Tag  `gorm:"many2many:post_tags;"`
}
