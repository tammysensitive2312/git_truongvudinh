package entities

import (
	"time"
)

type Project struct {
	ID               int64      `gorm:"primaryKey;autoIncrement"`
	Name             string     `gorm:"size:255;not null"`
	ProjectStartedAt time.Time  `gorm:"not null"`
	ProjectEndedAt   *time.Time `gorm:"default:null"`
	UserID           int64      `gorm:"not null"`
}
