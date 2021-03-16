package inventory

import (
	"market-fish-service/user"
)

type GetItemDetailInput struct {
	ID uint `uri:"id" binding:"required"`
}

type ItemInput struct {
	Name        string `form:"name" `
	Price       int64  `form:"price" binding:"required"`
	Description string `form:"description" binding:"required"`
	Stock       int    `form:"stock" binding:"required"`
	User        user.User
}
