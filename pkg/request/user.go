package request

type CreateUserRequest struct {
	Email           string `json:"email" valid:"email,required"`
	Password        string `json:"password" valid:"required,stringlength(6|20)"`
	ConfirmPassword string `json:"confirmPassword" valid:"required,stringlength(6|20)"`
}
