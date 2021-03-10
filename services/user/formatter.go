package user

type UserFormatter struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Address  string `json:"address"`
	RoleId   uint   `json:"role_id"`
	Token    string `json:"token"`
	Phone    string `json:"image_url"`
}

func FormatUser(user User, token string) UserFormatter {
	formatter := UserFormatter{
		ID:       user.ID,
		Name:     user.Name,
		Username: user.Username,
		Address:  user.Address,
		Phone:    user.Phone,
		RoleId:   user.RoleId,
		Token:    token,
	}

	return formatter
}
