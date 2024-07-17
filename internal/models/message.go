package models

import (
	"time"

	"gorm.io/gorm"
)

type Message struct {
	ID         uint      `gorm:"primaryKey"`
	SenderID   uint      `gorm:"not null"` // The user who sent the message
	ReceiverID uint      `gorm:"not null"` // The user who received the message
	Content    string    `gorm:"type:TEXT;not null"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	Sender     User      `gorm:"foreignKey:SenderID"`
	Receiver   User      `gorm:"foreignKey:ReceiverID"`
}
