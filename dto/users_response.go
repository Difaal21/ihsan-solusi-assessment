package dto

import (
	"time"

	"github.com/google/uuid"
)

type UserLoginResponse struct {
	LoginSession          uuid.UUID     `json:"login_session"`
	LoginSessionExpiresIn time.Duration `json:"login_session_expires_in"`
}

func NewUserLoginResponse(loginSesion uuid.UUID, ttl time.Duration) (result *UserLoginResponse) {
	return &UserLoginResponse{
		LoginSession:          loginSesion,
		LoginSessionExpiresIn: ttl,
	}
}
