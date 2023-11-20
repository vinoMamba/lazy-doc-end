package response

type UserRegisterResponse struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	UserId   int64  `json:"userId"`
	Token    string `json:"token"`
}
