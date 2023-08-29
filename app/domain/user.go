package domain

import (
	"time"
)

type User struct {
	ID          string
	Name        string
	Password    string
	PhoneNumber string
	Email       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}
