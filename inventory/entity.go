package inventory

import (
	"market-fish-service/user"
	"time"
)

type Inventory struct {
	ID          uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string
	Price       int64
	FileName    string
	Description string
	UserID      uint
	Stock       int
	User        user.User
}
