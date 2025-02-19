package messages

var Users = map[string]string{
	// success message

	// client error
	"not_found": "User not found",

	"user_already_exist":  "User already exist",
	"insuficient_balance": "Insufficient balance",

	// server error
	"secret_key_error": "An error occurred while generate secret key.",
}
