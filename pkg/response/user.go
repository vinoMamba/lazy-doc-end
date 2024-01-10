package response

type LoginResponse struct {
	Token string `json:"token"`
}

type UserInfoResponse struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}
