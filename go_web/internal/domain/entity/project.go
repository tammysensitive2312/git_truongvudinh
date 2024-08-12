package entity

import "time"

type Project struct {
	ID               int64      `gorm:"primaryKey,autoIncrement"`
	Name             string     `gorm:"size:255,notnull"`
	ProjectStartedAt time.Time  `gorm:"notnull"`
	ProjectEndedAt   *time.Time `gorm:"default:null"`
	UserID           int        `gorm:"not null"`
	User             User       `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE"`
}
