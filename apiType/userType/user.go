package userType

type RegisterRequest struct {
	UserName string `json:"user_name" validate:"required"`
	Password string `json:"password" validate:"required"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
}

type RegisterResponse struct {
	Token string `json:"token"`
}
