package entity

import "time"

type User struct {
	ID        int64
	Email     string
	Password  string
	FirstName string
	LastName  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
