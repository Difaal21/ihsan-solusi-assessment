package responses

import "net/http"

// Mapping Server Error Responses https://developer.mozilla.org/en-US/docs/Web/HTTP/Status#server_error_responses

func (r *ResponsesImpl) InternalServerError(status string) *ResponsesImpl {
	r.SetCode(http.StatusInternalServerError)
	r.SetStatus(status, "INTERNAL_SERVER_ERROR")
	return r
}
