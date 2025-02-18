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

type UserMFAStatusOnLoginResponse struct {
	IsMFAEnabled   bool `json:"mfa_is_enabled"`
	IsEmailEnabled bool `json:"email_is_enabled"`
}

func NewUserMFAStatusOnLoginResponse(isMFAEnabled bool, isEmailEnabled bool) (result *UserMFAStatusOnLoginResponse) {
	return &UserMFAStatusOnLoginResponse{
		IsMFAEnabled:   isMFAEnabled,
		IsEmailEnabled: isEmailEnabled,
	}
}

type UserTOTPURLResponse struct {
	TOTPURL string `json:"totp_url"`
}

func NewUserTOTPURLResponse(totpURL string) (result *UserTOTPURLResponse) {
	return &UserTOTPURLResponse{
		TOTPURL: totpURL,
	}
}

type UserTokenResponse struct {
	AccessToken AccessToken `json:"access_token"`
}

type AccessToken struct {
	Value     string `json:"value"`
	TokenType string `json:"token_type"`
	ExpiresIn int    `json:"expires_in"`
}

func NewUserTokenResponse(accessToken string, tokenType string, expiresIn int) (result *UserTokenResponse) {
	return &UserTokenResponse{
		AccessToken: AccessToken{
			Value:     accessToken,
			TokenType: tokenType,
			ExpiresIn: expiresIn,
		},
	}
}
