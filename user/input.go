package user

type RegisterUserInput struct {
	Name     string `json:"name" form:"name" binding:"required"`
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
	Phone    string `json:"phone" form:"phone" binding:"required"`
	Address  string `json:"address" form:"address" binding:"required"`
	Role     string `json:"role" form:"role" binding:"required"`
}

type LoginInput struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}
