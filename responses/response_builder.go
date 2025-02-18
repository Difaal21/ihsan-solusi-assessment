package responses

func (r *ResponsesImpl) SetCode(code int) Responses {
	r.Code = code
	return r
}

func (r *ResponsesImpl) SetData(data interface{}) Responses {
	r.Data = data
	return r
}

func (r *ResponsesImpl) SetMessage(message string) Responses {
	r.Message = message
	return r
}

func (r *ResponsesImpl) SetStatus(status string, defaultStatus string) Responses {
	if status != "" {
		r.Status = status
		return r
	}
	r.Status = defaultStatus
	return r
}

func (r *ResponsesImpl) SetPagination(pagination interface{}) Responses {
	r.Pagination = pagination
	return r
}

func (r *ResponsesImpl) Send() (response *ResponsesImpl) {
	return r
}
