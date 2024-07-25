package models

type CreateUserRequest struct {
	UserName *string `json:"user_name"`
	Password *string `json:"password"`
	Email    *string `json:"email"`
}

type LoginRequest struct {
	UserName *string `json:"user_name"`
	Password *string `json:"password"`
}

type LoginResponse struct {
	ID      *uint   `json:"id"`
	Token   *string `json:"token"`
	UseName *string `json:"user_name"`
	Email   *string `json:"email"`
}
