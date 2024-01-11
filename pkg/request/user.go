package request

type CreateUserRequest struct {
	Email           string `json:"email" valid:"email,required"`
	Password        string `json:"password" valid:"required,stringlength(6|20)"`
	ConfirmPassword string `json:"confirmPassword" valid:"required,stringlength(6|20)"`
}

type LoginRequest struct {
	Email    string `json:"email" valid:"email,required"`
	Password string `json:"password" valid:"required"`
}

type UpdatePasswordRequest struct {
	OldPassword string `json:"oldPassword" valid:"required"`
	NewPassword string `json:"newPassword" valid:"required"`
}
