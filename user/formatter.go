package user

type UserFormatter struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Address  string `json:"address"`
	Role     string `json:"role"`
	Token    string `json:"token"`
	Phone    string `json:"phone"`
}

func FormatUser(user User, token string) UserFormatter {
	formatter := UserFormatter{
		ID:       user.ID,
		Name:     user.Name,
		Username: user.Username,
		Address:  user.Address,
		Phone:    user.Phone,
		Role:     user.Role,
		Token:    token,
	}

	return formatter
}
