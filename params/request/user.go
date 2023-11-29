package request

type UserRegisterRequest struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserUpdateUsernameRequest struct {
	Username string `json:"username"`
}

type UserUpdateEmailRequest struct {
	Email string `json:"email"`
}

type UserUpdatePasswordRequest struct {
	NewPassword string `json:"newPassword"`
	OldPassword string `json:"oldPassword"`
}
