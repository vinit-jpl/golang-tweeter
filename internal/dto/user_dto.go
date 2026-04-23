package dto

type RegisterRequest struct {
	Email           string `json:"email"`
	Username        string `json:"username"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"passowrd_confirm"`
}

type RegisterResponse struct {
	ID int64 `json:"id"`
}
