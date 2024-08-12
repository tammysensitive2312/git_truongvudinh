package entities

import (
	"time"
)

type User struct {
	ID        int    `gorm:"primaryKey;autoIncrement"`
	Email     string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	FirstName string `gorm:"size:255;not null"`
	LastName  string `gorm:"size:255;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Projects []Project `gorm:"foreignKey:UserID"`
}
