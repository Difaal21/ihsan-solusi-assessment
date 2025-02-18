package dto

type UserLoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type UserTOTPVerification struct {
	Code string `json:"code" validate:"required,min=6"`
}
