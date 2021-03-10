package user

import "time"

type User struct {
	ID           uint
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Name         string
	Username     string
	PasswordHash string
	Phone        string
	Address      string
	Role         string
}
