package response

type LoginResponse struct {
	Token    string `json:"token"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
}

type UserInfoResponse struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}
