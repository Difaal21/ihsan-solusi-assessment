package messages

var Users = map[string]string{
	// success message
	"login_success": "Login success",
	"available_mfa": "List available Multi-Factor Authentication",
	"generate_otp":  "Generate OTP success",
	"totp_verify":   "TOTP verification success",

	// client error
	"not_found":          "User not found",
	"email_not_verified": "Email is not verified",
	"invalid_credential": "Invalid email or password",
	"mfa_inactive":       "Multi-Factor Authentication is inactive",
	"mfa_registered":     "User already registered Multi-Factor Authentication",
	"invalid_totp":       "Invalid TOTP code",

	// server error
	"secret_key_error": "An error occurred while generate secret key.",
}
