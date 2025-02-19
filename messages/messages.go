package messages

var Common = map[string]string{
	// Server error
	"internal_server_error": "An error occurred while processing your request.",
	"marshaling_data":       "An error occurred while processing the data.",
	"set_cache":             "An error occurred while setting the cache.",
	"get_cache":             "An error occurred while getting the cache.",

	// Client error
	"invalid_request":      "The request is invalid. Please check the input and try again.",
	"invalid_token":        "The token provided is invalid or expired.",
	"unprocessible_entity": "The request is invalid. Please check the input and try again.",
	"conflict":             "The data you entered already exists.",
}
