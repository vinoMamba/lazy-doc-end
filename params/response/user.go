package response

type UserRegisterResponse struct {
	Avatar   string `json:"avatar"`
	Email    string `json:"email"`
	Username string `json:"username"`
	UserId   string `json:"userId"`
	Token    string `json:"token"`
}
