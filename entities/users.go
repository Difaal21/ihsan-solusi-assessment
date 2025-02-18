package entities

import "time"

type Users struct {
	ID                    int        `json:"id"`
	Name                  string     `json:"name"`
	Email                 string     `json:"email"`
	IsEmailVerified       bool       `json:"is_email_verified"`
	EmailVerifiedAt       *time.Time `json:"email_verified_at"`
	PhoneNumber           string     `json:"phone_number"`
	IsPhoneNumberVerified bool       `json:"is_phone_number_verified"`
	PhoneNumberVerifiedAt *time.Time `json:"phone_number_verified_at"`
	Password              string     `json:"password"`
	IsMFAEnabled          bool       `json:"is_mfa_enabled"`
	MFASecretKey          string     `json:"mfa_secret_key"`
	MFARecoveryCode       string     `json:"mfa_recovery_code"`
	CreatedAt             time.Time  `json:"created_at"`
}

type UserAccount struct {
	ID                    int
	UserID                int
	Email                 string
	IsEmailVerified       bool
	EmailVerifiedAt       *time.Time
	PhoneNumber           string
	IsPhoneNumberVerified bool
	PhoneNumberVerifiedAt *time.Time
	Password              string
	PasswordSalt          string
	IsMFAEnabled          bool
	CreatedAt             time.Time
}

type UserLoginSessionCache struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	CreatedAt   time.Time `json:"created_at"`
}

type UserTOTP struct {
	UserID       int    `json:"user_id"`
	SecretKey    string `json:"secret"`
	RecoveryCode string `json:"recovery_code"`
}

type UserProfileCache struct {
	ID                    int        `json:"id"`
	Name                  string     `json:"name"`
	Email                 string     `json:"email"`
	IsEmailVerified       bool       `json:"is_email_verified"`
	EmailVerifiedAt       time.Time  `json:"email_verified_at"`
	PhoneNumber           string     `json:"phone_number"`
	IsPhoneNumberVerified bool       `json:"is_phone_number_verified"`
	PhoneNumberVerifiedAt *time.Time `json:"phone_number_verified_at"`
	CreatedAt             time.Time  `json:"created_at"`
}
