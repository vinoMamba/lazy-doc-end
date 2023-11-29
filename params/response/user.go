package response

type UserRegisterResponse struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	UserId   int64  `json:"userId"`
}
type UserLoginResponse struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Token    string `json:"token"`
}
