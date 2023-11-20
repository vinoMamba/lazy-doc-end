package request

type UserRegisterRequest struct {
	Username        string `json:"username"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

type UserLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
