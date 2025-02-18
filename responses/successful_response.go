package responses

import "net/http"

// Mapping Successful Responses https://developer.mozilla.org/en-US/docs/Web/HTTP/Status#successful_responses

func (r *ResponsesImpl) Ok(status string) *ResponsesImpl {
	r.SetCode(http.StatusOK)
	r.SetStatus(status, "OK")
	return r
}

func (r *ResponsesImpl) Created(status string) *ResponsesImpl {
	r.SetCode(http.StatusCreated)
	r.SetStatus(status, "CREATED")
	return r
}
