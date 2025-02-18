package responses

import "net/http"

// Mapping Client Error Responses https://developer.mozilla.org/en-US/docs/Web/HTTP/Status#client_error_responses

func (r *ResponsesImpl) BadRequest(status string) *ResponsesImpl {
	r.SetCode(http.StatusBadRequest)
	r.SetStatus(status, "BAD_REQUEST")
	return r
}

func (r *ResponsesImpl) Unauthorized(status string) *ResponsesImpl {
	r.SetCode(http.StatusUnauthorized)
	r.SetStatus(status, "UNAUTHORIZED")
	return r
}

func (r *ResponsesImpl) Forbidden(status string) *ResponsesImpl {
	r.SetCode(http.StatusForbidden)
	r.SetStatus(status, "FORBIDDEN")
	return r
}

func (r *ResponsesImpl) NotFound(status string) *ResponsesImpl {
	r.SetCode(http.StatusNotFound)
	r.SetStatus(status, "NOT_FOUND")
	return r
}

func (r *ResponsesImpl) Conflict(status string) *ResponsesImpl {
	r.SetCode(http.StatusConflict)
	r.SetStatus(status, "CONFLICT")
	return r
}

func (r *ResponsesImpl) UnprocessableEntity(status string) *ResponsesImpl {
	r.SetCode(http.StatusUnprocessableEntity)
	r.SetStatus(status, "UNPROCESSABLE_ENTITY")
	return r
}

func (r *ResponsesImpl) TooManyRequests(status string) *ResponsesImpl {
	r.SetCode(http.StatusTooManyRequests)
	r.SetStatus(status, "TOO_MANY_REQUESTS")
	return r
}
