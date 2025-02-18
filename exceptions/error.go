package exceptions

type LogError struct {
	ID   any    `json:"log_id"`
	Err  *error `json:"error,omitempty"`
	Data any    `json:"data,omitempty"`
}
