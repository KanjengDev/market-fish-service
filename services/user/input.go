package user

type RegisterUserInput struct {
	Name     string `json:"name" binding:"required"`
	Username string `json:"username"  binding:"required"`
	Password string `json:"password" binding:"required"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	RoleId   uint   `json:"role_id" binding:"required"`
}

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
