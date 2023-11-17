package response

type UserRegisterResponse struct {
	Avatar   string `json:"avatar"`
	Email    string `json:"email"`
	Username string `json:"username"`
	UserId   int64  `json:"userId"`
	Token    string `json:"token"`
}
